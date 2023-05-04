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

	"github.com/blevesearch/bleve/v2"
	"github.com/btnguyen2k/consu/reddo"
	"gopkg.in/yaml.v3"
)

const (
	logLevelDebug   = "DEBUG"
	logLevelInfo    = "INFO"
	logLevelWarning = "WARN"
	logLevelError   = "ERROR"
)

var (
	DEBUG_MODE = os.Getenv("DEBUG") == "true"

	gDataDir              string
	gSiteMeta             *SiteMeta
	gTopicList            = make([]*TopicMeta, 0)              // list of topics, sorted by index
	gTopicMeta            = make(map[string]*TopicMeta)        // map[topic-id]topic-metadata
	gDocumentListPerTopic = make(map[string][]*DocumentMeta)   // list of documents, per topic, sorted by index
	gDocumentMeta         = make(map[string]*DocumentMeta)     // map["topic-id:document-id"]document-metadata
	gDocumentContent      = make(map[string]map[string]string) // map["topic-id:document-id"]map[language-code]document-content

	gFti bleve.Index // Full-text index

	gTagAlias     = make(map[string]map[string]string)   // map[language-code]map[alias]tag
	gDocumentTags = make(map[string]map[string][]string) // map[language-code]map[tag][]"topic-id:document-id"

	gSpecialPages = make(map[string][]string) // map[special-page-name][]"topic-id:document-id"
	gDocumentList = make([]string, 0)         // list of all documents (format: index/topic-id:doc-id), sorted by index
)

func _resetGlobalVars() {
	gTopicList = make([]*TopicMeta, 0)                       // list of topics, sorted by index
	gTopicMeta = make(map[string]*TopicMeta)                 // map[topic-id]topic-metadata
	gDocumentListPerTopic = make(map[string][]*DocumentMeta) // list of documents, per topic, sorted by index
	gDocumentMeta = make(map[string]*DocumentMeta)           // map["topic-id:document-id"]document-metadata
	gDocumentContent = make(map[string]map[string]string)    // map["topic-id:document-id"]map[language-code]document-content

	if gFti != nil {
		gFti.Close()
		gFti = nil
	}

	gTagAlias = make(map[string]map[string]string)       // map[language-code]map[alias]tag
	gDocumentTags = make(map[string]map[string][]string) // map[language-code]map[tag][]"topic-id:document-id"

	gSpecialPages = make(map[string][]string) // map[special-page-name][]"topic-id:document-id"
	gDocumentList = make([]string, 0)         // list of all documents (format: index/topic-id:doc-id), sorted by index
}

const (
	DefaultSiteMode  = SiteModeDocument
	SiteModeDoc      = "doc"
	SiteModeDocument = "document"
	SiteModeBlog     = "blog"
)

var (
	AllSiteModes = []string{SiteModeDoc, SiteModeDocument, SiteModeBlog}
)

// Author is site's/document's author
type Author struct {
	Name   string `json:"name" yaml:"name"`
	Email  string `json:"email" yaml:"email"`
	Avatar string `json:"avatar" yaml:"avatar"`
}

// SiteMeta captures metadata of the website.
type SiteMeta struct {
	FileInfo        os.FileInfo            `json:"-" yaml:"-"`                                   // internal use only!
	Name            string                 `json:"name" yaml:"name"`                             // name of the website
	Description     interface{}            `json:"description" yaml:"description"`               // short description, can be a single string, or a map[language-code:string]string
	Languages       map[string]string      `json:"languages" yaml:"languages"`                   // available languages of the website content
	DefaultLanguage string                 `json:"-" yaml:"-"`                                   // site's default language
	Icon            string                 `json:"icon" yaml:"icon"`                             // website's icon
	Contacts        map[string]string      `json:"contacts,omitempty" yaml:"contacts,omitempty"` // site's contact info
	Tags            map[string]interface{} `json:"tags,omitempty" yaml:"tags,omitempty"`         // site's tags
	TagsAlias       interface{}            `json:"tagalias,omitempty" yaml:"tagalias,omitempty"` // tags-alias, can be map[tag][]string or map[language-code]map[tag][]string
	Mode            string                 `json:"mode" yaml:"mode"`                             // site's mode, current support modes are: document/doc and blog
	Author          *Author                `json:"author,omitempty" yaml:"author,omitempty"`     // site's author (also default document's author)
}

var (
	_typMapAny    = reflect.TypeOf(map[string]interface{}{})
	_typMapString = reflect.TypeOf(map[string]string{})
)

func (sm *SiteMeta) init() error {
	// init field "default language"
	defLang := sm.Languages["default"]
	if defLang == "" {
		for lang := range sm.Languages {
			defLang = lang
			break
		}
	}
	sm.DefaultLanguage = defLang

	// verify "site-mode"
	ok := false
	for _, mode := range AllSiteModes {
		if mode == sm.Mode {
			ok = true
			break
		}
	}
	if !ok {
		sm.Mode = DefaultSiteMode
	}

	// normalize field "contacts"
	for k, v := range sm.Contacts {
		if v == "" {
			delete(sm.Contacts, k)
		}
	}

	// normalize field "tags"
	if sm.Tags == nil {
		sm.Tags = make(map[string]interface{})
	}

	return nil
}

func (sm *SiteMeta) toMap() map[string]interface{} {
	return map[string]interface{}{
		"name":            sm.Name,
		"languages":       sm.Languages,
		"defaultLanguage": sm.DefaultLanguage,
		"icon":            sm.Icon,
		"description":     sm.GetDescriptionMap(),
		"contacts":        sm.Contacts,
		"tags":            sm.Tags,
		"mode":            sm.Mode,
		"author":          sm.Author,
	}
}

func (sm *SiteMeta) GetDescriptionMap() map[string]string {
	desc := make(map[string]string)
	if sm.Description != nil {
		switch reflect.TypeOf(sm.Description).Kind() {
		case reflect.String:
			desc[sm.DefaultLanguage] = fmt.Sprintf("%s", sm.Description)
		case reflect.Map:
			temp, err := reddo.Convert(sm.Description, _typMapString)
			if err == nil && temp != nil {
				desc = temp.(map[string]string)
			}
		}
	}
	return desc
}

// GetTagAliasMap return the tags as a map[lang-code]map[tag][]alias
//
// Available since v0.3.0
func (sm *SiteMeta) GetTagAliasMap() map[string]map[string][]string {
	empty := make(map[string]map[string][]string)
	outer, err := reddo.Convert(sm.TagsAlias, _typMapAny)
	if err != nil || outer == nil {
		// the top level must be a map
		return empty
	}

	var nextLevelKind reflect.Kind
	for _, v := range outer.(map[string]interface{}) {
		if v != nil {
			nextLevelKind = reflect.TypeOf(v).Kind()
		}
		break
	}
	if nextLevelKind == reflect.Array || nextLevelKind == reflect.Slice {
		// then the next level must be either array/slice
		if result, err := reddo.Convert(sm.TagsAlias, reflect.TypeOf(make(map[string][]string))); err == nil && result != nil {
			return map[string]map[string][]string{
				sm.DefaultLanguage: result.(map[string][]string),
			}
		}
	} else if nextLevelKind == reflect.Map {
		// or a map
		if result, err := reddo.Convert(sm.TagsAlias, reflect.TypeOf(empty)); err == nil && result != nil {
			return result.(map[string]map[string][]string)
		}
	}

	return empty
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
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *SiteMeta
	err = yaml.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.FileInfo = fi
		metadata.init()
		if os.Getenv("CLI") == "true" {
			gSiteMeta = metadata
		}
	}
	return metadata, err
}

func LoadSiteMetaFromJson(filePath string) (*SiteMeta, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *SiteMeta
	err = json.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.FileInfo = fi
		metadata.init()
		if os.Getenv("CLI") == "true" {
			gSiteMeta = metadata
		}
	}
	return metadata, err
}

/*----------------------------------------------------------------------*/

// TopicMeta capture metadata of a topic.
type TopicMeta struct {
	FileInfo    os.FileInfo `json:"-" yaml:"-"`                                         // internal use only!
	index       int         `json:"-" yaml:"-"`                                         // topic index, for ordering
	id          string      `json:"-" yaml:"-"`                                         // topic id
	dir         string      `json:"-" yaml:"-"`                                         // name of directory where topic's data locates
	numDocs     int         `json:"-" yaml:"-"`                                         // number of documents in this topic
	Title       interface{} `json:"title" yaml:"title"`                                 // topic's title, can be a single string, or a map[language-code:string]string
	Description interface{} `json:"description,omitempty" yaml:"description,omitempty"` // short description, can be a single string, or a map[language-code:string]string
	Icon        string      `json:"icon,omitempty" yaml:"icon,omitempty"`               // topic's icon
	EntryImage  string      `json:"img,omitempty" yaml:"img,omitempty"`                 // topic's entry image
	Hidden      bool        `json:"hidden,omitempty" yaml:"hidden,omitempty"`           // if 'true', this topic is "hidden" from GUI
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

func (tm *TopicMeta) toMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          tm.id,
		"num_docs":    tm.numDocs,
		"icon":        tm.Icon,
		"title":       tm.GetTitleMap(),
		"description": tm.GetDescriptionMap(),
		"img":         tm.EntryImage,
		"hidden":      tm.Hidden,
	}
}

func (tm *TopicMeta) GetDescriptionMap() map[string]string {
	desc := make(map[string]string)
	if tm.Description != nil {
		switch reflect.TypeOf(tm.Description).Kind() {
		case reflect.String:
			desc[gSiteMeta.DefaultLanguage] = fmt.Sprintf("%s", tm.Description)
		case reflect.Map:
			temp, err := reddo.Convert(tm.Description, _typMapString)
			if err == nil && temp != nil {
				desc = temp.(map[string]string)
			}
		}
	}
	return desc
}

func (tm *TopicMeta) GetTitleMap() map[string]string {
	title := make(map[string]string)
	if tm.Title != nil {
		switch reflect.TypeOf(tm.Title).Kind() {
		case reflect.String:
			title[gSiteMeta.DefaultLanguage] = fmt.Sprintf("%s", tm.Title)
		case reflect.Map:
			temp, err := reddo.Convert(tm.Title, _typMapString)
			if err == nil && temp != nil {
				title = temp.(map[string]string)
			}
		}
	}
	return title
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
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *TopicMeta
	err = yaml.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.FileInfo = fi
	}
	return metadata, err
}

func LoadTopicMetaFromJson(filePath string) (*TopicMeta, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *TopicMeta
	err = json.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.FileInfo = fi
	}
	return metadata, err
}

/*----------------------------------------------------------------------*/

// DocumentMeta capture metadata of a document.
type DocumentMeta struct {
	FileInfo        os.FileInfo `json:"-" yaml:"-"`                                 // internal use only!
	index           int         `json:"-" yaml:"-"`                                 // document index, for ordering
	id              string      `json:"-" yaml:"-"`                                 // document id
	dir             string      `json:"-" yaml:"-"`                                 // name of directory where document's data locates
	Title           interface{} `json:"title" yaml:"title"`                         // title of the document, can be a single string, or a map[language-code:string]string
	Summary         interface{} `json:"summary,omitempty" yaml:"summary,omitempty"` // document summary, can be a single string, or a map[language-code:string]string
	Icon            string      `json:"icon,omitempty" yaml:"icon,omitempty"`       // document's icon
	ContentFile     interface{} `json:"file" yaml:"file"`                           // name of document's content file, can be a single string, or a map[language-code:string]string
	Tags            interface{} `json:"tags,omitempty" yaml:"tags,omitempty"`       // document's tags, can be []string or map[language-code][]string
	EntryImage      string      `json:"img,omitempty" yaml:"img,omitempty"`         // document's entry image
	DocPage         string      `json:"page,omitempty" yaml:"page,omitempty"`       // document plays as the special page on site (such as "contact" or "about")
	DocStyle        string      `json:"style,omitempty" yaml:"style,omitempty"`     // document's special style
	TimestampCreate int64       `json:"tc" yaml:"tc"`                               // UNIX timestamp when the document was created
	TimestampUpdate int64       `json:"tu" yaml:"tu"`                               // UNIX timestamp when the document was last updated
	Author          *Author     `json:"author,omitempty" yaml:"author,omitempty"`   // document's author
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

func (dm *DocumentMeta) toMap() map[string]interface{} {
	return map[string]interface{}{
		"id":      dm.id,
		"icon":    dm.Icon,
		"title":   dm.GetTitleMap(),
		"summary": dm.GetSummaryMap(),
		"tags":    dm.GetTagsMap(),
		"img":     dm.EntryImage,
		"page":    dm.DocPage,
		"style":   dm.DocStyle,
		"tc":      dm.TimestampCreate,
		"tu":      dm.TimestampUpdate,
		"author":  dm.Author,
	}
}

func (dm *DocumentMeta) GetSummaryMap() map[string]string {
	summary := make(map[string]string)
	if dm.Summary != nil {
		switch reflect.TypeOf(dm.Summary).Kind() {
		case reflect.String:
			summary[gSiteMeta.DefaultLanguage] = fmt.Sprintf("%s", dm.Summary)
		case reflect.Map:
			temp, err := reddo.Convert(dm.Summary, _typMapString)
			if err == nil && temp != nil {
				summary = temp.(map[string]string)
			}
		}
	}
	return summary
}

func (dm *DocumentMeta) GetTitleMap() map[string]string {
	title := make(map[string]string)
	if dm.Title != nil {
		switch reflect.TypeOf(dm.Title).Kind() {
		case reflect.String:
			title[gSiteMeta.DefaultLanguage] = fmt.Sprintf("%s", dm.Title)
		case reflect.Map:
			temp, err := reddo.Convert(dm.Title, _typMapString)
			if err == nil && temp != nil {
				title = temp.(map[string]string)
			}
		}
	}
	return title
}

func (dm *DocumentMeta) GetContentFileMap() map[string]string {
	files := make(map[string]string)
	if dm.ContentFile != nil {
		switch reflect.TypeOf(dm.ContentFile).Kind() {
		case reflect.String:
			files[gSiteMeta.DefaultLanguage] = fmt.Sprintf("%s", dm.ContentFile)
		case reflect.Map:
			temp, err := reddo.Convert(dm.ContentFile, _typMapString)
			if err == nil && temp != nil {
				files = temp.(map[string]string)
			}
		}
	}
	return files
}

// GetTagsMap return the tags as a map[lang-code][]tags
//
// Available since v0.3.0
func (dm *DocumentMeta) GetTagsMap() map[string][]string {
	tagsMap := make(map[string][]string)
	if dm.Tags != nil {
		switch reflect.TypeOf(dm.Tags).Kind() {
		case reflect.Array, reflect.Slice:
			temp, err := reddo.Convert(dm.Tags, reflect.TypeOf([]string{}))
			if err == nil && temp != nil {
				tagsMap[gSiteMeta.DefaultLanguage] = temp.([]string)
			}
		case reflect.Map:
			temp, err := reddo.Convert(dm.Tags, reflect.TypeOf(map[string][]string{}))
			if err == nil && temp != nil {
				tagsMap = temp.(map[string][]string)
			}
		}
	}
	for k := range tagsMap {
		sort.Strings(tagsMap[k])
	}
	return tagsMap
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
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *DocumentMeta
	err = yaml.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.FileInfo = fi
	}
	return metadata, err
}

func LoadDocumentMetaFromJson(filePath string) (*DocumentMeta, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *DocumentMeta
	err = json.Unmarshal(buf, &metadata)
	if err == nil {
		metadata.FileInfo = fi
	}
	return metadata, err
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
