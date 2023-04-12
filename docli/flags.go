package main

import (
	"github.com/urfave/cli/v2"
)

const (
	fieldSrc    = "src"
	fieldOutput = "out"
	fieldPurge  = "purge"
)

var (
	flagSrc    = &cli.StringFlag{Name: fieldSrc, Value: defaultSrcFolder, Usage: "source directory"}
	flagOutput = &cli.StringFlag{Name: fieldOutput, Value: defaultOutputFolder, Usage: "output directory"}
	flagPurge  = &cli.BoolFlag{Name: fieldPurge, Usage: "purge output directory if not empty"}
)

type Options struct {
	SrcDir         string
	OutputDir      string
	PurgeOutputDir bool
}

func Opts(c *cli.Context) *Options {
	return &Options{
		SrcDir:         c.String(fieldSrc),
		OutputDir:      c.String(fieldOutput),
		PurgeOutputDir: c.Bool(fieldPurge),
	}
}
