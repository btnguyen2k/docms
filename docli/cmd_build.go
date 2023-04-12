package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/btnguyen2k/docms/be-api/src/docms"
	"github.com/urfave/cli/v2"
)

var commandBuild = &cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "Preprocess website content, ready for DO CMS runtime",
	Flags: []cli.Flag{
		flagSrc,
		flagOutput,
		flagPurge,
	},
	Action: actionBuild,
}

func _loadSiteMetadata(opts *Options) (*docms.SiteMeta, error) {
	for _, metaFileYaml := range []string{opts.SrcDir + "/meta.yaml", opts.SrcDir + "/meta.yml"} {
		log.Printf("[INFO] looking for file <%s>...\n", metaFileYaml)
		if isFile(metaFileYaml) {
			return docms.LoadSiteMetaFromYaml(metaFileYaml)
		}
	}

	for _, metaFileJson := range []string{opts.SrcDir + "/meta.json"} {
		log.Printf("[INFO] looking for file <%s>...\n", metaFileJson)
		if isFile(metaFileJson) {
			return docms.LoadSiteMetaFromJson(metaFileJson)
		}
	}

	return nil, fmt.Errorf("no metadata file found")
}

func _verifySiteMetadata(siteMeta *docms.SiteMeta) (*docms.SiteMeta, bool) {
	checkPass := true
	newMetadata := &docms.SiteMeta{}
	log.Printf("[INFO] veryfing metadata file...\n")

	// "name" must not empty
	{
		if strings.TrimSpace(siteMeta.Name) == "" {
			log.Printf("[ERROR] {name} must not be empty\n")
			checkPass = false
		}
		newMetadata.Name = strings.TrimSpace(siteMeta.Name)
	}

	// "languages" if any, must me a map[string]string
	if len(siteMeta.Languages) == 0 {
		log.Printf("[WARN] cannot parse {languages} config or it does not exist, falling back to English...\n")
		newMetadata.Languages = map[string]string{"en": "English"}
	} else {
		newMetadata.Languages = siteMeta.Languages
	}

	// "description" must be a string, or a map[string]string
	{
		desc := siteMeta.Description
		switch desc.(type) {
		case string:
			if strings.TrimSpace(desc.(string)) == "" {
				log.Printf("[ERROR] {description} must not be empty\n")
				checkPass = false
			} else {
				desc = strings.TrimSpace(desc.(string))
			}
		case map[string]interface{}:
			if len(desc.(map[string]interface{})) == 0 {
				log.Printf("[ERROR] {description} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range desc.(map[string]interface{}) {
				if _, ok := newMetadata.Languages[k]; !ok {
					log.Printf("[WARN] language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			desc = temp
		default:
			log.Printf("[ERROR] cannot parse {description} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Description = desc
	}

	// icon
	newMetadata.Icon = siteMeta.Icon

	// contacts
	newMetadata.Contacts = siteMeta.Contacts

	// tags
	{
		now := time.Now()
		strDate := now.Format("20060102")
		strTime := now.Format("150405")
		strDatetime := now.Format("20060102T150405")
		newMetadata.Tags = make(map[string]string)
		for k, v := range siteMeta.Tags {
			newMetadata.Tags[k] = strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(v, "${build_date}", strDate), "${build_time}", strTime,
				), "${build_datetime}", strDatetime,
			)
		}
	}

	return newMetadata, checkPass
}

func _loadTopicMetadata(opts *Options, dir os.DirEntry) (*docms.TopicMeta, error) {
	for _, metaFileYaml := range []string{opts.SrcDir + "/" + dir.Name() + "/meta.yaml", opts.SrcDir + "/" + dir.Name() + "/meta.yml"} {
		log.Printf("[INFO]\t looking for file <%s>...\n", metaFileYaml)
		if isFile(metaFileYaml) {
			return docms.LoadTopicMetaFromYaml(metaFileYaml)
		}
	}

	for _, metaFileJson := range []string{opts.SrcDir + "/" + dir.Name() + "/meta.json"} {
		log.Printf("[INFO]\t looking for file <%s>...\n", metaFileJson)
		if isFile(metaFileJson) {
			return docms.LoadTopicMetaFromJson(metaFileJson)
		}
	}

	return nil, fmt.Errorf("no metadata file found")
}

func _verifyTopicMetadata(siteMeta *docms.SiteMeta, topicMeta *docms.TopicMeta) (*docms.TopicMeta, bool) {
	checkPass := true
	newMetadata := &docms.TopicMeta{}
	log.Printf("[INFO]\t veryfing metadata file...\n")

	// "title" must be a string, or a map[string]string
	{
		title := topicMeta.Title
		switch title.(type) {
		case string:
			if strings.TrimSpace(title.(string)) == "" {
				log.Printf("[ERROR]\t {title} must not be empty\n")
				checkPass = false
			} else {
				title = strings.TrimSpace(title.(string))
			}
		case map[string]interface{}:
			if len(title.(map[string]interface{})) == 0 {
				log.Printf("[ERROR]\t {title} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range title.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					log.Printf("[WARN]\t language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			title = temp
		default:
			log.Printf("[ERROR]\t cannot parse {title} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Title = title
	}

	// "description" must be a string, or a map[string]string
	{
		desc := topicMeta.Description
		switch desc.(type) {
		case string:
			if strings.TrimSpace(desc.(string)) == "" {
				log.Printf("[ERROR]\t {description} must not be empty\n")
				checkPass = false
			} else {
				desc = strings.TrimSpace(desc.(string))
			}
		case map[string]interface{}:
			if len(desc.(map[string]interface{})) == 0 {
				log.Printf("[ERROR]\t {description} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range desc.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					log.Printf("[WARN]\t language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			desc = temp
		default:
			log.Printf("[ERROR]\t cannot parse {description} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Description = desc
	}

	// icon
	newMetadata.Icon = topicMeta.Icon

	return newMetadata, checkPass
}

func _buildTopicDir(opts *Options, siteMeta *docms.SiteMeta, topicDir os.DirEntry, fti bleve.Index) error {
	log.Printf("[INFO] building topic from <%s>...\n", opts.SrcDir+"/"+topicDir.Name())

	topicMeta, err := _loadTopicMetadata(opts, topicDir)
	if err != nil {
		return err
	}
	newTopicMeta, ok := _verifyTopicMetadata(siteMeta, topicMeta)
	if !ok {
		return fmt.Errorf("there is error while checking metadata file")
	}
	log.Printf("[INFO]\t metadata verification done.\n")
	if err := os.Mkdir(opts.OutputDir+"/"+topicDir.Name(), dirPerm); err != nil {
		return err
	}
	if err := writeFileYaml(opts.OutputDir+"/"+topicDir.Name()+"/meta.yaml", newTopicMeta); err != nil {
		return err
	}

	srcDirEntries, err := docms.GetDirContent(opts.SrcDir+"/"+topicDir.Name(), func(entry os.DirEntry) bool {
		return entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name())
	})
	if err != nil {
		return err
	}
	idmap := map[string]bool{}
	for _, dirEntry := range srcDirEntries {
		matches := docms.RexpContentDir.FindStringSubmatch(dirEntry.Name())
		if _, ok := idmap[matches[2]]; ok {
			return fmt.Errorf("duplicated document-id at directory <%s>", topicDir.Name()+"/"+dirEntry.Name())
		}
		idmap[matches[2]] = true
		if err := _buildDocDir(opts, siteMeta, topicDir, dirEntry, fti); err != nil {
			return err
		}
	}
	return nil
}

func _loadDocumentMetadata(opts *Options, topicDir, docDir os.DirEntry) (*docms.DocumentMeta, error) {
	for _, metaFileYaml := range []string{opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/meta.yaml", opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/meta.yml"} {
		log.Printf("[INFO]\t\t looking for file <%s>...\n", metaFileYaml)
		if isFile(metaFileYaml) {
			return docms.LoadDocumentMetaFromYaml(metaFileYaml)
		}
	}

	for _, metaFileJson := range []string{opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/meta.json"} {
		log.Printf("[INFO]\t\t looking for file <%s>...\n", metaFileJson)
		if isFile(metaFileJson) {
			return docms.LoadDocumentMetaFromJson(metaFileJson)
		}
	}

	return nil, fmt.Errorf("no metadata file found")
}

func _verifyDocumentMetadata(siteMeta *docms.SiteMeta, docMeta *docms.DocumentMeta) (*docms.DocumentMeta, bool) {
	checkPass := true
	newMetadata := &docms.DocumentMeta{}
	log.Printf("[INFO]\t\t veryfing metadata file...\n")

	// "title" must be a string, or a map[string]string
	{
		title := docMeta.Title
		switch title.(type) {
		case string:
			if strings.TrimSpace(title.(string)) == "" {
				log.Printf("[ERROR]\t\t {title} must not be empty\n")
				checkPass = false
			} else {
				title = strings.TrimSpace(title.(string))
			}
		case map[string]interface{}:
			if len(title.(map[string]interface{})) == 0 {
				log.Printf("[ERROR]\t\t {title} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range title.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					log.Printf("[WARN]\t\t language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			title = temp
		default:
			log.Printf("[ERROR]\t\t cannot parse {title} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Title = title
	}

	// "summary" must be a string, or a map[string]string
	{
		summary := docMeta.Summary
		switch summary.(type) {
		case string:
			if strings.TrimSpace(summary.(string)) == "" {
				log.Printf("[ERROR]\t\t {summary} must not be empty\n")
				checkPass = false
			} else {
				summary = strings.TrimSpace(summary.(string))
			}
		case map[string]interface{}:
			if len(summary.(map[string]interface{})) == 0 {
				log.Printf("[ERROR]\t\t {summary} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range summary.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					log.Printf("[WARN]\t\t language <%s> defined in {summary} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			summary = temp
		default:
			log.Printf("[ERROR]\t\t cannot parse {summary} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Summary = summary
	}

	// icon
	newMetadata.Icon = docMeta.Icon

	// "content file" must be a string, or a map[string]string
	{
		contentFile := docMeta.ContentFile
		switch contentFile.(type) {
		case string:
			if strings.TrimSpace(contentFile.(string)) == "" {
				log.Printf("[ERROR]\t\t {file} must not be empty\n")
				checkPass = false
			} else {
				contentFile = strings.TrimSpace(contentFile.(string))
			}
		case map[string]interface{}:
			if len(contentFile.(map[string]interface{})) == 0 {
				log.Printf("[ERROR]\t\t {file} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range contentFile.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					log.Printf("[WARN]\t\t language <%s> defined in {file} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			contentFile = temp
		default:
			log.Printf("[ERROR]\t\t cannot parse {file} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.ContentFile = contentFile
	}

	return newMetadata, checkPass
}

func _buildDocDir(opts *Options, siteMeta *docms.SiteMeta, topicDir, docDir os.DirEntry, fti bleve.Index) error {
	log.Printf("[INFO]\t building document from <%s>...\n", opts.SrcDir+"/"+topicDir.Name()+"/"+docDir.Name())

	docMeta, err := _loadDocumentMetadata(opts, topicDir, docDir)
	if err != nil {
		return err
	}
	newDocMeta, ok := _verifyDocumentMetadata(siteMeta, docMeta)
	if !ok {
		return fmt.Errorf("there is error while checking metadata file")
	}
	log.Printf("[INFO]\t\t metadata verification done.\n")

	contentMap := make(map[string][]byte)
	contentFiles := newDocMeta.GetContentFileMap()
	for lang, f := range contentFiles {
		contentFile := opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/" + f
		if !isFile(contentFile) {
			return fmt.Errorf("content file <%s> not exists", contentFile)
		}
		if buff, err := os.ReadFile(contentFile); err != nil {
			return err
		} else {
			contentMap[lang] = buff
		}
	}

	docId := extractId(topicDir) + ":" + extractId(docDir)
	for lang := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		doc := map[string]string{
			"lang":    lang,
			"content": string(contentMap[lang]),
			"title":   docMeta.GetTitleMap()[lang],
			"summary": docMeta.GetSummaryMap()[lang],
		}
		if err := fti.Index(docId+":"+lang, doc); err != nil {
			return err
		}
	}

	if err := os.Mkdir(opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name(), dirPerm); err != nil {
		return err
	}
	if err := writeFileYaml(opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name()+"/meta.yaml", newDocMeta); err != nil {
		return err
	}

	log.Printf("[INFO]\t\t building content directory <%s>...\n", opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name())
	if err := copyDir(opts.SrcDir+"/"+topicDir.Name()+"/"+docDir.Name(), opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name(), "meta.yaml", "meta.json"); err != nil {
		return err
	}
	return nil
}

// handle command "build"
func actionBuild(c *cli.Context) error {
	opts := Opts(c)
	if fi, err := os.Stat(opts.SrcDir); err != nil || !fi.IsDir() {
		return fmt.Errorf("<%s> is not an existing directory or not readable", opts.SrcDir)
	}

	if fi, err := os.Stat(opts.OutputDir); err == os.ErrNotExist {
		log.Printf("[INFO] directory <%s> does not exist, try creating...\n", opts.OutputDir)
		if err := os.Mkdir(opts.OutputDir, dirPerm); err != nil {
			return err
		}
	} else if err == nil && !fi.IsDir() {
		return fmt.Errorf("<%s> is not a directory", opts.OutputDir)
	} else if err != nil {
		return err
	}

	if outputDirEntries, err := docms.GetDirContent(opts.OutputDir, nil); err != nil {
		return err
	} else if len(outputDirEntries) > 0 {
		if !opts.PurgeOutputDir {
			return fmt.Errorf("output directory <%s> is not empty, empty it then retry or supply flag --%s", opts.OutputDir, flagPurge.Name)
		}
		log.Printf("[INFO] directory <%s> is not empty, clearning up...\n", opts.OutputDir)
		for _, path := range outputDirEntries {
			pathToRemove := opts.OutputDir + "/" + path.Name()
			log.Printf("[INFO] removing <%s>...\n", pathToRemove)
			if err := os.RemoveAll(pathToRemove); err != nil {
				return err
			}
		}
	}

	siteMeta, err := _loadSiteMetadata(opts)
	if err != nil {
		return err
	}
	newSiteMeta, ok := _verifySiteMetadata(siteMeta)
	if !ok {
		return fmt.Errorf("there is error while checking metadata file")
	} else {
		log.Printf("[INFO] metadata verification done.\n")
	}
	if err := writeFileYaml(opts.OutputDir+"/meta.yaml", newSiteMeta); err != nil {
		return err
	}

	ftiDir := opts.OutputDir + "/fti.bleve"
	ftiMapping := bleve.NewIndexMapping()
	fti, err := bleve.New(ftiDir, ftiMapping)
	if err != nil {
		return err
	}
	defer fti.Close()

	srcDirEntries, err := docms.GetDirContent(opts.SrcDir, func(entry os.DirEntry) bool {
		return entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name())
	})
	if err != nil {
		return err
	}

	idmap := map[string]bool{}
	for _, dirEntry := range srcDirEntries {
		matches := docms.RexpContentDir.FindStringSubmatch(dirEntry.Name())
		if _, ok := idmap[matches[2]]; ok {
			return fmt.Errorf("duplicated topic-id at directory <%s>", dirEntry.Name())
		}
		idmap[matches[2]] = true
		if err := _buildTopicDir(opts, siteMeta, dirEntry, fti); err != nil {
			return err
		}
	}

	return nil
}
