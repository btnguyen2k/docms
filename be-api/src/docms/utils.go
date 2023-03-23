package docms

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v3"
)

const (
	metaFileYaml    = "meta.yaml"
	metaFileJson    = "meta.json"
	logLevelInfo    = "INFO"
	logLevelWarning = "WARN"
	logLevelError   = "ERROR"
)

var (
	gDataDir         string
	gSiteMeta        *SiteMeta
	gTopicList       = make([]*TopicMeta, 0)              // list of categories, sorted by index
	gTopicMeta       = make(map[string]*TopicMeta)        // map[category-id]category-metadata
	gDocumentList    = make(map[string][]*DocumentMeta)   // list of documents, per category, sorted by index
	gDocumentMeta    = make(map[string]*DocumentMeta)     // map[category-id:document-id]document-metadata
	gDocumentContent = make(map[string]map[string]string) // map[category-id:document-id]map[language-code]document-content
)

// SiteMeta capture metadata of the website.
type SiteMeta struct {
	Name            string            `json:"name",yaml:"name"`               // name of the website
	Description     interface{}       `json:"description",yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Languages       map[string]string `json:"languages",yaml:"languages"`     // available languages of the website content
	DefaultLanguage string            `json:"-",yaml:"-"`                     // site's default language
	Icon            string            `json:"icon",yaml:"icon"`               // website's icon
	Contacts        map[string]string `json:"contacts",yaml:"contacts"`       // site's contact info
}

func (sm *SiteMeta) init() error {
	sm.DefaultLanguage = sm.Languages["default"]
	for k, v := range sm.Contacts {
		if v == "" {
			delete(sm.Contacts, k)
		}
	}
	return nil
}

func (sm *SiteMeta) GetDescriptionMap() map[string]string {
	result := make(map[string]string)
	switch sm.Description.(type) {
	case string:
		result[sm.DefaultLanguage] = sm.Description.(string)
	case map[string]string:
		for k, v := range sm.Description.(map[string]string) {
			result[k] = v
		}
	case map[string]interface{}:
		for k, v := range sm.Description.(map[string]interface{}) {
			result[k] = fmt.Sprintf("%s", v)
		}
	}
	return result
}

func LoadSiteMeta(yamlFilePath, jsonFilePath string) (*SiteMeta, error) {
	if _, err := os.Stat(yamlFilePath); err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		return LoadSiteMetaFromYaml(yamlFilePath)
	}
	if _, err := os.Stat(jsonFilePath); err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		return LoadSiteMetaFromJson(jsonFilePath)
	}
	return nil, fmt.Errorf("neither <%s> nor <%s> exist", yamlFilePath, jsonFilePath)
}

func LoadSiteMetaFromYaml(filePath string) (*SiteMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *SiteMeta
	err = yaml.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.init()
	}
	return metadata, err
}

func LoadSiteMetaFromJson(filePath string) (*SiteMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *SiteMeta
	err = json.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.init()
	}
	return metadata, err
}

/*----------------------------------------------------------------------*/

// TopicMeta capture metadata of a category.
type TopicMeta struct {
	index       int         `json:"-",yaml:"-"`                     // topic index, for ordering
	id          string      `json:"-",yaml:"-"`                     // topic id
	dir         string      `json:"-",yaml:"-"`                     // name of directory where topic's data locates
	Title       interface{} `json:"title",yaml:"title"`             // topic's title, can be a single string, or a map[language-code:string]string
	Description interface{} `json:"description",yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon",yaml:"icon"`               // topic's icon
}

func (tm *TopicMeta) setDirectory(dir string) bool {
	tm.dir = dir
	if !RexpContentDir.MatchString(dir) {
		return false
	}
	matches := RexpContentDir.FindStringSubmatch(dir)
	tm.index, _ = strconv.Atoi(matches[1])
	tm.id = matches[2]
	return true
}

func (tm *TopicMeta) GetDescriptionMap() map[string]string {
	result := make(map[string]string)
	switch tm.Description.(type) {
	case string:
		result[gSiteMeta.DefaultLanguage] = tm.Description.(string)
	case map[string]string:
		for k, v := range tm.Description.(map[string]string) {
			result[k] = v
		}
	case map[string]interface{}:
		for k, v := range tm.Description.(map[string]interface{}) {
			result[k] = fmt.Sprintf("%s", v)
		}
	}
	return result
}

func (tm *TopicMeta) GetTitleMap() map[string]string {
	result := make(map[string]string)
	switch tm.Title.(type) {
	case string:
		result[gSiteMeta.DefaultLanguage] = tm.Title.(string)
	case map[string]string:
		for k, v := range tm.Title.(map[string]string) {
			result[k] = v
		}
	case map[string]interface{}:
		for k, v := range tm.Title.(map[string]interface{}) {
			result[k] = fmt.Sprintf("%s", v)
		}
	}
	return result
}

func LoadTopicMeta(yamlFilePath, jsonFilePath string) (*TopicMeta, error) {
	if _, err := os.Stat(yamlFilePath); err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		return LoadTopicMetaFromYaml(yamlFilePath)
	}
	if _, err := os.Stat(jsonFilePath); err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		return LoadTopicMetaFromJson(jsonFilePath)
	}
	return nil, fmt.Errorf("neither <%s> nor <%s> exist", yamlFilePath, jsonFilePath)
}

func LoadTopicMetaFromYaml(filePath string) (*TopicMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *TopicMeta
	return metadata, yaml.Unmarshal(buf, &metadata)
}

func LoadTopicMetaFromJson(filePath string) (*TopicMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *TopicMeta
	return metadata, json.Unmarshal(buf, &metadata)
}

/*----------------------------------------------------------------------*/

// DocumentMeta capture metadata of a category.
type DocumentMeta struct {
	index       int         `json:"-",yaml:"-"`             // document index, for ordering
	id          string      `json:"-",yaml:"-"`             // document id
	dir         string      `json:"-",yaml:"-"`             // name of directory where document's data locates
	Title       interface{} `json:"title" yaml:"title"`     // title of the document, can be a single string, or a map[language-code:string]string
	Summary     interface{} `json:"summary" yaml:"summary"` // document summary, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon",yaml:"icon"`       // document's icon
	ContentFile interface{} `json:"file" yaml:"file"`       // name of document's content file, can be a single string, or a map[language-code:string]string
}

func (dm *DocumentMeta) setDirectory(dir string) bool {
	dm.dir = dir
	if !RexpContentDir.MatchString(dir) {
		return false
	}
	matches := RexpContentDir.FindStringSubmatch(dir)
	dm.index, _ = strconv.Atoi(matches[1])
	dm.id = matches[2]
	return true
}

func (dm *DocumentMeta) GetSummaryMap() map[string]string {
	result := make(map[string]string)
	switch dm.Summary.(type) {
	case string:
		result[gSiteMeta.DefaultLanguage] = dm.Summary.(string)
	case map[string]string:
		for k, v := range dm.Summary.(map[string]string) {
			result[k] = v
		}
	case map[string]interface{}:
		for k, v := range dm.Summary.(map[string]interface{}) {
			result[k] = fmt.Sprintf("%s", v)
		}
	}
	return result
}

func (dm *DocumentMeta) GetTitleMap() map[string]string {
	result := make(map[string]string)
	switch dm.Title.(type) {
	case string:
		result[gSiteMeta.DefaultLanguage] = dm.Title.(string)
	case map[string]string:
		for k, v := range dm.Title.(map[string]string) {
			result[k] = v
		}
	case map[string]interface{}:
		for k, v := range dm.Title.(map[string]interface{}) {
			result[k] = fmt.Sprintf("%s", v)
		}
	}
	return result
}

func (dm *DocumentMeta) GetContentFileMap() map[string]string {
	result := make(map[string]string)
	switch dm.ContentFile.(type) {
	case string:
		result[gSiteMeta.DefaultLanguage] = dm.ContentFile.(string)
	case map[string]string:
		for k, v := range dm.ContentFile.(map[string]string) {
			result[k] = v
		}
	case map[string]interface{}:
		for k, v := range dm.ContentFile.(map[string]interface{}) {
			result[k] = fmt.Sprintf("%s", v)
		}
	}
	return result
}

func LoadDocumentMeta(yamlFilePath, jsonFilePath string) (*DocumentMeta, error) {
	if _, err := os.Stat(yamlFilePath); err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		return LoadDocumentMetaFromYaml(yamlFilePath)
	}
	if _, err := os.Stat(jsonFilePath); err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else {
		return LoadDocumentMetaFromJson(jsonFilePath)
	}
	return nil, fmt.Errorf("neither <%s> nor <%s> exist", yamlFilePath, jsonFilePath)
}

func LoadDocumentMetaFromYaml(filePath string) (*DocumentMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *DocumentMeta
	return metadata, yaml.Unmarshal(buf, &metadata)
}

func LoadDocumentMetaFromJson(filePath string) (*DocumentMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *DocumentMeta
	return metadata, json.Unmarshal(buf, &metadata)
}

/*----------------------------------------------------------------------*/

var RexpContentDir = regexp.MustCompile(`^(\d)+-(\w+)$`)

var defaultDirFilter = func(entry os.DirEntry) bool {
	return entry.Name() != "." && entry.Name() != ".."
}

func GetDirContent(path string, filter func(entry os.DirEntry) bool) ([]os.DirEntry, error) {
	dirInfo, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if filter == nil {
		filter = defaultDirFilter
	}
	result := make([]os.DirEntry, 0)
	for _, entry := range dirInfo {
		if filter(entry) {
			result = append(result, entry)
		}
	}
	return result, nil
}
