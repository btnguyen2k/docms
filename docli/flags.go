package main

import (
	"github.com/btnguyen2k/docms/be-api/src/docms"
	"github.com/urfave/cli/v2"
)

const (
	fieldSrc          = "src"
	fieldOutput       = "out"
	fieldPurge        = "purge"
	fieldOverride     = "override"
	fieldDir          = "dir"
	fieldMode         = "mode"
	fieldName         = "name"
	fieldIcon         = "icon"
	fieldLanguages    = "languages"
	fieldId           = "id"
	fieldTopic        = "topic"
	fieldHidden       = "hidden"
	fieldEntryImage   = "img"
	fieldUseTimestamp = "use-timestamp"
	fieldAuthorName   = "author.name"
	fieldAuthorEmail  = "author.email"
	fieldAuthorAvatar = "author.avatar"
)

var (
	flagSrc            = &cli.StringFlag{Name: fieldSrc, Value: defaultSrcFolder, Usage: "source directory"}
	flagOutput         = &cli.StringFlag{Name: fieldOutput, Value: defaultOutputFolder, Usage: "output directory"}
	flagPurge          = &cli.BoolFlag{Name: fieldPurge, Usage: "purge output directory if not empty"}
	flagDir            = &cli.StringFlag{Name: fieldDir, Value: defaultDataFolder, Usage: "directory to store data"}
	flagOverride       = &cli.BoolFlag{Name: fieldOverride, Usage: "override if destination exists"}
	flagSiteName       = &cli.StringFlag{Name: fieldName, Usage: "(short) name of the website"}
	flagSiteIcon       = &cli.StringFlag{Name: fieldIcon, Value: defaultSiteIcon, Usage: "icon of the website (support FontAwesome icons)"}
	flagSiteLanguages  = &cli.StringFlag{Name: fieldLanguages, Value: defaultSiteLanguages, Usage: "supported languages (format: <code1:label1>[,<code2:label2>],<default:code>)"}
	flagSiteMode       = &cli.StringFlag{Name: fieldMode, Value: docms.DefaultSiteMode, Usage: "website's mode (valid values: document, blog)"}
	flagTopicId        = &cli.StringFlag{Name: fieldId, Usage: "topic's unique id"}
	flagTopicIcon      = &cli.StringFlag{Name: fieldIcon, Value: defaultTopicIcon, Usage: "icon of the topic (support FontAwesome icons)"}
	flagTopicEntryImge = &cli.StringFlag{Name: fieldEntryImage, Usage: "topic's entry image url"}
	flagTopicHidden    = &cli.BoolFlag{Name: fieldHidden, Usage: "hide this topic from listing on frontend"}
	flagDocTopic       = &cli.StringFlag{Name: fieldTopic, Usage: "id of document's topic"}
	flagDocId          = &cli.StringFlag{Name: fieldId, Usage: "document's unique id"}
	flagDocIdTimestamp = &cli.BoolFlag{Name: fieldUseTimestamp, Aliases: []string{"ts", "timestamp", "use-ts"}, Usage: "use current timestamp as document's id"}
	flagDocIcon        = &cli.StringFlag{Name: fieldIcon, Value: defaultDocumentIcon, Usage: "icon of the document (support FontAwesome icons)"}
	flagAuthorName     = &cli.StringFlag{Name: fieldAuthorName, Usage: "author's name"}
	flagAuthorEmail    = &cli.StringFlag{Name: fieldAuthorEmail, Usage: "author's email"}
	flagAuthorAvatar   = &cli.StringFlag{Name: fieldAuthorAvatar, Usage: "author's avatar url"}
)

type OptionsCmdBuild struct {
	SrcDir         string
	OutputDir      string
	PurgeOutputDir bool
}

func OptsCmdBuild(c *cli.Context) *OptionsCmdBuild {
	return &OptionsCmdBuild{
		SrcDir:         c.String(fieldSrc),
		OutputDir:      c.String(fieldOutput),
		PurgeOutputDir: c.Bool(fieldPurge),
	}
}

type OptionsCmdNew struct {
	OverrideTarget bool
	DataDir        string
	SiteName,
	SiteIcon,
	SiteLanguages,
	SiteMode string
	AuthorName,
	AuthorEmail,
	AuthorAvatar string
	TopicId,
	TopicIcon,
	TopicEntryImage string
	TopicHidden bool
	DocTopic,
	DocId,
	DocIcon string
	DocIdTimestamp bool
}

func OptsCmdNew(c *cli.Context) *OptionsCmdNew {
	return &OptionsCmdNew{
		DataDir:         c.String(fieldDir),
		OverrideTarget:  c.Bool(fieldOverride),
		SiteName:        c.String(fieldName),
		SiteIcon:        c.String(fieldIcon),
		SiteLanguages:   c.String(fieldLanguages),
		SiteMode:        c.String(fieldMode),
		AuthorName:      c.String(fieldAuthorName),
		AuthorEmail:     c.String(fieldAuthorEmail),
		AuthorAvatar:    c.String(fieldAuthorAvatar),
		TopicId:         c.String(fieldId),
		TopicIcon:       c.String(fieldIcon),
		TopicEntryImage: c.String(fieldEntryImage),
		TopicHidden:     c.Bool(fieldHidden),
		DocTopic:        c.String(fieldTopic),
		DocId:           c.String(fieldId),
		DocIdTimestamp:  c.Bool(fieldUseTimestamp),
		DocIcon:         c.String(fieldIcon),
	}
}
