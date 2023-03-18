package docms

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

// SiteMeta capture metadata of the website.
type SiteMeta struct {
	Name        string            `json:"name",yaml:"name"`               // name of the website
	Description interface{}       `json:"description",yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Languages   map[string]string `json:"languages",yaml:"languages"`     // available languages of the website content
	Icon        interface{}       `json:"icon",yaml:"icon"`               // website's main icon, can be a single string, or a map[string]string
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

// CategoryMeta capture metadata of a category.
type CategoryMeta struct {
	Title       interface{} `json:"title",yaml:"title"`             // title of the category, can be a single string, or a map[language-code:string]string
	Description interface{} `json:"description",yaml:"description"` // short description, can be a single string, or a map[language-code:string]string
	Icon        interface{} `json:"icon",yaml:"icon"`               // category icon, can be a single string, or a map[string]string
}

func LoadCategoryMetaFromYaml(filePath string) (*CategoryMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *CategoryMeta
	return metadata, yaml.Unmarshal(buf, &metadata)
}

func LoadCategoryMetaFromJson(filePath string) (*CategoryMeta, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var metadata *CategoryMeta
	return metadata, json.Unmarshal(buf, &metadata)
}

/*----------------------------------------------------------------------*/

// DocumentMeta capture metadata of a category.
type DocumentMeta struct {
	Title       interface{} `json:"title" yaml:"title"`     // title of the document, can be a single string, or a map[language-code:string]string
	Summary     interface{} `json:"summary" yaml:"summary"` // document summary, can be a single string, or a map[language-code:string]string
	Icon        interface{} `json:"icon" yaml:"icon"`       // document icon, can be a single string, or a map[string]string
	ContentFile interface{} `json:"file" yaml:"file"`       // name of document's content file, can be a single string, or a map[language-code:string]string
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
