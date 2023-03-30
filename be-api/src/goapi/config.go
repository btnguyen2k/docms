package goapi

import (
	"log"
	"os"
	"path"
	"path/filepath"

	hoconf "github.com/go-akka/configuration"
	"github.com/go-akka/configuration/hocon"
)

func loadAppConfig(file string) *hoconf.Config {
	// save the current directory and chdir back to it when done
	if curDir, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		defer os.Chdir(curDir)
	}

	log.Printf("[%s] Loading configurations from file [%s]", logLevelInfo, file)
	confDir, confFile := path.Split(file)
	os.Chdir(confDir)

	if data, err := os.ReadFile(confFile); err != nil {
		panic(err)
	} else {
		return hoconf.ParseString(string(data), myIncludeCallback)
	}
}

func myIncludeCallback(filename string) *hocon.HoconRoot {
	if files, err := filepath.Glob(filename); err != nil {
		panic(err)
	} else if len(files) == 0 {
		log.Printf("[%s] [%s] does not match any file", logLevelWarning, filename)
		return hocon.Parse("", nil)
	} else {
		var root = hocon.Parse("", nil)
		for _, f := range files {
			log.Printf("[%s] Loading configurations from file [%s]", logLevelInfo, f)
			if data, err := os.ReadFile(f); err != nil {
				panic(err)
			} else {
				node := hocon.Parse(string(data), myIncludeCallback)
				if node != nil {
					root.Value().GetObject().Merge(node.Value().GetObject())
					// merge substitutions
					subs := make([]*hocon.HoconSubstitution, 0)
					for _, s := range root.Substitutions() {
						subs = append(subs, s)
					}
					for _, s := range node.Substitutions() {
						subs = append(subs, s)
					}
					root = hocon.NewHoconRoot(root.Value(), subs...)
				}
			}
		}
		return root
	}
}
