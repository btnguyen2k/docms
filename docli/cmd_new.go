package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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

	return writeFileYaml(metaFile, siteMeta)
}
