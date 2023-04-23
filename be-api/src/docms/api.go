package docms

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/docms/be-api/src/itineris"
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
			"tags":            gSiteMeta.Tags,
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
			"tags":    docMeta.GetTagsMap(),
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
		"tags":    docMeta.GetTagsMap(),
		"content": gDocumentContent[topicId.(string)+":"+docId.(string)],
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(document)
}

var apiResultNoSearchResult = itineris.NewApiResult(itineris.StatusOk).SetData(map[string]interface{}{"total": 0, "duration": 0, "docs": make([]interface{}, 0)})
var apiResultFtiNotAvailable = itineris.NewApiResult(itineris.StatusNotImplemented).SetMessage("fulltext index not available")
var apiResultInvalidSearchQuery = itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("invalid search query")
var reLocale = regexp.MustCompile(`^[\w-]+$`)

// API handler "search"
func apiSearch(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	if gFti == nil {
		return apiResultFtiNotAvailable
	}
	query, err := params.GetParamAsType("query", reddo.TypeString)
	if query == nil || err != nil || strings.TrimSpace(query.(string)) == "" || len(strings.TrimSpace(query.(string))) > 100 {
		// return apiResultInvalidSearchQuery
		return apiResultNoSearchResult
	}

	clientLocale := ctx.GetClientLocale()
	if !reLocale.MatchString(clientLocale) {
		clientLocale = gSiteMeta.DefaultLanguage
	}
	searchQuery := bleve.NewBooleanQuery()
	searchQuery.AddMust(
		bleve.NewQueryStringQuery("lang:"+clientLocale),
		bleve.NewMatchQuery(query.(string)),
	)
	search := bleve.NewSearchRequest(searchQuery)
	searchResults, err := gFti.Search(search)
	if err != nil {
		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
	}

	result := map[string]interface{}{
		"total":    searchResults.Total,
		"duration": searchResults.Took.Seconds(),
	}
	docs := make([]map[string]interface{}, 0)
	for _, hit := range searchResults.Hits {
		// hit.ID is concatation of "topic-Id:doc-id:lang-code"
		topicDocId := hit.ID[:strings.LastIndex(hit.ID, ":")]
		if gDocumentMeta[topicDocId] != nil {
			topicId := hit.ID[:strings.Index(hit.ID, ":")]
			docs = append(docs, map[string]interface{}{
				"id":      gDocumentMeta[topicDocId].id,
				"icon":    gDocumentMeta[topicDocId].Icon,
				"title":   gDocumentMeta[topicDocId].GetTitleMap(),
				"summary": gDocumentMeta[topicDocId].GetSummaryMap(),
				"tags":    gDocumentMeta[topicDocId].GetTagsMap(),
				"topic": map[string]interface{}{
					"id":          topicId,
					"icon":        gTopicMeta[topicId].Icon,
					"title":       gTopicMeta[topicId].GetTitleMap(),
					"description": gTopicMeta[topicId].GetDescriptionMap(),
				},
			})
			if len(docs) >= 10 {
				break
			}
		}
	}
	result["docs"] = docs
	return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}

// API handler "tagSearch"
func apiTagSearch(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	query, err := params.GetParamAsType("query", reddo.TypeString)
	if query == nil || err != nil || strings.TrimSpace(query.(string)) == "" || len(strings.TrimSpace(query.(string))) > 100 {
		// return apiResultInvalidSearchQuery
		return apiResultNoSearchResult
	}

	start := time.Now()
	clientLocale := ctx.GetClientLocale()
	if !reLocale.MatchString(clientLocale) {
		clientLocale = gSiteMeta.DefaultLanguage
	}
	alias := strings.ToLower(strings.TrimSpace(query.(string)))
	tag := ""
	if tagList := gTagAlias[clientLocale]; tagList != nil {
		tag = gTagAlias[clientLocale][alias]
	}
	docIdList := gDocumentTags[clientLocale][tag]
	if docIdList == nil || len(docIdList) == 0 {
		return apiResultNoSearchResult
	}
	result := map[string]interface{}{
		"total":    len(docIdList),
		"duration": time.Now().Sub(start).Seconds(),
	}
	docs := make([]map[string]interface{}, 0)
	for _, topicDocId := range docIdList {
		if gDocumentMeta[topicDocId] != nil {
			topicId := topicDocId[:strings.Index(topicDocId, ":")]
			docs = append(docs, map[string]interface{}{
				"id":      gDocumentMeta[topicDocId].id,
				"icon":    gDocumentMeta[topicDocId].Icon,
				"title":   gDocumentMeta[topicDocId].GetTitleMap(),
				"summary": gDocumentMeta[topicDocId].GetSummaryMap(),
				"tags":    gDocumentMeta[topicDocId].GetTagsMap(),
				"topic": map[string]interface{}{
					"id":          topicId,
					"icon":        gTopicMeta[topicId].Icon,
					"title":       gTopicMeta[topicId].GetTitleMap(),
					"description": gTopicMeta[topicId].GetDescriptionMap(),
				},
			})
			if len(docs) >= 10 {
				break
			}
		}
	}
	result["docs"] = docs
	return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}

// API handler "getTagCloud"
func apiGetTagCloud(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	return itineris.NewApiResult(itineris.StatusOk).SetData(gDocumentTags)
}
