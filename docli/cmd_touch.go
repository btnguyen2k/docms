package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btnguyen2k/consu/g18"
	"github.com/btnguyen2k/docms/be-api/src/docms"
	"github.com/urfave/cli/v2"
)

var commandTouch = &cli.Command{
	Name:    "touch",
	Aliases: []string{"t"},
	Usage:   "Touch document metadata file to update timestamp",
	Flags: []cli.Flag{
		flagDir, flagDocTopic, flagDocId,
	},
	Action: actionTouchDocument,
}

// handle command "touch"
func actionTouchDocument(c *cli.Context) error {
	opts := OptsCmdTouch(c)

	// data dir
	if err := _validateDataDirMustExist(opts.DataDir); err != nil {
		return err
	}
	// siteMeta, err := docms.LoadSiteMetaAuto(opts.DataDir)
	// if err != nil {
	// 	return err
	// }

	// topic's id && directory
	topicId := opts.DocTopic
	if topicId == "" {
		return fmt.Errorf("id of topic must not be empty, supply topic's id with flag --%s", fieldTopic)
	}
	if !reId.MatchString(topicId) {
		return fmt.Errorf("invalid topic's id <%s>", topicId)
	}
	topicDir := ""
	docms.GetDirContent(opts.DataDir, func(entry os.DirEntry) bool {
		if entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name()) {
			if tokens := docms.RexpContentDir.FindStringSubmatch(entry.Name()); tokens[2] == topicId {
				topicDir = entry.Name()
			}
			return true
		}
		return false
	})
	if topicDir == "" {
		return fmt.Errorf("topic with id <%s> does not exist", topicId)
	}

	// document's id
	docId := opts.DocId
	if docId == "" {
		return fmt.Errorf("id of document must not be empty, supply topic's id with flag --%s", fieldId)
	}
	if !reId.MatchString(docId) {
		return fmt.Errorf("invalid document's id <%s>, should be lower cases and contain only letters and digits", docId)
	}

	// document's index
	var docDir *string = nil
	docms.GetDirContent(opts.DataDir+"/"+topicDir, func(entry os.DirEntry) bool {
		if entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name()) {
			if tokens := docms.RexpContentDir.FindStringSubmatch(entry.Name()); docId == tokens[2] {
				docDir = g18.PointerOf(entry.Name())
			}
			return true
		}
		return false
	})
	if docDir == nil {
		return fmt.Errorf("document with id <%s> does not exist", docId)
	}

	docMeta, err := docms.LoadDocumentMetaAuto(opts.DataDir + "/" + topicDir + "/" + *docDir)
	if err != nil {
		return err
	}
	docMeta.TimestampUpdate = now.UTC().Unix()

	metaFile := opts.DataDir + "/" + topicDir + "/" + *docDir + "/meta.yaml"
	if err := writeFileYaml(metaFile, docMeta); err != nil {
		return err
	}
	log.Printf("[INFO] document metadata has been updated at <%s>\n", metaFile)
	return nil
}
