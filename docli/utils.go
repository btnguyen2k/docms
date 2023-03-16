package main

import (
	"fmt"
	"os"
)

const (
	defaultSrcFolder    = "dosrc"
	defaultOutputFolder = "dodata"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	switch {
	case err != nil:
		return false
	default:
		return fi.IsDir()
	}
}
