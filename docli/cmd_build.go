package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"main/src/docms"

	"github.com/urfave/cli"
)

var reDirContent = regexp.MustCompile(`^(\d)+-(\w+)$`)

var commandBuild = cli.Command{
	Name:        "build",
	ShortName:   "b",
	Usage:       "Build DO CMS data from source",
	Description: "Build DO CMS data from source directory and write to output directory",
	Flags: []cli.Flag{
		flagSrc,
		flagOutput,
		flagPurge,
	},
	Action: actionBuild,
}

func _loadSiteMetadata(opts *Options) *docms.SiteMeta {
	metaFileYaml := opts.SrcDir + "/meta.yaml"
	fmt.Printf("[INFO] looking for file <%s>...", metaFileYaml)
	if isFile(metaFileYaml) {
		fmt.Printf("found.\n")
		siteMeta, err := docms.LoadSiteMetaFromYaml(metaFileYaml)
		exitIfError(err)
		return siteMeta
	}

	fmt.Printf("not found.\n")
	metaFileJson := opts.SrcDir + "/meta.json"
	fmt.Printf("[INFO] looking for file <%s>...", metaFileJson)
	if isFile(metaFileJson) {
		fmt.Printf("found.\n")
		siteMeta, err := docms.LoadSiteMetaFromJson(metaFileJson)
		exitIfError(err)
		return siteMeta
	}

	fmt.Printf("not found.\n")
	exitIfError(fmt.Errorf("no <%s> or <%s> file found", metaFileYaml, metaFileJson))
	return nil
}

func _verifySiteMetadata(siteMeta *docms.SiteMeta) (*docms.SiteMeta, bool) {
	checkPass := true
	newMetadata := &docms.SiteMeta{}
	fmt.Printf("[INFO] veryfing metadata file...\n")

	// "name" must not empty
	{
		if strings.TrimSpace(siteMeta.Name) == "" {
			fmt.Printf("[ERROR] {name} must not be empty\n")
			checkPass = false
		}
		newMetadata.Name = strings.TrimSpace(siteMeta.Name)
	}

	// "languages" if any, must me a map[string]string
	if len(siteMeta.Languages) == 0 {
		fmt.Printf("[WARN] cannot parse {languages} config or it does not exist, falling back to English...\n")
		newMetadata.Languages = map[string]string{"en": "English"}
	} else {
		newMetadata.Languages = siteMeta.Languages
	}

	// "description" must be a string, or a map[string]string
	{
		desc := siteMeta.Description
		switch desc.(type) {
		case string:
			if strings.TrimSpace(desc.(string)) == "" {
				fmt.Printf("[ERROR] {description} must not be empty\n")
				checkPass = false
			} else {
				desc = strings.TrimSpace(desc.(string))
			}
		case map[string]interface{}:
			if len(desc.(map[string]interface{})) == 0 {
				fmt.Printf("[ERROR] {description} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range desc.(map[string]interface{}) {
				if _, ok := newMetadata.Languages[k]; !ok {
					fmt.Printf("[WARN] language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			desc = temp
		default:
			fmt.Printf("[ERROR] cannot parse {description} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Description = desc
	}

	// icon
	newMetadata.Icon = siteMeta.Icon

	// contacts
	newMetadata.Contacts = siteMeta.Contacts

	return newMetadata, checkPass
}

func _loadTopicMetadata(opts *Options, dir os.DirEntry) *docms.TopicMeta {
	metaFileYaml := opts.SrcDir + "/" + dir.Name() + "/meta.yaml"
	fmt.Printf("[INFO]\t looking for file <%s>...", metaFileYaml)
	if isFile(metaFileYaml) {
		fmt.Printf("found.\n")
		topicMeta, err := docms.LoadTopicMetaFromYaml(metaFileYaml)
		exitIfError(err)
		return topicMeta
	}

	fmt.Printf("not found.\n")
	metaFileJson := opts.SrcDir + "/" + dir.Name() + "/meta.json"
	fmt.Printf("[INFO]\t looking for file <%s>...", metaFileJson)
	if isFile(metaFileJson) {
		fmt.Printf("found.\n")
		topicMeta, err := docms.LoadTopicMetaFromJson(metaFileJson)
		exitIfError(err)
		return topicMeta
	}

	fmt.Printf("not found.\n")
	exitIfError(fmt.Errorf("no <%s> or <%s> file found", metaFileYaml, metaFileJson))
	return nil
}

func _verifyTopicMetadata(siteMeta *docms.SiteMeta, topicMeta *docms.TopicMeta) (*docms.TopicMeta, bool) {
	checkPass := true
	newMetadata := &docms.TopicMeta{}
	fmt.Printf("[INFO]\t veryfing metadata file...\n")

	// "title" must be a string, or a map[string]string
	{
		title := topicMeta.Title
		switch title.(type) {
		case string:
			if strings.TrimSpace(title.(string)) == "" {
				fmt.Printf("[ERROR]\t {title} must not be empty\n")
				checkPass = false
			} else {
				title = strings.TrimSpace(title.(string))
			}
		case map[string]interface{}:
			if len(title.(map[string]interface{})) == 0 {
				fmt.Printf("[ERROR]\t {title} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range title.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					fmt.Printf("[WARN]\t language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			title = temp
		default:
			fmt.Printf("[ERROR]\t cannot parse {title} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Title = title
	}

	// "description" must be a string, or a map[string]string
	{
		desc := topicMeta.Description
		switch desc.(type) {
		case string:
			if strings.TrimSpace(desc.(string)) == "" {
				fmt.Printf("[ERROR]\t {description} must not be empty\n")
				checkPass = false
			} else {
				desc = strings.TrimSpace(desc.(string))
			}
		case map[string]interface{}:
			if len(desc.(map[string]interface{})) == 0 {
				fmt.Printf("[ERROR]\t {description} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range desc.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					fmt.Printf("[WARN]\t language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			desc = temp
		default:
			fmt.Printf("[ERROR]\t cannot parse {description} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Description = desc
	}

	// icon
	newMetadata.Icon = topicMeta.Icon

	return newMetadata, checkPass
}

func _buildTopicDir(opts *Options, siteMeta *docms.SiteMeta, topicDir os.DirEntry) {
	fmt.Printf("[INFO] building topic from <%s>...\n", opts.SrcDir+"/"+topicDir.Name())

	topicMeta := _loadTopicMetadata(opts, topicDir)
	newTopicMeta, ok := _verifyTopicMetadata(siteMeta, topicMeta)
	if !ok {
		exitIfError(fmt.Errorf("there is error while checking metadata file"))
	} else {
		fmt.Printf("[INFO]\t metadata verification done.\n")
	}
	exitIfError(os.Mkdir(opts.OutputDir+"/"+topicDir.Name(), dirPerm))
	exitIfError(writeFileYaml(opts.OutputDir+"/"+topicDir.Name()+"/meta.yaml", newTopicMeta))

	srcDirEntries, err := getDirContent(opts.SrcDir + "/" + topicDir.Name())
	exitIfError(err)
	for _, dirEntry := range srcDirEntries {
		if !dirEntry.IsDir() {
			fmt.Printf("[WARN]\t ignore <%s>: not directory\n", opts.SrcDir+"/"+topicDir.Name()+"/"+dirEntry.Name())
			continue
		}
		matches := reDirContent.FindStringSubmatch(dirEntry.Name())
		if len(matches) == 0 {
			fmt.Printf("[WARN]\t ignore <%s>: invalid name for content directory\n", opts.SrcDir+"/"+dirEntry.Name())
			continue
		}
		_buildDocDir(opts, siteMeta, topicDir, dirEntry)
	}
}

func _loadDocumentMetadata(opts *Options, topicDir, docDir os.DirEntry) *docms.DocumentMeta {
	metaFileYaml := opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/meta.yaml"
	fmt.Printf("[INFO]\t\t looking for file <%s>...", metaFileYaml)
	if isFile(metaFileYaml) {
		fmt.Printf("found.\n")
		docMeta, err := docms.LoadDocumentMetaFromYaml(metaFileYaml)
		exitIfError(err)
		return docMeta
	}

	fmt.Printf("not found.\n")
	metaFileJson := opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/meta.json"
	fmt.Printf("[INFO]\t\t looking for file <%s>...", metaFileJson)
	if isFile(metaFileJson) {
		fmt.Printf("found.\n")
		docMeta, err := docms.LoadDocumentMetaFromJson(metaFileJson)
		exitIfError(err)
		return docMeta
	}

	fmt.Printf("not found.\n")
	exitIfError(fmt.Errorf("no <%s> or <%s> file found", metaFileYaml, metaFileJson))
	return nil
}

func _verifyDocumentMetadata(siteMeta *docms.SiteMeta, docMeta *docms.DocumentMeta) (*docms.DocumentMeta, bool) {
	checkPass := true
	newMetadata := &docms.DocumentMeta{}
	fmt.Printf("[INFO]\t\t veryfing metadata file...\n")

	// "title" must be a string, or a map[string]string
	{
		title := docMeta.Title
		switch title.(type) {
		case string:
			if strings.TrimSpace(title.(string)) == "" {
				fmt.Printf("[ERROR]\t\t {title} must not be empty\n")
				checkPass = false
			} else {
				title = strings.TrimSpace(title.(string))
			}
		case map[string]interface{}:
			if len(title.(map[string]interface{})) == 0 {
				fmt.Printf("[ERROR]\t\t {title} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range title.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					fmt.Printf("[WARN]\t\t language <%s> defined in {description} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			title = temp
		default:
			fmt.Printf("[ERROR]\t\t cannot parse {title} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Title = title
	}

	// "summary" must be a string, or a map[string]string
	{
		summary := docMeta.Summary
		switch summary.(type) {
		case string:
			if strings.TrimSpace(summary.(string)) == "" {
				fmt.Printf("[ERROR]\t\t {summary} must not be empty\n")
				checkPass = false
			} else {
				summary = strings.TrimSpace(summary.(string))
			}
		case map[string]interface{}:
			if len(summary.(map[string]interface{})) == 0 {
				fmt.Printf("[ERROR]\t\t {summary} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range summary.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					fmt.Printf("[WARN]\t\t language <%s> defined in {summary} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			summary = temp
		default:
			fmt.Printf("[ERROR]\t\t cannot parse {summary} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.Summary = summary
	}

	// icon
	newMetadata.Icon = docMeta.Icon

	// "content file" must be a string, or a map[string]string
	{
		contentFile := docMeta.ContentFile
		switch contentFile.(type) {
		case string:
			if strings.TrimSpace(contentFile.(string)) == "" {
				fmt.Printf("[ERROR]\t\t {file} must not be empty\n")
				checkPass = false
			} else {
				contentFile = strings.TrimSpace(contentFile.(string))
			}
		case map[string]interface{}:
			if len(contentFile.(map[string]interface{})) == 0 {
				fmt.Printf("[ERROR]\t\t {file} must not be empty\n")
				checkPass = false
			}
			temp := make(map[string]string)
			for k, v := range contentFile.(map[string]interface{}) {
				if _, ok := siteMeta.Languages[k]; !ok {
					fmt.Printf("[WARN]\t\t language <%s> defined in {file} does not exist\n", k)
				}
				temp[k] = fmt.Sprintf("%s", v)
			}
			contentFile = temp
		default:
			fmt.Printf("[ERROR]\t\t cannot parse {file} config or it does not exist\n")
			checkPass = false
		}
		newMetadata.ContentFile = contentFile
	}

	return newMetadata, checkPass
}

func _buildDocDir(opts *Options, siteMeta *docms.SiteMeta, topicDir, docDir os.DirEntry) {
	fmt.Printf("[INFO]\t building document from <%s>...\n", opts.SrcDir+"/"+topicDir.Name()+"/"+docDir.Name())

	docMeta := _loadDocumentMetadata(opts, topicDir, docDir)
	newDocMeta, ok := _verifyDocumentMetadata(siteMeta, docMeta)
	if !ok {
		exitIfError(fmt.Errorf("there is error while checking metadata file"))
	} else {
		fmt.Printf("[INFO]\t\t metadata verification done.\n")
	}

	contentFiles := newDocMeta.GetContentFileNames()
	for _, f := range contentFiles {
		contentFile := opts.SrcDir + "/" + topicDir.Name() + "/" + docDir.Name() + "/" + f
		if !isFile(contentFile) {
			exitIfError(fmt.Errorf("content file <%s> not exists", contentFile))
		}
	}

	exitIfError(os.Mkdir(opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name(), dirPerm))
	exitIfError(writeFileYaml(opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name()+"/meta.yaml", newDocMeta))

	fmt.Printf("[INFO]\t\t building content directory <%s>...", opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name())
	err := copyDir(opts.SrcDir+"/"+topicDir.Name()+"/"+docDir.Name(), opts.OutputDir+"/"+topicDir.Name()+"/"+docDir.Name(), "meta.yaml", "meta.json")
	if err == nil {
		fmt.Printf("done.\n")
	} else {
		fmt.Printf("\n")
		exitIfError(err)
	}
}

// handle command "build"
func actionBuild(c *cli.Context) {
	opts := Opts(c)
	if !isDir(opts.SrcDir) {
		exitIfError(fmt.Errorf("<%s> is not a directory or not readable", opts.SrcDir))
	}
	if !isDir(opts.OutputDir) {
		fmt.Printf("[INFO] directory <%s> does not exist, try to create...\n", opts.OutputDir)
		err := os.Mkdir(opts.OutputDir, dirPerm)
		exitIfError(err)
	}

	outputDirEntries, err := getDirContent(opts.OutputDir)
	exitIfError(err)
	if len(outputDirEntries) > 0 {
		if !opts.PurgeOutputDir {
			exitIfError(fmt.Errorf("output directory <%s> is not empty, either empty it and retry or supply argument --%s", opts.OutputDir, flagPurge.Name))
		}
		fmt.Printf("[INFO] directory <%s> is not empty, clearning up...\n", opts.OutputDir)
		for _, path := range outputDirEntries {
			pathToRemove := opts.OutputDir + "/" + path.Name()
			fmt.Printf("[INFO] removing <%s>...", pathToRemove)
			err := os.RemoveAll(pathToRemove)
			if err == nil {
				fmt.Printf("done.\n")
			} else {
				fmt.Printf("\n")
			}
			exitIfError(err)
		}
	}

	siteMeta := _loadSiteMetadata(opts)
	newSiteMeta, ok := _verifySiteMetadata(siteMeta)
	if !ok {
		exitIfError(fmt.Errorf("there is error while checking metadata file"))
	} else {
		fmt.Printf("[INFO] metadata verification done.\n")
	}
	exitIfError(writeFileYaml(opts.OutputDir+"/meta.yaml", newSiteMeta))

	srcDirEntries, err := getDirContent(opts.SrcDir)
	exitIfError(err)
	for _, dirEntry := range srcDirEntries {
		if !dirEntry.IsDir() {
			fmt.Printf("[WARN] ignore <%s>: not directory\n", opts.SrcDir+"/"+dirEntry.Name())
			continue
		}
		matches := reDirContent.FindStringSubmatch(dirEntry.Name())
		if len(matches) == 0 {
			fmt.Printf("[WARN] ignore <%s>: invalid name for content directory\n", opts.SrcDir+"/"+dirEntry.Name())
			continue
		}
		_buildTopicDir(opts, siteMeta, dirEntry)
	}
}
