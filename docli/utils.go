package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	defaultSrcFolder    = "dosrc"
	defaultOutputFolder = "dodata"

	dirPerm  = 0711
	filePerm = 0600
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

func getDirContent(path string) ([]os.DirEntry, error) {
	dirInfo, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]os.DirEntry, 0)
	for _, entry := range dirInfo {
		if entry.Name() != "." && entry.Name() != ".." {
			result = append(result, entry)
		}
	}
	return result, nil
}

func isFile(path string) bool {
	fi, err := os.Stat(path)
	switch {
	case err != nil:
		return false
	default:
		return !fi.IsDir()
	}
}

func writeFileYaml(filePath string, content interface{}) error {
	buff, err := yaml.Marshal(content)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, buff, 0600)
}
