package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	// Version of DO CMS CLI
	Version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "docli"
	app.Usage = "Pre-process and build content for DO CMS"
	app.Version = Version
	app.Commands = []cli.Command{
		commandBuild,
	}
	app.Run(os.Args)
}
