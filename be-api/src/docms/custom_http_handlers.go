package docms

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/btnguyen2k/consu/checksum"
	"github.com/btnguyen2k/consu/reddo"
	"github.com/btnguyen2k/consu/semita"
	"github.com/btnguyen2k/docms/be-api/src/goapi"
	"github.com/btnguyen2k/docms/be-api/src/itineris"
	"github.com/gorilla/feeds"
	"github.com/labstack/echo/v4"
)

func _extractQueryParam(c echo.Context, paramName string, typ reflect.Type, defValue interface{}, regexp *regexp.Regexp) interface{} {
	v, _ := reddo.Convert(c.QueryParam(paramName), typ)
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

var mediaFileMime = make(map[string]string)

// var mediaFileMime = map[string]string{
// 	".jpg":  "image/jpeg",
// 	".jpeg": "image/jpeg",
// 	".png":  "image/png",
// 	".gif":  "image/gif",
// 	".svg":  "image/svg+xml",
// }

var reFilename = regexp.MustCompile(`^[0-9a-zA-Z_\-\.]+$`)

func serveMedia(c echo.Context) error {
	topicId := c.Param("tid")
	docId := c.Param("did")
	mediaName := c.Param("media")
	topicMeta := gTopicMeta[topicId]
	docMeta := gDocumentMeta[topicId+":"+docId]
	if topicMeta == nil || docMeta == nil || !reFilename.MatchString(mediaName) {
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, mediaName))
	}

	ext := filepath.Ext(mediaName)
	mimeType, ok := mediaFileMime[ext]
	if !ok {
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, mediaName))
	}

	fileName := gDataDir + "/" + topicMeta.dir + "/" + docMeta.dir + "/" + mediaName
	buff, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("[%s] Error reading file [%s]: %s", logLevelError, fileName, err)
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, mediaName))
	}
	fi, err := os.Stat(fileName)
	if err != nil {
		log.Printf("[%s] Error reading file info [%s]: %s", logLevelError, fileName, err)
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, mediaName))
	}

	c.Response().Header().Set("Cache-Control", "1024")
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000")      // cache for 1y , TODO: "public" part is configurable
	c.Response().Header().Set("Expires", time.Now().UTC().Format(time.RFC1123)) // expire in 1y
	c.Response().Header().Set("Date", fi.ModTime().UTC().Format(time.RFC1123))
	c.Response().Header().Set("Last-Modified", fi.ModTime().UTC().Format(time.RFC1123))
	c.Response().Header().Set("Content-Length", strconv.Itoa(len(buff)))
	c.Response().Header().Set("Etag", fmt.Sprintf("%x", checksum.Md5Checksum(buff)))
	return c.Blob(http.StatusOK, mimeType, buff)
}

func serveFeeds(c echo.Context) error {
	docs := make([]interface{}, 0)
	if gSiteMeta.Mode == SiteModeBlog {
		// only supported in "blog" mode
		topicIdList := make([]string, 0)
		for _, t := range gTopicList {
			topicIdList = append(topicIdList, t.id)
		}

		num := _extractQueryParam(c, "n", reddo.TypeInt, int64(10), nil)
		if num == nil || num.(int64) <= 0 || num.(int64) > 10 {
			num = int64(10)
		}

		docs = latestDocumentsForTopics(int(num.(int64)), topicIdList)
	}

	now := time.Now().UTC()
	siteUrl := strings.TrimSuffix(os.Getenv("SITE_URL"), "/")
	frontendPath := strings.TrimSuffix(goapi.AppConfig.GetString("docms.frontend.path"), "/")
	feed := &feeds.Feed{
		Title:       gSiteMeta.Name,
		Link:        &feeds.Link{Href: siteUrl},
		Description: gSiteMeta.GetDescriptionMap()[gSiteMeta.DefaultLanguage],
		Created:     now,
	}

	for _, doc := range docs {
		s := semita.NewSemita(doc)
		title, _ := s.GetValueOfType("title."+gSiteMeta.DefaultLanguage, reddo.TypeString)
		if title == nil {
			title = ""
		}
		summary, _ := s.GetValueOfType("summary."+gSiteMeta.DefaultLanguage, reddo.TypeString)
		if summary == nil {
			summary = ""
		}
		topicId, _ := s.GetValueOfType("topic.id", reddo.TypeString)
		if topicId == nil {
			topicId = ""
		}
		docId, _ := s.GetValueOfType("id", reddo.TypeString)
		if docId == nil {
			docId = ""
		}
		timestamp, _ := s.GetValueOfType("tc", reddo.TypeInt)
		if timestamp == nil {
			timestamp = now.Unix()
		}
		item := &feeds.Item{
			Title:       title.(string),
			Link:        &feeds.Link{Href: siteUrl + frontendPath + "/" + topicId.(string) + "/" + docId.(string) + "/"},
			Description: summary.(string),
			Created:     time.Unix(timestamp.(int64), 0).UTC(),
		}
		feed.Items = append(feed.Items, item)
	}

	typ := _extractQueryParam(c, "t", reddo.TypeString, "rss", nil)
	typ = strings.ToLower(typ.(string))
	if typ != "atom" && typ != "rss" && typ != "json" {
		typ = "rss"
	}

	switch typ.(string) {
	case "atom":
		feeds, err := feed.ToAtom()
		if err != nil {
			return c.String(itineris.StatusErrorServer, err.Error())
		}
		return c.Blob(itineris.StatusOk, "application/atom+xml", []byte(feeds))
	case "json":
		feeds, err := feed.ToJSON()
		if err != nil {
			return c.String(itineris.StatusErrorServer, err.Error())
		}
		return c.Blob(itineris.StatusOk, "application/json", []byte(feeds))
	default:
		json, err := feed.ToRss()
		if err != nil {
			return c.String(itineris.StatusErrorServer, err.Error())
		}
		return c.Blob(itineris.StatusOk, "application/rss+xml", []byte(json))
	}
}
