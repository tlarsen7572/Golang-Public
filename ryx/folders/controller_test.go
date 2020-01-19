package folders_test

import (
	"encoding/json"
	server "github.com/tlarsen7572/Golang-Public/ryx/folders"
	"path/filepath"
	"testing"
)

var testFolder, _ = filepath.Abs(filepath.Join(`..`, `testdocs`))

func TestInitializeController(t *testing.T) {
	controller := server.InitializeFolderController(testFolder)
	paths, _ := controller.ReadFolder(``)
	if len(paths) != 1 {
		t.Fatalf(`expected 1 visible path but got %v`, len(paths))
	}
	if paths[0] != testFolder {
		t.Fatalf(`expected visible path of '%v' but got '%v'`, testFolder, paths[0])
	}
}

func TestReadFolder(t *testing.T) {
	controller := server.InitializeFolderController(testFolder)
	paths, _ := controller.ReadFolder(testFolder)
	if count := len(paths); count == 0 {
		t.Fatalf(`expected more than 0 items in the working directory`)
	}
	if paths[0] == testFolder {
		t.Fatalf(`expected a path that was not the working directory`)
	}
	if !filepath.IsAbs(paths[0]) {
		t.Fatalf(`expected absolute paths but got '%v'`, paths[0])
	}
	encoded, _ := json.Marshal(paths)
	t.Logf(string(encoded))
}

func TestReadInvalidFolder(t *testing.T) {
	invalidFolder := filepath.Dir(testFolder)
	controller := server.InitializeFolderController(testFolder)
	_, err := controller.ReadFolder(invalidFolder)
	if err == nil {
		t.Fatalf(`expected an error but none occurred`)
	}
	t.Logf(err.Error())
}

func TestReadRelativeFolderReturnsError(t *testing.T) {
	controller := server.InitializeFolderController(testFolder)
	_, err := controller.ReadFolder(filepath.Join(`..`, `testdocs`, `macros`))
	if err == nil {
		t.Fatalf(`expected an error reading relative path but non occurred`)
	}
	t.Logf(err.Error())
}

func TestReadEmptyFolder(t *testing.T) {
	controller := server.InitializeFolderController(testFolder)
	result, _ := controller.ReadFolder(filepath.Join(testFolder, `macros`))
	if result == nil {
		t.Fatalf(`expected non-nil result but got nil`)
	}
}
