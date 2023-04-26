package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/btnguyen2k/docms/be-api/src/docms"
	"github.com/urfave/cli/v2"
)

var commandNew = &cli.Command{
	Name:    "new",
	Aliases: []string{"n", "create", "c"},
	Usage:   "Helper to create assets with default metadata",
	Flags: []cli.Flag{
		flagDir,
		flagOverride,
	},
	Subcommands: []*cli.Command{
		{
			Name:    "site",
			Aliases: []string{"s"},
			Usage:   "Create new site content metadata",
			Flags: []cli.Flag{
				flagSiteName, flagSiteIcon, flagSiteLanguages,
			},
			Action: actionNewSite,
		},
		{
			Name:    "topic",
			Aliases: []string{"t"},
			Usage:   "Create new topic metadata",
			Flags: []cli.Flag{
				flagTopicId, flagTopicIcon,
			},
			Action: actionNewTopic,
		},
		{
			Name:    "document",
			Aliases: []string{"doc", "d"},
			Usage:   "Create new document metadata",
			Flags: []cli.Flag{
				flagDocTopic, flagDocId, flagDocIcon, flagDocIdTimestamp,
			},
			Action: actionNewDocument,
		},
	},
}

func _validateDataDirMustBeValid(dir string) error {
	fi, err := os.Stat(dir)
	if dir == "" || (err == nil && !fi.IsDir()) {
		return fmt.Errorf("invalid data directory <%s>", dir)
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func _validateDataDirMustExist(dir string) error {
	if fi, err := os.Stat(dir); err == nil && fi.IsDir() {
		return nil
	}
	return fmt.Errorf("<%s> is not a directory or does not exist", dir)
}

// handle command "new site"
func actionNewSite(c *cli.Context) error {
	opts := Opts(c)

	// data dir
	if err := _validateDataDirMustBeValid(opts.DataDir); err != nil {
		return err
	}
	metaFile := opts.DataDir + "/meta.yaml"
	if isDir(opts.DataDir) && isFile(metaFile) && !opts.OverrideTarget {
		return fmt.Errorf("file <%s> exists, remove it then retry or supply flag --%s", metaFile, fieldOverride)
	}

	if !isDir(opts.DataDir) {
		if err := os.Mkdir(opts.DataDir, dirPerm); err != nil {
			return err
		}
	}

	// init site meta
	siteMeta, err := docms.LoadSiteMetaAuto(opts.DataDir)
	if err == nil && siteMeta != nil {
		if siteMeta.Name == "" || opts.SiteName != "" {
			siteMeta.Name = opts.SiteName
		}
		if siteMeta.Icon == "" || (opts.SiteIcon != "" && opts.SiteIcon != defaultSiteIcon) {
			siteMeta.Icon = opts.SiteIcon
		}
		if siteMeta.Languages == nil {
			siteMeta.Languages = make(map[string]string)
		}
	} else {
		siteMeta = &docms.SiteMeta{
			Name: opts.SiteName,
			Icon: opts.SiteIcon,
			Contacts: map[string]string{
				"website":  "https://my/awesome/website(optional)",
				"email":    "my-email-address-(optional)",
				"github":   "https://github.com/my-github/(optional)",
				"facebook": "https://www.facebook.com/my-fb/(optional)",
				"linkedin": "https://www.linkedin.com/in/my-linkedin/(optional)",
				"slack":    "https://join/my/slack/channel/(optional)",
				"twitter":  "https://follow/me/on/twitter/(optional)",
				"discord":  "https://join/my/discord/(optional)",
			},
			Tags: map[string]interface{}{
				"build": "${build_datetime}",
			},
		}
	}

	// site's name
	if siteMeta.Name == "" {
		return fmt.Errorf("name of website must not be empty, supply website's name with flag --%s", fieldName)
	}

	// site's languages
	if opts.SiteLanguages == "" && len(siteMeta.Languages) == 0 {
		opts.SiteLanguages = defaultSiteLanguages
	}
	codesLabels := strings.Split(opts.SiteLanguages, ",")
	for _, codeLabel := range codesLabels {
		tokens := strings.SplitN(strings.TrimSpace(codeLabel), ":", 2)
		if len(tokens) != 2 {
			log.Printf("[WARN] invalid language code:label pair <%s>", codeLabel)
		} else {
			siteMeta.Languages[strings.TrimSpace(tokens[0])] = strings.TrimSpace(tokens[1])
		}
	}
	// default language
	if len(siteMeta.Languages) == 0 {
		siteMeta.Languages = map[string]string{"default": "en"}
	} else if siteMeta.Languages["default"] == "" {
		for k := range siteMeta.Languages {
			siteMeta.Languages["default"] = k
			break
		}
	}

	// site's description
	siteDescMap := siteMeta.GetDescriptionMap()
	for lang, _ := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		if desc, ok := siteDescMap[lang]; desc == "" || !ok {
			siteDescMap[lang] = fmt.Sprintf("My awesome website (in %s)", lang)
		}
	}
	siteMeta.Description = siteDescMap

	// site's tag alias
	siteTagAliasMap := siteMeta.GetTagAliasMap()
	for lang, _ := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		if tagAlias := siteTagAliasMap[lang]; tagAlias == nil {
			siteTagAliasMap[lang] = map[string][]string{
				"mytag": {"variances", "of", "mytag"},
			}
		}
	}
	siteMeta.TagsAlias = siteTagAliasMap

	if err := writeFileYaml(metaFile, siteMeta); err != nil {
		return err
	}
	log.Printf("[INFO] site metadata has been created at <%s>\n", metaFile)
	return nil
}

// handle command "new topic"
func actionNewTopic(c *cli.Context) error {
	opts := Opts(c)

	// data dir
	if err := _validateDataDirMustExist(opts.DataDir); err != nil {
		return err
	}
	siteMeta, err := docms.LoadSiteMetaAuto(opts.DataDir)
	if err != nil {
		return err
	}

	// topic's id
	topicId := opts.TopicId
	if topicId == "" {
		return fmt.Errorf("id of topic must not be empty, supply topic's id with flag --%s", fieldId)
	}
	if !reId.MatchString(topicId) {
		return fmt.Errorf("invalid topic's id <%s>, should be lower cases and contain only letters and digits", topicId)
	}

	// topic's index
	index := 1
	if dirContent, err := docms.GetDirContent(opts.DataDir, func(entry os.DirEntry) bool {
		return entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name())
	}); err != nil {
		return err
	} else {
		for _, dir := range dirContent {
			tokens := docms.RexpContentDir.FindStringSubmatch(dir.Name())
			if topicId == tokens[2] {
				if !opts.OverrideTarget {
					return fmt.Errorf("topic with id <%s> has already existed, remove it then retry or supply flag --%s", topicId, fieldOverride)
				}
				index, _ = strconv.Atoi(tokens[1])
				break
			}
			if i, _ := strconv.Atoi(tokens[1]); i >= index {
				index = i + 1
			}
		}
	}

	// init topic meta
	topicDir := fmt.Sprintf("%02d-%s", index, topicId)
	metaFile := opts.DataDir + "/" + topicDir + "/meta.yaml"
	if err := os.Mkdir(opts.DataDir+"/"+topicDir, dirPerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	topicMeta, err := docms.LoadTopicMetaAuto(opts.DataDir + "/" + topicDir)
	if err == nil && topicMeta != nil {
		if topicMeta.Icon == "" || (opts.TopicIcon != "" && opts.TopicIcon != defaultTopicIcon) {
			topicMeta.Icon = opts.TopicIcon
		}
	} else {
		topicMeta = &docms.TopicMeta{
			Icon: opts.TopicIcon,
		}
	}

	// topic's title and description
	topicTitleMap := topicMeta.GetTitleMap()
	topicDescMap := topicMeta.GetDescriptionMap()
	for lang, _ := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		if title, ok := topicTitleMap[lang]; title == "" || !ok {
			topicTitleMap[lang] = fmt.Sprintf("The topic title (in %s)", lang)
		}
		if desc, ok := topicDescMap[lang]; desc == "" || !ok {
			topicDescMap[lang] = fmt.Sprintf("Short description about the topic (in %s)", lang)
		}
	}
	topicMeta.Title = topicTitleMap
	topicMeta.Description = topicDescMap

	if err := writeFileYaml(metaFile, topicMeta); err != nil {
		return err
	}
	log.Printf("[INFO] topic metadata has been created at <%s>\n", metaFile)
	return nil
}

func _pint(val int) *int {
	return &val
}

// handle command "new document"
func actionNewDocument(c *cli.Context) error {
	opts := Opts(c)

	// data dir
	if err := _validateDataDirMustExist(opts.DataDir); err != nil {
		return err
	}
	siteMeta, err := docms.LoadSiteMetaAuto(opts.DataDir)
	if err != nil {
		return err
	}

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
	var index *int = nil
	if dirContent, err := docms.GetDirContent(opts.DataDir+"/"+topicDir, func(entry os.DirEntry) bool {
		if entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name()) {
			if tokens := docms.RexpContentDir.FindStringSubmatch(entry.Name()); docId == tokens[2] {
				val, _ := strconv.Atoi(tokens[1])
				index = &val
			}
			return true
		}
		return false
	}); err != nil {
		return err
	} else {
		if index != nil {
			if !opts.OverrideTarget {
				return fmt.Errorf("document with id <%s> has already existed, remove it then retry or supply flag --%s", docId, fieldOverride)
			}
		} else if opts.DocIdTimestamp {
			strTime := time.Now().UTC().Format("200601021504")
			val, _ := strconv.Atoi(strTime)
			index = &val
		} else {
			index = _pint(1)
			for _, dir := range dirContent {
				if i, _ := strconv.Atoi(docms.RexpContentDir.FindStringSubmatch(dir.Name())[1]); i >= *index {
					index = _pint(i + 1)
				}
			}
		}
	}

	// init document meta
	docDir := fmt.Sprintf("%03d-%s", *index, docId)
	metaFile := opts.DataDir + "/" + topicDir + "/" + docDir + "/meta.yaml"
	if err := os.Mkdir(opts.DataDir+"/"+topicDir+"/"+docDir, dirPerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	docMeta, err := docms.LoadDocumentMetaAuto(opts.DataDir + "/" + topicDir + "/" + docDir)
	if err == nil && docMeta != nil {
		if docMeta.Icon == "" || (opts.DocIcon != "" && opts.DocIcon != defaultDocumentIcon) {
			docMeta.Icon = opts.DocIcon
		}
	} else {
		docMeta = &docms.DocumentMeta{
			Icon: opts.DocIcon,
		}
	}

	// document's title, summary, tags and content files
	docTitleMap := docMeta.GetTitleMap()
	docSummaryMap := docMeta.GetSummaryMap()
	docContentFileMap := docMeta.GetContentFileMap()
	docTags := docMeta.GetTagsMap()
	for lang, _ := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		if title, ok := docTitleMap[lang]; title == "" || !ok {
			docTitleMap[lang] = fmt.Sprintf("The document title (in %s)", lang)
		}
		if summary, ok := docSummaryMap[lang]; summary == "" || !ok {
			docSummaryMap[lang] = fmt.Sprintf("Summary about document content (in %s)", lang)
		}
		if tags, ok := docTags[lang]; len(tags) == 0 || !ok {
			re := regexp.MustCompile(`[\s\W]+`)
			docTags[lang] = strings.Split(strings.TrimSpace(re.ReplaceAllString(strings.ToLower(docTitleMap[lang]), " ")), " ")
		}
		if file, ok := docContentFileMap[lang]; file == "" || !ok {
			docContentFileMap[lang] = fmt.Sprintf("index-%s.md", lang)
			contentFile := opts.DataDir + "/" + topicDir + "/" + docDir + "/" + docContentFileMap[lang]
			if !isFile(contentFile) {
				mdContent := fmt.Sprintf("# Document content (in %s)\n\nThe awesome content goes here.", lang)
				os.WriteFile(contentFile, []byte(mdContent), filePerm)
			}
		}
	}
	docMeta.Title = docTitleMap
	docMeta.Summary = docSummaryMap
	docMeta.ContentFile = docContentFileMap
	docMeta.Tags = docTags

	if err := writeFileYaml(metaFile, docMeta); err != nil {
		return err
	}
	log.Printf("[INFO] document metadata has been created at <%s>\n", metaFile)
	return nil
}
