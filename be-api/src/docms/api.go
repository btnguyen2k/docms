package docms

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/btnguyen2k/consu/reddo"
	"main/src/itineris"
)

var _apiResultGetSiteMeta *itineris.ApiResult

// API handler "getSiteMeta"
func apiGetSiteMeta(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	if _apiResultGetSiteMeta == nil {
		_apiResultGetSiteMeta = itineris.NewApiResult(itineris.StatusOk).SetData(map[string]interface{}{
			"name":            gSiteMeta.Name,
			"languages":       gSiteMeta.Languages,
			"defaultLanguage": gSiteMeta.DefaultLanguage,
			"icon":            gSiteMeta.Icon,
			"description":     gSiteMeta.GetDescriptionMap(),
			"contacts":        gSiteMeta.Contacts,
		})
	}
	return _apiResultGetSiteMeta
}

var _apiResultGetTopics *itineris.ApiResult

// API handler "getTopics"
func apiGetTopics(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	if _apiResultGetTopics == nil {
		topics := make([]map[string]interface{}, len(gTopicList))
		for i, topic := range gTopicList {
			topics[i] = map[string]interface{}{
				"id":          topic.id,
				"icon":        topic.Icon,
				"title":       topic.GetTitleMap(),
				"description": topic.GetDescriptionMap(),
			}
		}
		_apiResultGetTopics = itineris.NewApiResult(itineris.StatusOk).SetData(topics)
	}
	return _apiResultGetTopics
}

func _extractParam(params *itineris.ApiParams, paramName string, typ reflect.Type, defValue interface{}, regexp *regexp.Regexp) interface{} {
	v, _ := params.GetParamAsType(paramName, typ)
	if v == nil {
		v = defValue
	}
	if v != nil {
		if _, ok := v.(string); ok {
			v = strings.TrimSpace(v.(string))
			if regexp != nil && !regexp.Match([]byte(v.(string))) {
				return nil
			}
		}
	}
	return v
}

// API handler "getDocumentsForTopic"
func apiGetDocumentsForTopic(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	topicId := _extractParam(params, "tid", reddo.TypeString, "", nil)
	docMetaList := gDocumentList[topicId.(string)]
	if docMetaList == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage(fmt.Sprintf("Topic <%s> not found", topicId))
	}
	documents := make([]map[string]interface{}, len(docMetaList))
	for i, docMeta := range docMetaList {
		documents[i] = map[string]interface{}{
			"id":      docMeta.id,
			"icon":    docMeta.Icon,
			"title":   docMeta.GetTitleMap(),
			"summary": docMeta.GetSummaryMap(),
		}
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(documents)
}

// API handler "getDocument"
func apiGetDocument(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	topicId := _extractParam(params, "tid", reddo.TypeString, "", nil)
	docId := _extractParam(params, "did", reddo.TypeString, "", nil)
	docList := gDocumentList[topicId.(string)]
	if docList == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage(fmt.Sprintf("Topic <%s> not found", topicId))
	}
	docMeta := gDocumentMeta[topicId.(string)+":"+docId.(string)]
	if docMeta == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage(fmt.Sprintf("Document <%s/%s> not found", topicId, docId))
	}
	document := map[string]interface{}{
		"id":      docMeta.id,
		"icon":    docMeta.Icon,
		"title":   docMeta.GetTitleMap(),
		"summary": docMeta.GetSummaryMap(),
		"content": gDocumentContent[topicId.(string)+":"+docId.(string)],
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(document)
}
