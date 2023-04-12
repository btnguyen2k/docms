package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/btnguyen2k/docms/be-api/src/docms"
	"github.com/urfave/cli/v2"
)

var commandNew = &cli.Command{
	Name:    "new",
	Aliases: []string{"n"},
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

	siteMeta := &docms.SiteMeta{
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
		Tags: map[string]string{
			"build": "${build_datetime}",
		},
		Languages:   map[string]string{},
		Description: map[string]string{},
	}

	// site's name
	if siteMeta.Name == "" {
		return fmt.Errorf("name of website must not be empty, supply website's name with flag --%s", fieldName)
	}

	// site's languages
	if opts.SiteLanguages == "" {
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

	// site's description
	for lang, _ := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		siteMeta.Description.(map[string]string)[lang] = fmt.Sprintf("My awesome website (in %s)", lang)
	}

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
	siteMeta, err := docms.LoadSiteMeta(opts.DataDir+"/meta.yaml", opts.DataDir+"/meta.json")
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
	dirContent, err := docms.GetDirContent(opts.DataDir, func(entry os.DirEntry) bool {
		return entry.IsDir() && docms.RexpContentDir.MatchString(entry.Name())
	})
	index := 1
	for _, dir := range dirContent {
		tokens := docms.RexpContentDir.FindStringSubmatch(dir.Name())
		if topicId == tokens[2] {
			if !opts.OverrideTarget {
				return fmt.Errorf("topic with id <%s> has already existed, remove it then retry or supply flag --%s", topicId, fieldOverride)
			}
			index, _ = strconv.Atoi(tokens[1])
			break
		}
		i, _ := strconv.Atoi(tokens[1])
		if i > index {
			index = i
		}
	}

	topicDir := fmt.Sprintf("%02d-%s", index, topicId)
	metaFile := opts.DataDir + "/" + topicDir + "/meta.yaml"

	if err := os.Mkdir(opts.DataDir+"/"+topicDir, dirPerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	topicMeta := &docms.TopicMeta{
		Title:       map[string]string{},
		Description: map[string]string{},
		Icon:        opts.TopicIcon,
	}

	for lang, _ := range siteMeta.Languages {
		if lang == "default" {
			continue
		}
		topicMeta.Title.(map[string]string)[lang] = fmt.Sprintf("The topic title (in %s)", lang)
		topicMeta.Description.(map[string]string)[lang] = fmt.Sprintf("Short description about the topic (in %s)", lang)
	}

	if err := writeFileYaml(metaFile, topicMeta); err != nil {
		return err
	}

	log.Printf("[INFO] topic metadata has been created at <%s>\n", metaFile)

	return nil
}
