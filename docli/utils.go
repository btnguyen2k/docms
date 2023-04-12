package main

import (
	"io"
	"os"

	"github.com/btnguyen2k/docms/be-api/src/docms"
	"gopkg.in/yaml.v3"
)

const (
	defaultSrcFolder     = "dosrc"
	defaultOutputFolder  = "dodata"
	defaultDataFolder    = "dosrc"
	defaultSiteIcon      = "fas fa-globe"
	defaultSiteLanguages = "en:English,default:en"

	dirPerm  = 0711
	filePerm = 0600
)

func isDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

func isFile(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && !fi.IsDir()
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

func extractId(dir os.DirEntry) string {
	matches := docms.RexpContentDir.FindStringSubmatch(dir.Name())
	return matches[2]
}
