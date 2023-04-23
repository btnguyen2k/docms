package docms

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/btnguyen2k/consu/reddo"
	"gopkg.in/yaml.v3"
)

const (
	logLevelInfo    = "INFO"
	logLevelWarning = "WARN"
	logLevelError   = "ERROR"
)

var (
	gDataDir         string
	gSiteMeta        *SiteMeta
	gTopicList       = make([]*TopicMeta, 0)              // list of topics, sorted by index
	gTopicMeta       = make(map[string]*TopicMeta)        // map[topic-id]topic-metadata
	gDocumentList    = make(map[string][]*DocumentMeta)   // list of documents, per topic, sorted by index
	gDocumentMeta    = make(map[string]*DocumentMeta)     // map["topic-id:document-id"]document-metadata
	gDocumentContent = make(map[string]map[string]string) // map["topic-id:document-id"]map[language-code]document-content
	gFti             bleve.Index                          // Full-text index

	gTagAlias     = make(map[string]map[string]string)   // map[language-code]map[alias]tag
	gDocumentTags = make(map[string]map[string][]string) // map[language-code]map[tag][]"topic-id:document-id"
)

// SiteMeta capture metadata of the website.
type SiteMeta struct {
	Name            string                 `json:"name" yaml:"name"`               // name of the website
	Description     interface{}            `json:"description" yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Languages       map[string]string      `json:"languages" yaml:"languages"`     // available languages of the website content
	DefaultLanguage string                 `json:"-" yaml:"-"`                     // site's default language
	Icon            string                 `json:"icon" yaml:"icon"`               // website's icon
	Contacts        map[string]string      `json:"contacts" yaml:"contacts"`       // site's contact info
	Tags            map[string]interface{} `json:"tags" yaml:"tags"`               // site's tags
	TagsAlias       interface{}            `json:"tagalias" yaml:"tagalias"`       // tags-alias, can be map[tag][]string or map[language-code]map[tag][]string
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

func _checkNextLevelType(input map[string]interface{}) string {
	for _, v := range input {
		switch v.(type) {
		case []string:
			return "[]string"
		case []interface{}:
			return "[]interface{}"
		case map[string]interface{}:
			return "map[string]interface{}"
		}
	}
	return ""
}

var (
	_typMap = reflect.TypeOf(map[string]interface{}{})
)

// GetTagAliasMap return the tags as a map[lang-code]map[tag][]alias
//
// Available since v0.3.0
func (sm *SiteMeta) GetTagAliasMap() map[string]map[string][]string {
	outer, err := reddo.Convert(sm.TagsAlias, _typMap)
	if err != nil {
		// the top level must be a map
		return make(map[string]map[string][]string)
	}
	nextLevelType := _checkNextLevelType(outer.(map[string]interface{}))
	if strings.HasPrefix(nextLevelType, "[]") {
		// then the next level must be either []string
		if result, err := reddo.Convert(sm.TagsAlias, reflect.TypeOf(make(map[string][]string))); err == nil {
			return map[string]map[string][]string{
				gSiteMeta.DefaultLanguage: result.(map[string][]string),
			}
		}
	} else if strings.HasPrefix(nextLevelType, "map[") {
		// or a map[string][]string
		if result, err := reddo.Convert(sm.TagsAlias, reflect.TypeOf(make(map[string]map[string][]string))); err == nil {
			return result.(map[string]map[string][]string)
		}
	}
	return make(map[string]map[string][]string)
}

func LoadSiteMetaAuto(dir string) (*SiteMeta, error) {
	yamlFiles := []string{dir + "/meta.yaml", dir + "/meta.yml"}
	for _, yamlFilePath := range yamlFiles {
		siteMeta, err := LoadSiteMetaFromYaml(yamlFilePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return siteMeta, nil
	}

	jsonFiles := []string{dir + "/meta.json"}
	for _, jsonFilePath := range jsonFiles {
		siteMeta, err := LoadSiteMetaFromJson(jsonFilePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return siteMeta, nil
	}

	return nil, fmt.Errorf("no meta file found")
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

// TopicMeta capture metadata of a topic.
type TopicMeta struct {
	index       int         `json:"-" yaml:"-"`                     // topic index, for ordering
	id          string      `json:"-" yaml:"-"`                     // topic id
	dir         string      `json:"-" yaml:"-"`                     // name of directory where topic's data locates
	Title       interface{} `json:"title" yaml:"title"`             // topic's title, can be a single string, or a map[language-code:string]string
	Description interface{} `json:"description" yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon" yaml:"icon"`               // topic's icon
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

func LoadTopicMetaAuto(dir string) (*TopicMeta, error) {
	yamlFiles := []string{dir + "/meta.yaml", dir + "/meta.yml"}
	for _, yamlFilePath := range yamlFiles {
		topicMeta, err := LoadTopicMetaFromYaml(yamlFilePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return topicMeta, nil
	}

	jsonFiles := []string{dir + "/meta.json"}
	for _, jsonFilePath := range jsonFiles {
		topicMeta, err := LoadTopicMetaFromJson(jsonFilePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return topicMeta, nil
	}

	return nil, fmt.Errorf("no meta file found")
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

// DocumentMeta capture metadata of a document.
type DocumentMeta struct {
	index       int         `json:"-" yaml:"-"`             // document index, for ordering
	id          string      `json:"-" yaml:"-"`             // document id
	dir         string      `json:"-" yaml:"-"`             // name of directory where document's data locates
	Title       interface{} `json:"title" yaml:"title"`     // title of the document, can be a single string, or a map[language-code:string]string
	Summary     interface{} `json:"summary" yaml:"summary"` // document summary, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon" yaml:"icon"`       // document's icon
	ContentFile interface{} `json:"file" yaml:"file"`       // name of document's content file, can be a single string, or a map[language-code:string]string
	Tags        interface{} `json:"tags" yaml:"tags"`       // document's tags, can be []string or map[language-code][]string
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

// GetTagsMap return the tags as a map[lang-code][]tags
//
// Available since v0.3.0
func (dm *DocumentMeta) GetTagsMap() map[string][]string {
	result := make(map[string][]string)
	switch dm.Tags.(type) {
	case []string:
		result[gSiteMeta.DefaultLanguage] = dm.Tags.([]string)
	case []interface{}:
		temp, err := reddo.Convert(dm.Tags, reflect.TypeOf([]string{}))
		if err == nil && temp != nil {
			result[gSiteMeta.DefaultLanguage] = temp.([]string)
		}
	case map[string]interface{}, map[string][]interface{}, map[string][]string:
		temp, err := reddo.Convert(dm.Tags, reflect.TypeOf(map[string][]string{}))
		if err == nil && temp != nil {
			result = temp.(map[string][]string)
		}
	}
	for k := range result {
		sort.Strings(result[k])
	}
	return result
}

func LoadDocumentMetaAuto(dir string) (*DocumentMeta, error) {
	yamlFiles := []string{dir + "/meta.yaml", dir + "/meta.yml"}
	for _, yamlFilePath := range yamlFiles {
		documentMeta, err := LoadDocumentMetaFromYaml(yamlFilePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return documentMeta, nil
	}

	jsonFiles := []string{dir + "/meta.json"}
	for _, jsonFilePath := range jsonFiles {
		documentMeta, err := LoadDocumentMetaFromJson(jsonFilePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return documentMeta, nil
	}

	return nil, fmt.Errorf("no meta file found")
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

var RexpContentDir = regexp.MustCompile(`^(\d+)-(\w+)$`)

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

func removeDuplicateStrings(s []string) []string {
	if len(s) < 1 {
		return s
	}
	sort.Strings(s)
	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			s[prev] = s[curr]
			prev++
		}
	}
	return s[:prev]
}
