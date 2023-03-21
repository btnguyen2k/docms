package docms

import (
	"main/src/itineris"
)

var _apiResultGetSiteMeta *itineris.ApiResult

// API handler "getSiteMeta"
func apiGetSiteMeta(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	if _apiResultGetSiteMeta == nil {
		_apiResultGetSiteMeta = itineris.NewApiResult(itineris.StatusOk).SetData(map[string]interface{}{
			"name":        gSiteMeta.Name,
			"languages":   gSiteMeta.Languages,
			"icon":        gSiteMeta.Icon,
			"description": gSiteMeta.GetDescriptionMap(),
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
