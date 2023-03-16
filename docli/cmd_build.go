package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var commandBuild = cli.Command{
	Name:        "build",
	ShortName:   "b",
	Usage:       "Build DO CMS data from source",
	Description: "Build DO CMS data from source directory and write to output directory",
	Flags: []cli.Flag{
		flagSrc,
		flagOutput,
	},
	Action: actionBuild,
}

// handle command "build"
func actionBuild(c *cli.Context) {
	opts := Opts(c)
	if !isDir(opts.SrcDir) {
		exitIfError(fmt.Errorf("<%s> is not a directory or not readable", opts.SrcDir))
	}
	if !isDir(opts.OutputDir) {
		fmt.Printf("[INFO] directory <%s> does not exist, try to create...\n", opts.OutputDir)
		err := os.Mkdir(opts.OutputDir, 0711)
		exitIfError(err)
	}
}
