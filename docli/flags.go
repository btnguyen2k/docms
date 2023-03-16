package main

import (
	"github.com/urfave/cli"
)

const (
	fieldSrc    = "src"
	fieldOutput = "out"
)

var (
	flagSrc    = cli.StringFlag{Name: fieldSrc, Value: defaultSrcFolder, Usage: "source directory"}
	flagOutput = cli.StringFlag{Name: fieldOutput, Value: defaultOutputFolder, Usage: "output directory"}
)

type Options struct {
	SrcDir    string
	OutputDir string
}

func Opts(c *cli.Context) *Options {
	return &Options{
		SrcDir:    c.String(fieldSrc),
		OutputDir: c.String(fieldOutput),
	}
}
