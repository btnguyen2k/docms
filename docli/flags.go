package main

import (
	"github.com/urfave/cli/v2"
)

const (
	fieldSrc       = "src"
	fieldOutput    = "out"
	fieldPurge     = "purge"
	fieldOverride  = "override"
	fieldDir       = "dir"
	fieldName      = "name"
	fieldIcon      = "icon"
	fieldLanguages = "languages"
	fieldId        = "id"
)

var (
	flagSrc           = &cli.StringFlag{Name: fieldSrc, Value: defaultSrcFolder, Usage: "source directory"}
	flagOutput        = &cli.StringFlag{Name: fieldOutput, Value: defaultOutputFolder, Usage: "output directory"}
	flagPurge         = &cli.BoolFlag{Name: fieldPurge, Usage: "purge output directory if not empty"}
	flagDir           = &cli.StringFlag{Name: fieldDir, Value: defaultDataFolder, Usage: "directory to store data"}
	flagOverride      = &cli.BoolFlag{Name: fieldOverride, Usage: "override if destination exists"}
	flagSiteName      = &cli.StringFlag{Name: fieldName, Usage: "(short) name of the website"}
	flagSiteIcon      = &cli.StringFlag{Name: fieldIcon, Value: defaultSiteIcon, Usage: "icon of the website (support FontAwesome icons)"}
	flagSiteLanguages = &cli.StringFlag{Name: fieldLanguages, Value: defaultSiteLanguages, Usage: "supported languages (format: <code1:label1>[,<code2:label2>].<default:code>)"}
	flagTopicId       = &cli.StringFlag{Name: fieldId, Usage: "topic's unique id"}
	flagTopicIcon     = &cli.StringFlag{Name: fieldIcon, Value: defaultTopicIcon, Usage: "icon of the topic (support FontAwesome icons)"}
)

type Options struct {
	SrcDir         string
	OutputDir      string
	PurgeOutputDir bool
	OverrideTarget bool
	DataDir        string
	SiteName       string
	SiteIcon       string
	SiteLanguages  string
	TopicId        string
	TopicIcon      string
}

func Opts(c *cli.Context) *Options {
	return &Options{
		SrcDir:         c.String(fieldSrc),
		OutputDir:      c.String(fieldOutput),
		PurgeOutputDir: c.Bool(fieldPurge),
		DataDir:        c.String(fieldDir),
		OverrideTarget: c.Bool(fieldOverride),
		SiteName:       c.String(fieldName),
		SiteIcon:       c.String(fieldIcon),
		SiteLanguages:  c.String(fieldLanguages),
		TopicId:        c.String(fieldId),
		TopicIcon:      c.String(fieldIcon),
	}
}
