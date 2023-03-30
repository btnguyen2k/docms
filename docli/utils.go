package main

import (
	"fmt"
	"io"
	"os"

	"github.com/btnguyen2k/docms/be-api/src/docms"
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

// func getDirContent(path string) ([]os.DirEntry, error) {
// 	dirInfo, err := os.ReadDir(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result := make([]os.DirEntry, 0)
// 	for _, entry := range dirInfo {
// 		if entry.Name() != "." && entry.Name() != ".." {
// 			result = append(result, entry)
// 		}
// 	}
// 	return result, nil
// }

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

func _inList(list []string, target string) bool {
	for _, e := range list {
		if e == target {
			return true
		}
	}
	return false
}

func _copyFile(srcPath, destPath string) error {
	srcFile, err := os.OpenFile(srcPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(destPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}

func copyDir(srcPath, destPath string, ignoreList ...string) error {
	ignoreList = append(ignoreList, ".", "..")
	srcDirEntries, err := docms.GetDirContent(srcPath, func(entry os.DirEntry) bool {
		return !_inList(ignoreList, entry.Name())
	})
	if err != nil {
		return err
	}
	for _, srcEntry := range srcDirEntries {
		// if _inList(ignoreList, srcEntry.Name()) {
		// 	continue
		// }
		if srcEntry.IsDir() {
			if err := os.Mkdir(destPath+"/"+srcEntry.Name(), dirPerm); err != nil {
				return err
			}
			if err := copyDir(srcPath+"/"+srcEntry.Name(), destPath+"/"+srcEntry.Name(), ignoreList...); err != nil {
				return err
			}
		} else {
			if err := _copyFile(srcPath+"/"+srcEntry.Name(), destPath+"/"+srcEntry.Name()); err != nil {
				return err
			}
		}
	}
	return nil
}
