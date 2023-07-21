package docms

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"golang.org/x/exp/maps"

	"github.com/btnguyen2k/consu/g18"
	"github.com/btnguyen2k/docms/be-api/src/goapi"
	"github.com/labstack/echo/v4"
)

var (
	crawlerBotNeedles = []string{
		"googlebot",
		"bingbot",
		"facebot",
		"twitterbot",
		"yandexbot",
		"duckduckbot",
		"slurp",
		"ia_archiver",
		"baiduspider",
		"teoma",
		"msnbot",
		"yandex",
		"bingpreview",
		"linkedinbot",
		"embedly",
		"quora link preview",
	}
)

type Template struct {
	tpl *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.tpl.ExecuteTemplate(w, name, data)
}

func setupCrawlerBotHandlers(e *echo.Echo) {
	e.Renderer = &Template{
		tpl: template.Must(template.New("").Funcs(map[string]interface{}{
			"unescapeHtml": func(s string) template.HTML {
				return template.HTML(s)
			},
		}).ParseGlob("./frontend/seo/*.html")),
	}
}

func isCrawlerBot(c echo.Context) bool {
	userAgent := strings.ToLower(c.Request().Header.Get("User-Agent"))
	for _, needle := range crawlerBotNeedles {
		if strings.Contains(userAgent, needle) {
			return true
		}
	}
	return false
}

func handleCrawlerBotRequest(fePath string, c echo.Context) error {
	uri := strings.TrimSpace(c.Request().RequestURI[len(fePath):])
	if uri == "/" || uri == "/index" || uri == "/index.html" {
		return c.Render(200, "index.html", seoModelIndex(fePath))
	}
	if !strings.HasSuffix(uri, "/") {
		return c.Redirect(301, fePath+uri+"/")
	}
	parts := strings.Split(strings.Trim(uri, "/"), "/")
	if len(parts) == 1 {
		topicId := parts[0]
		if topic := gTopicMeta[topicId]; topic != nil && !topic.Hidden {
			return c.Render(200, "topic.html", seoModelTopic(fePath, topic))
		}
	}
	if len(parts) == 2 {
		topicId := parts[0]
		docId := parts[1]
		if topic := gTopicMeta[topicId]; topic != nil && !topic.Hidden {
			if doc := gDocumentMeta[topicId+":"+docId]; doc != nil {
				return c.Render(200, "post.html", seoModelPost(fePath, topic, doc))
			}
		}
	}
	return c.JSON(404, map[string]interface{}{
		"code":  404,
		"error": "not found",
		"path":  uri,
	})
}

func seoModelIndex(fePath string) map[string]interface{} {
	topics := make([]map[string]interface{}, 0)
	for _, topic := range gTopicList {
		if !topic.Hidden || gSiteMeta.Mode != SiteModeBlog {
			topics = append(topics, _modelTopic(fePath, topic))
		}
	}

	posts := make([]map[string]interface{}, 0)
	if gSiteMeta.Mode == SiteModeBlog {
		topicIdList := make([]string, 0)
		for _, t := range gTopicList {
			topicIdList = append(topicIdList, t.id)
		}
		num := 10
		for _, indexTopicDocId := range gDocumentList {
			tokens := strings.Split(indexTopicDocId, "/")
			topicDocId := tokens[1]
			topicId := topicDocId[:strings.Index(topicDocId, ":")]
			if gDocumentMeta[topicDocId] != nil && !gTopicMeta[topicId].Hidden && g18.FindInSlice(topicId, topicIdList) >= 0 {
				doc := gDocumentMeta[topicDocId]
				posts = append(posts, _modelPost(fePath, doc, topicId))
				if len(posts) >= num {
					break
				}
			}
		}
	}

	return map[string]interface{}{
		"app": _modelApp(),
		"page": map[string]interface{}{
			"title":       gSiteMeta.Name,
			"description": gSiteMeta.GetDescriptionMap()[gSiteMeta.DefaultLanguage],
			"keywords":    "",
		},
		"topics": topics,
		"posts":  posts,
	}
}

func seoModelTopic(fePath string, topic *TopicMeta) map[string]interface{} {
	posts := make([]map[string]interface{}, 0)
	for _, doc := range gDocumentListPerTopic[topic.id] {
		posts = append(posts, _modelPost(fePath, doc, topic.id))
	}

	return map[string]interface{}{
		"app": _modelApp(),
		"page": map[string]interface{}{
			"title":       strings.Join(maps.Values(topic.GetTitleMap()), " - ") + " | " + gSiteMeta.Name,
			"description": strings.Join(maps.Values(topic.GetDescriptionMap()), " - "),
			"keywords":    "",
		},
		"breadcrumbs": []map[string]interface{}{
			{
				"title": gSiteMeta.Name,
				"url":   fePath + "/",
			},
		},
		"topic": _modelTopic(fePath, topic),
		"posts": posts,
	}
}

func seoModelPost(fePath string, topic *TopicMeta, doc *DocumentMeta) map[string]interface{} {
	tags := map[string]bool{}
	for _, tagList := range doc.GetTagsMap() {
		for _, tag := range tagList {
			tags[tag] = true
		}
	}

	postModel := _modelPost(fePath, doc, topic.id)
	contentMap := gDocumentContent[topic.id+":"+doc.id]
	mdParser := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	for lang, content := range contentMap {
		var buf bytes.Buffer
		if err := mdParser.Convert([]byte(content), &buf); err == nil {
			contentMap[lang] = buf.String()
		}
	}
	postModel["contents"] = contentMap

	return map[string]interface{}{
		"app": _modelApp(),
		"page": map[string]interface{}{
			"title":       strings.Join(maps.Values(doc.GetTitleMap()), " - ") + " | " + gSiteMeta.Name,
			"description": strings.Join(maps.Values(doc.GetSummaryMap()), " - "),
			"keywords":    strings.Join(maps.Keys(tags), ", "),
		},
		"breadcrumbs": []map[string]interface{}{
			{
				"title": gSiteMeta.Name,
				"url":   fePath + "/",
			},
			{
				"title": strings.Join(maps.Values(topic.GetTitleMap()), " - "),
				"url":   fePath + "/" + topic.id + "/",
			},
		},
		"topic":     _modelTopic(fePath, topic),
		"post":      postModel,
		"languages": maps.Keys(doc.GetContentFileMap()),
	}
}

func _modelApp() map[string]interface{} {
	return map[string]interface{}{
		"name":    goapi.AppConfig.GetString("app.name"),
		"version": goapi.AppConfig.GetString("app.version"),
	}
}

func _modelTopic(fePath string, topic *TopicMeta) map[string]interface{} {
	return map[string]interface{}{
		"id":           topic.id,
		"name":         topic.GetTitleMap()[gSiteMeta.DefaultLanguage],
		"names":        topic.GetTitleMap(),
		"description":  topic.GetDescriptionMap()[gSiteMeta.DefaultLanguage],
		"descriptions": topic.GetDescriptionMap(),
		"url":          fmt.Sprintf("%s/%s/", fePath, topic.id),
	}
}

func _modelPost(fePath string, doc *DocumentMeta, topicId string) map[string]interface{} {
	return map[string]interface{}{
		"id":        doc.id,
		"title":     doc.GetTitleMap()[gSiteMeta.DefaultLanguage],
		"titles":    doc.GetTitleMap(),
		"summary":   doc.GetSummaryMap()[gSiteMeta.DefaultLanguage],
		"summaries": doc.GetSummaryMap(),
		"url":       fmt.Sprintf("%s/%s/%s/", fePath, topicId, doc.id),
	}
}
