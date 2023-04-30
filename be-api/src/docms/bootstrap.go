package docms

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/btnguyen2k/docms/be-api/src/goapi"
	"github.com/btnguyen2k/docms/be-api/src/itineris"
	"github.com/labstack/echo/v4"
	"github.com/radovskyb/watcher"
)

var Bootstrapper = &MyBootstrapper{name: "docms"}

// MyBootstrapper implements goapi.IBootstrapper
type MyBootstrapper struct {
	name string
}

// Bootstrap implements goapi.IBootstrapper.Bootstrap
//
// Bootstrapper usually does the following:
// - register api-handlers with the global ApiRouter
// - other initializing work (e.g. creating DAO, initializing database, etc)
func (m MyBootstrapper) Bootstrap() error {
	goapi.PostInitEchoSetup = append(goapi.PostInitEchoSetup, func(e *echo.Echo) error {
		if err := postInitEchoSetup(e); err != nil {
			panic(err)
		}
		return nil
	})
	initPublicData()
	initCMSData()
	initApiHandlers(goapi.ApiRouter)
	return nil
}

func initPublicData() {

}

func postInitEchoSetup(e *echo.Echo) error {
	const confKeyFePath = "docms.frontend.path"
	fePath := goapi.AppConfig.GetString(confKeyFePath)
	const confKeyFeDir = "docms.frontend.directory"
	feDir := goapi.AppConfig.GetString(confKeyFeDir)
	const confKeyFeTemplate = "docms.frontend.template"
	feTemplate := goapi.AppConfig.GetString(confKeyFeTemplate)

	if os.Getenv("DEBUG") != "true" {
		if fePath == "" || feDir == "" || feTemplate == "" {
			return fmt.Errorf("frontend path/directory/template is not defined at key [%s/%s/%s]", confKeyFePath, confKeyFeDir, confKeyFeTemplate)
		}

		var re = regexp.MustCompile(`^[0-9a-zA-Z_-]+$`)
		if !re.MatchString(feTemplate) {
			return fmt.Errorf("invalid frontend template: %s", feTemplate)
		}
	}

	feTemplateDir := feDir + "/" + feTemplate
	if fi, err := os.Stat(feTemplateDir); err != nil || !fi.IsDir() {
		return fmt.Errorf("invalid frontend template directory: %s", feTemplateDir)
	}

	// register handler for image files attached to documents
	e.GET("/img/:tid/:did/:img", serveImage)
	e.GET(fePath+"/:tid/:did/:img", serveImage)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, fePath+"/")
	})

	// map frontend's static assets
	dirContent, err := GetDirContent(feTemplateDir, nil)
	if err != nil {
		return err
	}
	for _, entry := range dirContent {
		if entry.IsDir() {
			e.Static(fePath+"/"+entry.Name()+"/*", feTemplateDir+"/"+entry.Name())
		} else {
			e.File(fePath+"/"+entry.Name(), feTemplateDir+"/"+entry.Name())
		}
	}

	// finally route everything else to "index.html:
	e.GET(fePath+"/*", func(c echo.Context) error {
		if fcontent, err := os.ReadFile(feTemplateDir + "/index.html"); err != nil {
			if os.IsNotExist(err) {
				return c.HTML(http.StatusNotFound, "Not found: "+c.Request().RequestURI)
			} else {
				return err
			}
		} else {
			return c.HTMLBlob(http.StatusOK, fcontent)
		}
	})

	return nil
}

var imgFileMime = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".svg":  "image/svg+xml",
}

var reFilename = regexp.MustCompile(`^[0-9a-zA-Z_\-\.]+$`)

func serveImage(c echo.Context) error {
	topicId := c.Param("tid")
	docId := c.Param("did")
	imgName := c.Param("img")
	topicMeta := gTopicMeta[topicId]
	docMeta := gDocumentMeta[topicId+":"+docId]
	if topicMeta == nil || docMeta == nil || !reFilename.MatchString(imgName) {
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, imgName))
	}

	ext := filepath.Ext(imgName)
	mimeType, ok := imgFileMime[ext]
	if !ok {
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, imgName))
	}

	fileName := gDataDir + "/" + topicMeta.dir + "/" + docMeta.dir + "/" + imgName
	buff, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("[%s] Error reading file [%s]: %s", logLevelError, fileName, err)
		return c.HTML(http.StatusNotFound, fmt.Sprintf("Not found: %s/%s/%s", topicId, docId, imgName))
	}

	return c.Blob(http.StatusOK, mimeType, buff)
}

func initCMSData() {
	if os.Getenv("DEBUG") == "true" {
		w := watcher.New()
		// w.SetMaxEvents(1)
		w.FilterOps(watcher.Create, watcher.Write, watcher.Move)
		go func() {
			for {
				select {
				case <-w.Event:
					_reloadCMSData()
				case err := <-w.Error:
					log.Fatalln(err)
				case <-w.Closed:
					return
				}
			}
		}()
		if err := w.AddRecursive(gDataDir); err != nil {
			log.Fatalln(err)
		}
		go w.Start(10 * time.Second)
	}
	_reloadCMSData()
}

func _loadSiteMeta() {
	var err error

	// load site's metadata
	if gSiteMeta, err = LoadSiteMetaAuto(gDataDir); err != nil {
		panic(err)
	}
	// normalize tag-alias
	for lang, inner := range gSiteMeta.GetTagAliasMap() {
		gTagAlias[lang] = make(map[string]string)
		for tag, aliasList := range inner {
			tag = strings.ToLower(strings.TrimSpace(tag))
			gTagAlias[lang][tag] = tag
			for _, alias := range aliasList {
				gTagAlias[lang][strings.ToLower(strings.TrimSpace(alias))] = tag
			}
		}
	}
	if os.Getenv("DEBUG") == "true" {
		log.Printf("[%s] site's tags: %#v", logLevelDebug, gSiteMeta.Tags)
		log.Printf("[%s] site's tag-alias: %#v", logLevelDebug, gTagAlias)
	}
}

func _loadDocumentsForTopic(topicMeta *TopicMeta) {
	// fetch all documents inside the topic directory
	topicDirPath := gDataDir + "/" + topicMeta.dir
	docDirList, err := GetDirContent(topicDirPath, func(entry os.DirEntry) bool {
		return entry.IsDir() && RexpContentDir.MatchString(entry.Name())
	})
	if err != nil {
		panic(err)
	}
	topicMeta.numDocs = len(docDirList)
	for _, docDir := range docDirList {
		docDirPath := gDataDir + "/" + topicMeta.dir + "/" + docDir.Name()
		log.Printf("[%s] Loading Document data from <%s>...", logLevelInfo, docDirPath)

		// load document metadata
		docMeta, err := LoadDocumentMetaAuto(docDirPath)
		if err != nil {
			panic(err)
		}
		{
			// temp fix
			fi, _ := os.Stat(docDirPath)
			if docMeta.TimestampUpdate <= 0 {
				docMeta.TimestampUpdate = fi.ModTime().Unix()
			}
			if docMeta.AuthorName == "" {
				docMeta.AuthorName = gSiteMeta.AuthorName
				if docMeta.AuthorName == "" {
					docMeta.AuthorName = goapi.AppConfig.GetString("app.shortname")
				}
			}
			if docMeta.AuthorEmail == "" {
				docMeta.AuthorEmail = gSiteMeta.AuthorEmail
				if docMeta.AuthorEmail == "" {
					docMeta.AuthorEmail = goapi.AppConfig.GetString("app.shortname") + "@domain.com"
				}
			}
		}
		docMeta.setDirectory(docDir.Name())
		topicDocId := topicMeta.id + ":" + docMeta.id
		gDocumentListPerTopic[topicMeta.id] = append(gDocumentListPerTopic[topicMeta.id], docMeta)
		gDocumentList = append(gDocumentList, fmt.Sprintf("%d", docMeta.index)+"/"+topicDocId)
		gDocumentMeta[topicDocId] = docMeta
		gDocumentContent[topicDocId] = make(map[string]string)
		if docMeta.DocPage != "" {
			specialPages := gSpecialPages[docMeta.DocPage]
			if specialPages == nil {
				specialPages = make([]string, 0)
				gSpecialPages[docMeta.DocPage] = specialPages
			}
			specialPages = append(specialPages, topicDocId)
		}

		if os.Getenv("DEBUG") == "true" {
			log.Printf("[%s] document's tags: %#v", logLevelDebug, docMeta.Tags)
			log.Printf("[%s] document's tags: %#v", logLevelDebug, docMeta.GetTagsMap())
		}

		// normalize document's tags
		docTagsMap := make(map[string][]string)
		for lang, docTags := range docMeta.GetTagsMap() {
			tagDocMap := gDocumentTags[lang]
			if tagDocMap == nil {
				tagDocMap = make(map[string][]string)
				gDocumentTags[lang] = tagDocMap
			}

			docTagsMap[lang] = make([]string, 0)
			for _, alias := range docTags {
				alias = strings.ToLower(strings.TrimSpace(alias))
				tag, ok := gTagAlias[lang][alias]
				if !ok {
					tag = alias
					gTagAlias[lang][alias] = tag
				}
				docTagsMap[lang] = append(docTagsMap[lang], tag)
				tagDocMap[tag] = append(tagDocMap[tag], topicDocId)
			}
			docTagsMap[lang] = removeDuplicateStrings(docTagsMap[lang])
		}
		for lang := range docTagsMap {
			docTagsMap[lang] = removeDuplicateStrings(docTagsMap[lang])
		}
		docMeta.Tags = docTagsMap

		// load document content
		docFileContentMap := docMeta.GetContentFileMap()
		if len(docFileContentMap) == 0 {
			panic(fmt.Errorf("document <%s> has no content defined at metadata key <file>", docDirPath))
		}
		for k, v := range docFileContentMap {
			buff, err := os.ReadFile(docDirPath + "/" + v)
			if err != nil {
				panic(err)
			}
			gDocumentContent[topicDocId][k] = string(buff)
		}
	}
	if gSiteMeta.Mode == SiteModeBlog {
		sort.Slice(gDocumentListPerTopic[topicMeta.id], func(i, j int) bool {
			return gDocumentListPerTopic[topicMeta.id][i].index > gDocumentListPerTopic[topicMeta.id][j].index
		})
	} else {
		sort.Slice(gDocumentListPerTopic[topicMeta.id], func(i, j int) bool {
			return gDocumentListPerTopic[topicMeta.id][i].index < gDocumentListPerTopic[topicMeta.id][j].index
		})
	}
}

func _loadTopics() {
	// fetch all topics
	topicDirList, err := GetDirContent(gDataDir, func(entry os.DirEntry) bool {
		return entry.IsDir() && RexpContentDir.MatchString(entry.Name())
	})
	if err != nil {
		panic(err)
	}
	for _, topicDir := range topicDirList {
		topicDirPath := gDataDir + "/" + topicDir.Name()
		log.Printf("[%s] Loading Topic data from <%s>...", logLevelInfo, topicDirPath)

		// load topic metadata
		topicMeta, err := LoadTopicMetaAuto(topicDirPath)
		if err != nil {
			panic(err)
		}
		topicMeta.setDirectory(topicDir.Name())
		gTopicList = append(gTopicList, topicMeta)
		gTopicMeta[topicMeta.id] = topicMeta
		gDocumentListPerTopic[topicMeta.id] = make([]*DocumentMeta, 0)

		_loadDocumentsForTopic(topicMeta)
	}
	sort.Slice(gTopicList, func(i, j int) bool {
		return gTopicList[i].index < gTopicList[j].index
	})
	if gSiteMeta.Mode == SiteModeBlog {
		sort.Slice(gDocumentList, func(i, j int) bool {
			return gDocumentList[i] > gDocumentList[j]
		})
	} else {
		gDocumentList = make([]string, 0)
	}
}

func _normalizeTags() {
	for lang, tagDocMap := range gDocumentTags {
		for tag := range tagDocMap {
			tagDocMap[tag] = removeDuplicateStrings(tagDocMap[tag])
		}
		gDocumentTags[lang] = tagDocMap
	}
}

func _loadFTI() {
	var err error

	// load fti if exists
	if gFti, err = bleve.OpenUsing(gDataDir+"/fti.bleve", map[string]interface{}{
		"read_only": true,
	}); err != nil {
		log.Printf("[%s] error while opening fulltext index: %s", logLevelError, err)
		gFti = nil
	}
}

var dataLock sync.Mutex

func _reloadCMSData() {
	if !dataLock.TryLock() {
		log.Printf("[%s] There is data lock, ignore loading...", logLevelWarning)
		return
	}
	defer dataLock.Unlock()
	gDataDir = goapi.AppConfig.GetString("docms.data_dir")
	log.Printf("[%s] Loading CMS data from <%s>...", logLevelInfo, gDataDir)

	_resetGlobalVars()
	_loadSiteMeta()
	_loadTopics()
	_normalizeTags()
	_loadFTI()
}

// Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)
//   - api-handler function must have the following signature:
//     func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
func initApiHandlers(router *itineris.ApiRouter) {
	router.SetHandler("getSiteMeta", apiGetSiteMeta)
	router.SetHandler("getTopics", apiGetTopics)
	router.SetHandler("getDocuments", apiGetDocuments)
	router.SetHandler("getDocumentsForTopic", apiGetDocumentsForTopic)
	router.SetHandler("getDocument", apiGetDocument)
	router.SetHandler("search", apiSearch)
	router.SetHandler("tagSearch", apiTagSearch)
	router.SetHandler("getTagCloud", apiGetTagCloud)

	router.SetHandler("health", apiHealthCheck)
	router.SetHandler("info", apiInfo)
}

var (
	apiResultOk = itineris.NewApiResult(itineris.StatusOk).SetMessage("ok")
)

func apiHealthCheck(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	return apiResultOk
}

func apiInfo(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	appData := make(map[string]interface{})
	appCfg := goapi.AppConfig.GetValue("app")
	if appCfg != nil {
		if appObj := appCfg.GetObject(); appObj != nil {
			keys := appObj.GetKeys()
			for _, k := range keys {
				appData[k] = appObj.GetKey(k).GetString()
			}
		}
	}
	tracking := map[string]interface{}{
		"gtag": goapi.AppConfig.GetString("docms.tracking.gtag"),
	}
	data := map[string]interface{}{
		"app":      appData,
		"tracking": tracking,
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	extras := map[string]interface{}{
		"memory": map[string]interface{}{
			"heap_alloc":          m.HeapAlloc,
			"heap_alloc_readable": bToReadable(m.HeapAlloc),
			"heap_sys":            m.HeapSys,
			"heap_sys_readable":   bToReadable(m.HeapSys),
			"sys":                 m.Sys,
			"sys_readable":        bToReadable(m.Sys),
		},
	}

	return itineris.NewApiResult(itineris.StatusOk).SetData(data).SetExtras(extras)
}

func bToReadable(b uint64) string {
	if b > 200_000_000 {
		return fmt.Sprintf("%.2f Gb", bToGb(b))
	}
	if b > 400_000 {
		return fmt.Sprintf("%.2f Mb", bToMb(b))
	}
	if b > 800 {
		return fmt.Sprintf("%.2f Kb", bToKb(b))
	}
	return fmt.Sprintf("%d bytes", b)
}

func bToKb(b uint64) float64 {
	return float64(b) / 1024
}

func bToMb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}

func bToGb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}
