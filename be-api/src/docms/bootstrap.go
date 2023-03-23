package docms

import (
	"fmt"
	"log"
	"os"
	"sort"

	"main/src/goapi"
	"main/src/itineris"
)

var Bootstrapper = &MyBootstrapper{name: "gvabe"}

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
	initCMSData()
	initApiHandlers(goapi.ApiRouter)
	return nil
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
