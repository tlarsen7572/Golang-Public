package ryxfolder

import (
	h "github.com/tlarsen7572/Golang-Public/helpers"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type RyxFolder struct {
	Path    string
	Folders []*RyxFolder
	Docs    []string
}

var ryxExt = h.StringArray{`.yxmc`, `.yxmd`, `.yxwz`}

func Build(path string) (*RyxFolder, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	entries, err := ioutil.ReadDir(absPath)
	if err != nil {
		return nil, err
	}

	folders := make([]*RyxFolder, 0)
	docs := make([]string, 0)
	for _, entry := range entries {
		newPath := filepath.Join(absPath, entry.Name())
		if entry.IsDir() {
			subfolder, err := Build(newPath)
			if err == nil {
				folders = append(folders, subfolder)
			}
			continue
		}
		if ryxExt.Contains(strings.ToLower(filepath.Ext(newPath))) {
			docs = append(docs, newPath)
		}
	}
	return &RyxFolder{Path: absPath, Folders: folders, Docs: docs}, nil
}

func (ryxFolder *RyxFolder) TotalFiles() int {
	files := len(ryxFolder.Docs)
	for _, folder := range ryxFolder.Folders {
		files += folder.TotalFiles()
	}
	return files
}

func (ryxFolder *RyxFolder) AllFiles() []string {
	var files []string
	for _, folder := range ryxFolder.Folders {
		files = append(files, folder.AllFiles()...)
	}
	files = append(files, ryxFolder.Docs...)
	return files
}
