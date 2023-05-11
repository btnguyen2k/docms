package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	// Version defined the version number of DO CMS CLI
	Version = "0.3.1.3"
)

func main() {
	os.Setenv("CLI", "true")
	app := &cli.App{
		Name:    "docli",
		Usage:   "DO CMS website content preprocessing tool",
		Version: Version,
		Authors: []*cli.Author{
			{Name: "Thanh Nguyen", Email: "btnguyen2k (at) gmail (dot) com"},
		},
		Copyright: "Copyright (c) 2023 - DO CMS",
		Commands: []*cli.Command{
			commandBuild,
			commandNew,
			commandTouch,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
