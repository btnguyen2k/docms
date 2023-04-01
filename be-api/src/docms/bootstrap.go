package docms

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/btnguyen2k/docms/be-api/src/goapi"
	"github.com/btnguyen2k/docms/be-api/src/itineris"
	"github.com/labstack/echo/v4"
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
	goapi.PostInitEchoSetup = append(goapi.PostInitEchoSetup, postInitEchoSetup)
	initCMSData()
	initApiHandlers(goapi.ApiRouter)
	return nil
}

func postInitEchoSetup(e *echo.Echo) error {
	const confKeyFePath = "docms.frontend.path"
	fePath := goapi.AppConfig.GetString(confKeyFePath)
	const confKeyFeDir = "docms.frontend.directory"
	feDir := goapi.AppConfig.GetString(confKeyFeDir)
	const confKeyFeTemplate = "docms.frontend.template"
	feTemplate := goapi.AppConfig.GetString(confKeyFeTemplate)

	if fePath == "" || feDir == "" || feTemplate == "" {
		return fmt.Errorf("frontend path/directory/template is not defined at key [%s/%s/%s]", confKeyFePath, confKeyFeDir, confKeyFeTemplate)
	}

	// register handler for image files attached to documents
	e.GET("/img/:tid/:did/:img", serveImage)
	e.GET(fePath+"/:tid/:did/:img", serveImage)

	feTemplateDir := feDir + "/" + feTemplate
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
	gDataDir = goapi.AppConfig.GetString("docms.data_dir")
	log.Printf("[%s] Loading CMS data from <%s>...", logLevelInfo, gDataDir)

	var err error
	// load site's metadata
	gSiteMeta, err = LoadSiteMeta(gDataDir+"/"+metaFileYaml, gDataDir+"/"+metaFileJson)
	if err != nil {
		panic(err)
	}

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
		topicMeta, err := LoadTopicMeta(topicDirPath+"/"+metaFileYaml, topicDirPath+"/"+metaFileJson)
		if err != nil {
			panic(err)
		}
		topicMeta.setDirectory(topicDir.Name())
		gTopicList = append(gTopicList, topicMeta)
		gTopicMeta[topicMeta.id] = topicMeta
		gDocumentList[topicMeta.id] = make([]*DocumentMeta, 0)

		// fetch all documents inside the topic directory
		docDirList, err := GetDirContent(topicDirPath, func(entry os.DirEntry) bool {
			return entry.IsDir() && RexpContentDir.MatchString(entry.Name())
		})
		if err != nil {
			panic(err)
		}
		for _, docDir := range docDirList {
			docDirPath := gDataDir + "/" + topicDir.Name() + "/" + docDir.Name()
			log.Printf("[%s] Loading Document data from <%s>...", logLevelInfo, docDirPath)

			// load document metadata
			docMeta, err := LoadDocumentMeta(docDirPath+"/"+metaFileYaml, docDirPath+"/"+metaFileYaml)
			if err != nil {
				panic(err)
			}
			docMeta.setDirectory(docDir.Name())
			gDocumentList[topicMeta.id] = append(gDocumentList[topicMeta.id], docMeta)
			gDocumentMeta[topicMeta.id+":"+docMeta.id] = docMeta
			gDocumentContent[topicMeta.id+":"+docMeta.id] = make(map[string]string)

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
				gDocumentContent[topicMeta.id+":"+docMeta.id][k] = string(buff)
			}
		}
		sort.Slice(gDocumentList[topicMeta.id], func(i, j int) bool {
			return gDocumentList[topicMeta.id][i].index < gDocumentList[topicMeta.id][j].index
		})
	}
	sort.Slice(gTopicList, func(i, j int) bool {
		return gTopicList[i].index < gTopicList[j].index
	})
}

// Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)
//   - api-handler function must have the following signature:
//     func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
func initApiHandlers(router *itineris.ApiRouter) {
	router.SetHandler("getSiteMeta", apiGetSiteMeta)
	router.SetHandler("getTopics", apiGetTopics)
	router.SetHandler("getDocumentsForTopic", apiGetDocumentsForTopic)
	router.SetHandler("getDocument", apiGetDocument)
}
