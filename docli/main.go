package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	// Version of DO CMS CLI
	Version = "0.1.0"
)

func main() {
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
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
