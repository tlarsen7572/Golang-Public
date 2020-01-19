package ryxfolder_test

import (
	"encoding/json"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxfolder"
	"path/filepath"
	"testing"
)

func TestBuildStructure(t *testing.T) {
	path := filepath.Join(`..`, `testdocs`)
	pathAbs, _ := filepath.Abs(path)
	structure, err := ryxfolder.Build(path)
	if err != nil {
		t.Fatalf(`expected no errors but got %v`, err.Error())
	}
	if structure.Path != pathAbs {
		t.Fatalf(`expected path of '%v' but got '%v'`, pathAbs, structure.Path)
	}
	if docs := len(structure.Docs); docs != 5 {
		t.Fatalf(`expected 5 docs but got %v`, docs)
	}
	if structure.Folders[0].Folders == nil {
		t.Fatalf(`expected empty list but got nil`)
	}
	if folders := len(structure.Folders); folders != 1 {
		t.Fatalf(`expected 1 folder but got %v`, folders)
	}
	if docs := len(structure.Folders[0].Docs); docs != 1 {
		t.Fatalf(`expected 1 doc in the subfolder but got %v`, docs)
	}
	if folders := len(structure.Folders[0].Folders); folders != 0 {
		t.Fatalf(`expected 0 folders in the subfolder but got %v`, folders)
	}
	if files := structure.TotalFiles(); files != 6 {
		t.Fatalf(`expected 6 total files but got %v`, files)
	}
	allFiles := structure.AllFiles()
	if len(allFiles) != 6 {
		t.Fatalf(`expected 6 file paths but got %v`, allFiles)
	}
	allFilesStr, _ := json.Marshal(allFiles)
	t.Logf(string(allFilesStr))
}
