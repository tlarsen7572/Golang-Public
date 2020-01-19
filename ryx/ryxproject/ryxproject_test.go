package ryxproject_test

import (
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxproject"
	r "github.com/tlarsen7572/Golang-Public/ryx/testdocbuilder"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var baseFolder = filepath.Join(`..`, `testdocs`)

func TestOpenProject(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	proj, err := ryxproject.Open(baseFolder)
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	if path := proj.ReadPath(); !filepath.IsAbs(path) {
		t.Fatalf(`expected an absolute path but got '%v'`, path)
	}
	t.Logf(proj.ReadPath())
}

func TestOpenInvalidFolder(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	proj, err := ryxproject.Open("Invalid")
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	if proj != nil {
		t.Fatalf(`expected a nil project`)
	}
}

func TestOpenFileRatherThanFolder(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	proj, err := ryxproject.Open(`ryxproject.go`)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	if proj != nil {
		t.Fatalf(`expected a nil project`)
	}
}

func TestRenameMacro(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	proj, _ := ryxproject.Open(baseFolder)
	oldFile, _ := generateAbsPath(baseFolder, `macros`, `Tag with Sets.yxmc`)
	newFile, _ := generateAbsPath(baseFolder, `macros`, `Tag.yxmc`)
	err := proj.RenameMacro(oldFile, newFile)
	if err != nil {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected no error but got: %v`, err)
	}
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected '%v' to not exist but it does`, oldFile)
	}
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected '%v' to exist but it does not`, newFile)
	}
	workflow, err := ryxdoc.ReadFile(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`))
	if err != nil {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected no error opening workflow but got: %v`, err.Error())
	}
	nodes := workflow.ReadMappedNodes()
	macroPath := nodes[18].ReadMacro()
	expected, _ := generateAbsPath(baseFolder, `macros`, `Tag.yxmc`)
	expected = strings.Replace(expected, string(os.PathSeparator), `\`, -1)
	if macroPath.StoredPath != expected {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expected, macroPath.StoredPath)
	}
	r.RebuildTestdocs(baseFolder)
}

func TestMakeAllMacrosAbsoluteAndRelative(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	proj, _ := ryxproject.Open(baseFolder)
	changed := proj.MakeAllMacrosAbsolute()
	if changed != 2 {
		t.Fatalf(`expected 2 doc changed but got %v`, changed)
	}
	workflowPath, _ := generateAbsPath(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	workflow, _ := ryxdoc.ReadFile(workflowPath)
	nodes := workflow.ReadMappedNodes()
	expected1, _ := generateAbsPath(baseFolder, `Calculate Filter Expression.yxmc`)
	expected1 = strings.Replace(expected1, string(os.PathSeparator), `\`, -1)
	expected2, _ := generateAbsPath(baseFolder, `macros`, `Tag with Sets.yxmc`)
	expected2 = strings.Replace(expected2, string(os.PathSeparator), `\`, -1)
	if actual := nodes[12].ReadMacro().StoredPath; actual != expected1 {
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected1, actual)
	}
	if actual := nodes[18].ReadMacro().StoredPath; actual != expected2 {
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected2, actual)
	}

	proj.MakeAllMacrosRelative()
	if changed != 2 {
		t.Fatalf(`expected 2 doc changed but got %v`, changed)
	}
	workflow, _ = ryxdoc.ReadFile(workflowPath)
	nodes = workflow.ReadMappedNodes()
	expected1 = `Calculate Filter Expression.yxmc`
	expected2 = `macros\Tag with Sets.yxmc`
	if actual := nodes[12].ReadMacro().StoredPath; actual != expected1 {
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected1, actual)
	}
	if actual := nodes[18].ReadMacro().StoredPath; actual != expected2 {
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected2, actual)
	}

	r.RebuildTestdocs(baseFolder)
}

func TestMakeMacroAbsoluteAndRelative(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	proj, _ := ryxproject.Open(baseFolder)
	macro, _ := generateAbsPath(baseFolder, `Calculate Filter Expression.yxmc`)
	changed := proj.MakeMacroAbsolute(macro)
	if changed != 1 {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected 1 doc changed but got %v`, changed)
	}
	workflowPath, _ := generateAbsPath(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	workflow, _ := ryxdoc.ReadFile(workflowPath)
	nodes := workflow.ReadMappedNodes()
	expected1, _ := generateAbsPath(baseFolder, `Calculate Filter Expression.yxmc`)
	expected1 = strings.Replace(expected1, string(os.PathSeparator), `\`, -1)
	expected2 := `macros\Tag with Sets.yxmc`
	if actual := nodes[12].ReadMacro().StoredPath; actual != expected1 {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected1, actual)
	}
	if actual := nodes[18].ReadMacro().StoredPath; actual != expected2 {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected2, actual)
	}

	proj.MakeMacroRelative(macro)
	if changed != 1 {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected 1 doc changed but got %v`, changed)
	}
	workflow, _ = ryxdoc.ReadFile(workflowPath)
	nodes = workflow.ReadMappedNodes()
	expected1 = `Calculate Filter Expression.yxmc`
	expected2 = `macros\Tag with Sets.yxmc`
	if actual := nodes[12].ReadMacro().StoredPath; actual != expected1 {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected1, actual)
	}
	if actual := nodes[18].ReadMacro().StoredPath; actual != expected2 {
		r.RebuildTestdocs(baseFolder)
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected2, actual)
	}

	r.RebuildTestdocs(baseFolder)
}

func TestRetrieveDocument(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	docPath := filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	proj, _ := ryxproject.Open(baseFolder)
	doc, err := proj.RetrieveDocument(docPath)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if doc == nil {
		t.Fatalf(`expected a non-nil document`)
	}
}

func TestWhereUsed(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	docPath, _ := filepath.Abs(filepath.Join(baseFolder, `Calculate Filter Expression.yxmc`))
	proj, _ := ryxproject.Open(baseFolder)
	usages := proj.WhereUsed(docPath)
	if count := len(usages); count != 1 {
		t.Fatalf(`expected 1 usage but got %v`, count)
	}
	usedPath, _ := filepath.Abs(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`))
	if usages[0] != usedPath {
		t.Fatalf(`expected usage in '%v' but got '%v'`, usedPath, usages[0])
	}
}

func generateAbsPath(path ...string) (string, error) {
	return filepath.Abs(filepath.Join(path...))
}
