package docms

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v3"
)

const (
	metaFile = "meta.yaml"
)

var (
	gDataDir         string
	gDefaultLanguage string
	gSiteMeta        *SiteMeta
	gTopicList       = make([]*TopicMeta, 0)            // list of categories, sorted by index
	gTopicMeta       = make(map[string]*TopicMeta)      // map[category-id]category-metadata
	gDocumentList    = make(map[string][]*DocumentMeta) // list of documents, per category, sorted by index
	gDocumentMeta    = make(map[string]*DocumentMeta)   // map[category-id:document-id]document-metadata
)

// SiteMeta capture metadata of the website.
type SiteMeta struct {
	Name        string            `json:"name",yaml:"name"`               // name of the website
	Description interface{}       `json:"description",yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Languages   map[string]string `json:"languages",yaml:"languages"`     // available languages of the website content
	Icon        string            `json:"icon",yaml:"icon"`               // website's icon
}

func (sm *SiteMeta) GetDescriptionMap() map[string]string {
	result := make(map[string]string)
	switch sm.Description.(type) {
	case string:
		result[gDefaultLanguage] = sm.Description.(string)
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

func LoadSiteMetaFromYaml(filePath string) (*SiteMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *SiteMeta
	return metadata, yaml.Unmarshal(buf, &metadata)
}

func LoadSiteMetaFromJson(filePath string) (*SiteMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *SiteMeta
	return metadata, json.Unmarshal(buf, &metadata)
}

/*----------------------------------------------------------------------*/

// TopicMeta capture metadata of a category.
type TopicMeta struct {
	index       int         `json:"-",yaml:"-"`                     // topic index, for ordering
	id          string      `json:"-",yaml:"-"`                     // topic id
	Title       interface{} `json:"title",yaml:"title"`             // topic's title, can be a single string, or a map[language-code:string]string
	Description interface{} `json:"description",yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon",yaml:"icon"`               // topic's icon
}

func (tm *TopicMeta) setIndexAndId(input string) bool {
	if !RexpContentDir.MatchString(input) {
		return false
	}
	matches := RexpContentDir.FindStringSubmatch(input)
	tm.index, _ = strconv.Atoi(matches[1])
	tm.id = matches[2]
	return true
}

func (tm *TopicMeta) GetDescriptionMap() map[string]string {
	result := make(map[string]string)
	switch tm.Description.(type) {
	case string:
		result[gDefaultLanguage] = tm.Description.(string)
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
		result[gDefaultLanguage] = tm.Title.(string)
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
	Title       interface{} `json:"title" yaml:"title"`     // title of the document, can be a single string, or a map[language-code:string]string
	Summary     interface{} `json:"summary" yaml:"summary"` // document summary, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon",yaml:"icon"`       // document's icon
	ContentFile interface{} `json:"file" yaml:"file"`       // name of document's content file, can be a single string, or a map[language-code:string]string
}

func (dm *DocumentMeta) setIndexAndId(input string) bool {
	if !RexpContentDir.MatchString(input) {
		return false
	}
	matches := RexpContentDir.FindStringSubmatch(input)
	dm.index, _ = strconv.Atoi(matches[1])
	dm.id = matches[2]
	return true
}

func (dm *DocumentMeta) GetContentFileNames() []string {
	switch dm.ContentFile.(type) {
	case string:
		return []string{dm.ContentFile.(string)}
	case map[string]string:
		result := make([]string, 0)
		for _, v := range dm.ContentFile.(map[string]string) {
			result = append(result, v)
		}
		return result
	default:
		return make([]string, 0)
	}
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
