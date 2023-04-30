package docms

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/btnguyen2k/consu/g18"
	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/docms/be-api/src/itineris"
)

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

var _apiResultGetSiteMeta *itineris.ApiResult

// API handler "getSiteMeta"
func apiGetSiteMeta(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	if _apiResultGetSiteMeta == nil {
		_apiResultGetSiteMeta = itineris.NewApiResult(itineris.StatusOk).SetData(gSiteMeta.toMap())
	}
	return _apiResultGetSiteMeta
}

var _apiResultGetTopics *itineris.ApiResult

// API handler "getTopics"
func apiGetTopics(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	if _apiResultGetTopics == nil {
		topics := make([]map[string]interface{}, len(gTopicList))
		for i, topic := range gTopicList {
			if !topic.Hidden {
				topics[i] = topic.toMap()
			}
		}
		_apiResultGetTopics = itineris.NewApiResult(itineris.StatusOk).SetData(topics)
	}
	return _apiResultGetTopics
}

var _apiEmptyResultGetDocuments = itineris.NewApiResult(itineris.StatusOk).SetData(make([]interface{}, 0))

func latestDocumentsForTopics(num int, topicIdList []string) []interface{} {
	result := make([]interface{}, 0)
	for len(topicIdList) == 0 {
		return result
	}
	for _, indexTopicDocId := range gDocumentList {
		tokens := strings.Split(indexTopicDocId, "/")
		topicDocId := tokens[1]
		topicId := topicDocId[:strings.Index(topicDocId, ":")]
		if gDocumentMeta[topicDocId] != nil && !gTopicMeta[topicId].Hidden && g18.FindInSlice(topicId, topicIdList) >= 0 {
			docData := gDocumentMeta[topicDocId].toMap()
			docData["topic"] = gTopicMeta[topicId].toMap()
			result = append(result, docData)
			if len(result) >= num {
				break
			}
		}
	}
	return result
}

// API handler "getDocuments"
func apiGetDocuments(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	purpose := _extractParam(params, "p", reddo.TypeString, "", nil)
	if purpose == "latest" {
		// fetching "latest" documents is supported only in blog mode
		if gSiteMeta.Mode != SiteModeBlog {
			return _apiEmptyResultGetDocuments
		}
		num := _extractParam(params, "n", reddo.TypeInt, 10, nil)
		if num == nil || num.(int64) <= 0 || num.(int64) > 10 {
			num = 10
		}
		topics := _extractParam(params, "t", reddo.TypeString, "", nil)
		if topics == nil {
			topics = ""
		}
		topicIdList := regexp.MustCompile(`[\s,;:]+`).Split(topics.(string), -1)
		// fmt.Printf("[DEBUG] num: %#v / param: %#v / topics: %#v\n", num, topics, topicIdList)
		if topics.(string) == "" || topicIdList == nil || len(topicIdList) == 0 {
			topicIdList = make([]string, 0)
			for _, t := range gTopicList {
				topicIdList = append(topicIdList, t.id)
			}
		}
		// fmt.Printf("[DEBUG] num: %#v / param: %#v / topics: %#v / gTopicList: %#v\n", num, topics, topicIdList, gTopicList)
		docs := latestDocumentsForTopics(int(num.(int64)), topicIdList)
		return itineris.NewApiResult(itineris.StatusOk).SetData(docs)
	}
	return _apiEmptyResultGetDocuments
}

// API handler "getDocumentsForTopic"
func apiGetDocumentsForTopic(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	topicId := _extractParam(params, "tid", reddo.TypeString, "", nil)
	docMetaList := gDocumentListPerTopic[topicId.(string)]
	if docMetaList == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage(fmt.Sprintf("Topic <%s> not found", topicId))
	}
	documents := make([]map[string]interface{}, len(docMetaList))
	for i, docMeta := range docMetaList {
		documents[i] = docMeta.toMap()
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(documents)
}

// API handler "getDocument"
func apiGetDocument(_ *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
	topicId := _extractParam(params, "tid", reddo.TypeString, "", nil)
	docId := _extractParam(params, "did", reddo.TypeString, "", nil)
	docList := gDocumentListPerTopic[topicId.(string)]
	if docList == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage(fmt.Sprintf("Topic <%s> not found", topicId))
	}
	docMeta := gDocumentMeta[topicId.(string)+":"+docId.(string)]
	if docMeta == nil {
		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage(fmt.Sprintf("Document <%s/%s> not found", topicId, docId))
	}
	document := docMeta.toMap()
	document["content"] = gDocumentContent[topicId.(string)+":"+docId.(string)]
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
		// hit.ID is concatenation of "topic-Id:doc-id:lang-code"
		topicDocId := hit.ID[:strings.LastIndex(hit.ID, ":")]
		if gDocumentMeta[topicDocId] != nil {
			topicId := hit.ID[:strings.Index(hit.ID, ":")]
			docData := gDocumentMeta[topicDocId].toMap()
			docData["topic"] = gTopicMeta[topicId].toMap()
			docs = append(docs, docData)
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
			docData := gDocumentMeta[topicDocId].toMap()
			docData["topic"] = gTopicMeta[topicId].toMap()
			docs = append(docs, docData)
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
