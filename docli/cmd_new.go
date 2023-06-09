package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/btnguyen2k/consu/g18"
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
				flagSiteName, flagSiteIcon, flagSiteLanguages, flagSiteMode, flagAuthorName, flagAuthorEmail, flagAuthorAvatar,
			},
			Action: actionNewSite,
		},
		{
			Name:    "topic",
			Aliases: []string{"t"},
			Usage:   "Create new topic metadata",
			Flags: []cli.Flag{
				flagTopicId, flagTopicIcon, flagTopicHidden, flagTopicEntryImage,
			},
			Action: actionNewTopic,
		},
		{
			Name:    "document",
			Aliases: []string{"doc", "d"},
			Usage:   "Create new document metadata",
			Flags: []cli.Flag{
				flagDocTopic, flagDocId, flagDocIcon, flagDocIdTimestamp, flagDocEntryImage, flagDocPage, flagDocStyle, flagAuthorName, flagAuthorEmail, flagAuthorAvatar,
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
	opts := OptsCmdNew(c)

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
	} else {
		siteMeta = &docms.SiteMeta{
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
	if siteMeta.Languages == nil {
		siteMeta.Languages = make(map[string]string)
	}
	if siteMeta.Name == "" || opts.SiteName != "" {
		siteMeta.Name = opts.SiteName
	}
	if siteMeta.Icon == "" || (opts.SiteIcon != "" && opts.SiteIcon != defaultSiteIcon) {
		siteMeta.Icon = opts.SiteIcon
	}
	if siteMeta.Mode == "" || opts.SiteMode != "" {
		siteMeta.Mode = opts.SiteMode
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

	// site's mode
	validSiteModes := []string{docms.SiteModeBlog, docms.SiteModeDoc, docms.SiteModeDocument}
	if g18.FindInSlice(siteMeta.Mode, validSiteModes) < 0 {
		return fmt.Errorf("invalid website mode, valid values are <%s>", validSiteModes)
	}

	// site's author
	if opts.AuthorName != "" || opts.AuthorEmail != "" || opts.AuthorAvatar != "" {
		if siteMeta.Author == nil {
			siteMeta.Author = &docms.Author{}
		}
		if siteMeta.Author.Name == "" {
			siteMeta.Author.Name = opts.AuthorName
			if siteMeta.Author.Name == "" {
				siteMeta.Author.Name = "DOCLI"
			}
		}
		if siteMeta.Author.Email == "" {
			siteMeta.Author.Email = opts.AuthorEmail
			if siteMeta.Author.Email == "" {
				siteMeta.Author.Email = "docli@docms"
			}
		}
		if siteMeta.Author.Avatar == "" {
			siteMeta.Author.Avatar = opts.AuthorAvatar
		}
	}

	if err := writeFileYaml(metaFile, siteMeta); err != nil {
		return err
	}
	log.Printf("[INFO] site metadata has been created at <%s>\n", metaFile)
	return nil
}

// handle command "new topic"
func actionNewTopic(c *cli.Context) error {
	opts := OptsCmdNew(c)

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
	var index *int = nil
	var topicDir *string = nil
	if dirContent, err := docms.GetDirContent(opts.DataDir, func(entry os.DirEntry) bool {
		if entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name()) {
			if tokens := docms.RexpContentDir.FindStringSubmatch(entry.Name()); topicId == tokens[2] {
				val, _ := strconv.Atoi(tokens[1])
				index = &val
				topicDir = g18.PointerOf(entry.Name())
			}
			return true
		}
		return false
	}); err != nil {
		return err
	} else {
		if index != nil {
			if !opts.OverrideTarget {
				return fmt.Errorf("topic with id <%s> has already existed, remove it then retry or supply flag --%s", topicId, fieldOverride)
			}
		} else {
			index = g18.PointerOf(1)
			for _, dir := range dirContent {
				if i, _ := strconv.Atoi(docms.RexpContentDir.FindStringSubmatch(dir.Name())[1]); i >= *index {
					index = g18.PointerOf(i + 1)
				}
			}
		}
	}

	// init topic meta
	if topicDir == nil {
		topicDir = g18.PointerOf(fmt.Sprintf("%02d-%s", *index, topicId))
	}
	metaFile := opts.DataDir + "/" + *topicDir + "/meta.yaml"
	if err := os.Mkdir(opts.DataDir+"/"+*topicDir, dirPerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	topicMeta, err := docms.LoadTopicMetaAuto(opts.DataDir + "/" + *topicDir)
	if err == nil && topicMeta != nil {
		if opts.TopicHidden {
			topicMeta.Hidden = opts.TopicHidden
		}
	} else {
		topicMeta = &docms.TopicMeta{
			Icon:   opts.TopicIcon,
			Hidden: opts.TopicHidden,
		}
	}
	if topicMeta.Icon == "" || (opts.TopicIcon != "" && opts.TopicIcon != defaultTopicIcon) {
		topicMeta.Icon = opts.TopicIcon
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

	// topic's entry-image
	if opts.TopicEntryImage != "" {
		topicMeta.EntryImage = opts.TopicEntryImage
	}

	if err := writeFileYaml(metaFile, topicMeta); err != nil {
		return err
	}
	log.Printf("[INFO] topic metadata has been created at <%s>\n", metaFile)
	return nil
}

// handle command "new document"
func actionNewDocument(c *cli.Context) error {
	opts := OptsCmdNew(c)

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
	var docDir *string = nil
	if dirContent, err := docms.GetDirContent(opts.DataDir+"/"+topicDir, func(entry os.DirEntry) bool {
		if entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name()) {
			if tokens := docms.RexpContentDir.FindStringSubmatch(entry.Name()); docId == tokens[2] {
				val, _ := strconv.Atoi(tokens[1])
				index = &val
				docDir = g18.PointerOf(entry.Name())
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
		} else if opts.DocIdTimestamp || siteMeta.Mode == docms.SiteModeBlog {
			strTime := now.UTC().Format("200601021504")
			val, _ := strconv.Atoi(strTime)
			index = &val
		} else {
			index = g18.PointerOf(1)
			for _, dir := range dirContent {
				if i, _ := strconv.Atoi(docms.RexpContentDir.FindStringSubmatch(dir.Name())[1]); i >= *index {
					index = g18.PointerOf(i + 1)
				}
			}
		}
	}

	// init document meta
	if docDir == nil {
		docDir = g18.PointerOf(fmt.Sprintf("%03d-%s", *index, docId))
	}
	metaFile := opts.DataDir + "/" + topicDir + "/" + *docDir + "/meta.yaml"
	if err := os.Mkdir(opts.DataDir+"/"+topicDir+"/"+*docDir, dirPerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	docMeta, err := docms.LoadDocumentMetaAuto(opts.DataDir + "/" + topicDir + "/" + *docDir)
	if err == nil && docMeta != nil {
		docMeta.TimestampUpdate = now.UTC().Unix()
		if docMeta.TimestampCreate <= 0 {
			docMeta.TimestampCreate = docMeta.FileInfo.ModTime().UTC().Unix()
		}
	} else {
		docMeta = &docms.DocumentMeta{
			Icon:            opts.DocIcon,
			TimestampCreate: now.UTC().Unix(),
			TimestampUpdate: now.UTC().Unix(),
		}
	}

	// document's icon
	if docMeta.Icon == "" || (opts.DocIcon != "" && opts.DocIcon != defaultDocumentIcon) {
		docMeta.Icon = opts.DocIcon
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
			contentFile := opts.DataDir + "/" + topicDir + "/" + *docDir + "/" + docContentFileMap[lang]
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

	// document's entry-image
	if opts.TopicEntryImage != "" {
		docMeta.EntryImage = opts.TopicEntryImage
	}
	// document's page
	if opts.DocPage != "" {
		docMeta.DocPage = opts.DocPage
	}
	// document's style
	if opts.DocStyle != "" {
		docMeta.DocStyle = opts.DocStyle
	}

	// document's author
	if opts.AuthorName != "" || opts.AuthorEmail != "" || opts.AuthorAvatar != "" {
		if docMeta.Author == nil {
			docMeta.Author = &docms.Author{}
		}
		if docMeta.Author.Name == "" {
			docMeta.Author.Name = opts.AuthorName
			if docMeta.Author.Name == "" {
				docMeta.Author.Name = "DOCLI"
			}
		}
		if docMeta.Author.Email == "" {
			docMeta.Author.Email = opts.AuthorEmail
			if docMeta.Author.Email == "" {
				docMeta.Author.Email = "docli@docms"
			}
		}
		if docMeta.Author.Avatar == "" {
			docMeta.Author.Avatar = opts.AuthorAvatar
		}
	}

	if err := writeFileYaml(metaFile, docMeta); err != nil {
		return err
	}
	log.Printf("[INFO] document metadata has been created at <%s>\n", metaFile)
	return nil
}
