package ryxproject_test

import (
	"encoding/json"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxproject"
	r "github.com/tlarsen7572/Golang-Public/ryx/testdocbuilder"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var baseFolder, _ = filepath.Abs(filepath.Join(`..`, `testdocs`))

func TestOpenProject(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

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
	defer r.RebuildTestdocs(baseFolder)

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
	defer r.RebuildTestdocs(baseFolder)

	proj, err := ryxproject.Open(`ryxproject.go`)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	if proj != nil {
		t.Fatalf(`expected a nil project`)
	}
}

func TestRenameFiles(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	oldFile, _ := generateAbsPath(baseFolder, `macros`, `Tag with Sets.yxmc`)
	newFile, _ := generateAbsPath(baseFolder, `macros`, `Tag.yxmc`)
	errFiles, err := proj.RenameFiles([]string{oldFile}, []string{newFile})
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err)
	}
	if count := len(errFiles); count > 0 {
		t.Fatalf(`expected 0 file errors but got %v`, count)
	}
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Fatalf(`expected '%v' to not exist but it does`, oldFile)
	}
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Fatalf(`expected '%v' to exist but it does not`, newFile)
	}
	workflow, err := ryxdoc.ReadFile(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`))
	if err != nil {
		t.Fatalf(`expected no error opening workflow but got: %v`, err.Error())
	}
	nodes := workflow.ReadMappedNodes()
	macroPath := nodes[18].ReadMacro()
	expected, _ := generateAbsPath(baseFolder, `macros`, `Tag.yxmc`)
	expected = strings.Replace(expected, string(os.PathSeparator), `\`, -1)
	if macroPath.StoredPath != expected {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expected, macroPath.StoredPath)
	}
}

func TestRenameFilesWithUnequalLists(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	_, err := proj.RenameFiles([]string{`A`}, []string{`B`, `C`})
	if err == nil {
		t.Fatalf(`expected an error but none occurred`)
	}
}

func TestMakeAllMacrosAbsoluteAndRelative(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	changed := proj.MakeAllFilesAbsolute()
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

	proj.MakeAllFilesRelative()
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
}

func TestMakeMacroAbsoluteAndRelative(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	macro, _ := generateAbsPath(baseFolder, `Calculate Filter Expression.yxmc`)
	changed := proj.MakeFilesAbsolute([]string{macro})
	if changed != 1 {
		t.Fatalf(`expected 1 doc changed but got %v`, changed)
	}
	workflowPath, _ := generateAbsPath(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	workflow, _ := ryxdoc.ReadFile(workflowPath)
	nodes := workflow.ReadMappedNodes()
	expected1, _ := generateAbsPath(baseFolder, `Calculate Filter Expression.yxmc`)
	expected1 = strings.Replace(expected1, string(os.PathSeparator), `\`, -1)
	expected2 := `macros\Tag with Sets.yxmc`
	if actual := nodes[12].ReadMacro().StoredPath; actual != expected1 {
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected1, actual)
	}
	if actual := nodes[18].ReadMacro().StoredPath; actual != expected2 {
		t.Fatalf(`expected stored path of '%v' but got '%v'`, expected2, actual)
	}

	proj.MakeFilesRelative([]string{macro})
	if changed != 1 {
		t.Fatalf(`expected 1 doc changed but got %v`, changed)
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
}

func TestMakeWorkflowRelativeAndAbsolute(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	workflow, _ := generateAbsPath(baseFolder, `01 SETLEAF Equations Completed.yxmd`)

	changed := proj.MakeFilesAbsolute([]string{workflow})
	if changed != 1 {
		t.Fatalf(`expected 1 changed file but got %v`, changed)
	}
	doc, _ := ryxdoc.ReadFile(workflow)
	macro := doc.ReadMappedNodes()[18].ReadMacro()
	expectedStoredPath := strings.Replace(filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`), string(os.PathSeparator), `\`, -1)
	if macro.StoredPath != expectedStoredPath {
		t.Fatalf(`expected stored path '%v' but got '%v'`, expectedStoredPath, macro.StoredPath)
	}

	changed = proj.MakeFilesRelative([]string{workflow})
	if changed != 1 {
		t.Fatalf(`expected 1 changed file but got %v`, changed)
	}
	doc, _ = ryxdoc.ReadFile(workflow)
	macro = doc.ReadMappedNodes()[18].ReadMacro()
	expectedStoredPath = strings.Replace(filepath.Join(`macros`, `Tag with Sets.yxmc`), string(os.PathSeparator), `\`, -1)
	if macro.StoredPath != expectedStoredPath {
		t.Fatalf(`expected stored path '%v' but got '%v'`, expectedStoredPath, macro.StoredPath)
	}

}

func TestRetrieveDocument(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

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
	defer r.RebuildTestdocs(baseFolder)

	docPath := filepath.Join(baseFolder, `Calculate Filter Expression.yxmc`)
	proj, _ := ryxproject.Open(baseFolder)
	usages := proj.WhereUsed(docPath)
	if count := len(usages); count != 1 {
		t.Fatalf(`expected 1 usage but got %v`, count)
	}
	usedPath := filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	if usages[0] != usedPath {
		t.Fatalf(`expected usage in '%v' but got '%v'`, usedPath, usages[0])
	}
}

func TestMoveFiles(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	files := []string{
		filepath.Join(baseFolder, `Calculate Filter Expression.yxmc`),
		filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`),
	}
	moveTo := filepath.Join(baseFolder, `macros`)
	failedMoves, err := proj.MoveFiles(files, moveTo)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if count := len(failedMoves); count != 0 {
		t.Fatalf(`expected 0 move errors but got %v`, count)
	}
	newFiles := []string{
		filepath.Join(baseFolder, `macros`, `Calculate Filter Expression.yxmc`),
		filepath.Join(baseFolder, `macros`, `01 SETLEAF Equations Completed.yxmd`),
	}
	if _, err := os.Stat(newFiles[0]); os.IsNotExist(err) {
		t.Fatalf(`file '%v' did not exist after the rename`, newFiles[0])
	}
	if _, err := os.Stat(newFiles[1]); os.IsNotExist(err) {
		t.Fatalf(`file '%v' did not exist after the rename`, newFiles[1])
	}
	if _, err := os.Stat(files[0]); !os.IsNotExist(err) {
		t.Fatalf(`file '%v' still exists after the rename`, files[0])
	}
	if _, err := os.Stat(files[1]); !os.IsNotExist(err) {
		t.Fatalf(`file '%v' still exist after the rename`, files[1])
	}
	doc, _ := ryxdoc.ReadFile(newFiles[1])
	macroNode := doc.ReadMappedNodes()[12]
	macro := macroNode.ReadMacro(filepath.Dir(newFiles[1]))
	if macro.FoundPath != newFiles[0] {
		t.Fatalf(`expected tool 12 to be a macro at '%v', but got '%v': stored path was '%v'`, newFiles[0], macro.FoundPath, macro.StoredPath)
	}
}

func TestMoveWorkflowWithRelativeMacros(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	proj.MakeAllFilesRelative()
	files := []string{
		filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`),
	}
	moveTo := filepath.Join(baseFolder, `macros`)
	failedMoves, err := proj.MoveFiles(files, moveTo)
	if err != nil || len(failedMoves) > 0 {
		t.Fatalf(`errors occurred moving the workflow`)
	}

	newLocation := filepath.Join(baseFolder, `macros`, `01 SETLEAF Equations Completed.yxmd`)
	expectedMacro := filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`)
	doc, err := ryxdoc.ReadFile(newLocation)
	if err != nil {
		t.Fatalf(`error loading moved file`)
	}
	node := doc.ReadMappedNodes()[18]
	macroPath := node.ReadMacro()
	if macroPath.FoundPath != expectedMacro {
		t.Fatalf("could not find expected macro.\nexpected: %v\nfound: %v\nstored: %v", expectedMacro, macroPath.FoundPath, macroPath.StoredPath)
	}
}

func TestRenameFolder(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	from := filepath.Join(baseFolder, `macros`)
	to := `stuff`
	proj, _ := ryxproject.Open(baseFolder)
	err := proj.RenameFolder(from, to)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	toPath := filepath.Join(baseFolder, to)
	if _, err = os.Stat(from); os.IsExist(err) {
		t.Fatalf(`expected the from folder to no longer exist, but it does`)
	}
	if _, err = os.Stat(toPath); os.IsNotExist(err) {
		t.Fatalf(`expected the to folder to exist, but it does not`)
	}

	workflow := filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	expectedMacro := filepath.Join(baseFolder, `stuff`, `Tag with Sets.yxmc`)
	doc, err := ryxdoc.ReadFile(workflow)
	if err != nil {
		t.Fatalf(`error loading workflow`)
	}
	node := doc.ReadMappedNodes()[18]
	macroPath := node.ReadMacro()
	if macroPath.FoundPath != expectedMacro {
		t.Fatalf("could not find expected macro.\nexpected: %v\nfound: %v\nstored: %v", expectedMacro, macroPath.FoundPath, macroPath.StoredPath)
	}
}

func TestListMacrosUsedInProject(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	macros, err := proj.ListMacrosUsedInProject()
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	tagWithSets, ok := macros[`Tag with Sets.yxmc`]

	if !ok {
		t.Fatalf(`could not find 'Tag with Sets.yxmc' macro but it is used in the project`)
	}

	found := filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`)
	foundPath, ok := tagWithSets.FoundPaths[found]
	if !ok {
		t.Fatalf(`found path '%v' is not in list but it exists`, found)
	}

	stored := filepath.Join(`macros`, `Tag with Sets.yxmc`)
	storedPath, ok := foundPath.StoredPaths[stored]
	if !ok {
		t.Fatalf(`stored path '%v' is not in list but it exists`, stored)
	}

	if count := len(storedPath.WhereUsed); count != 1 {
		t.Fatalf(`expected 1 where used but got %v`, count)
	}
	where := filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`)
	if storedPath.WhereUsed[0] != where {
		t.Fatalf(`expected where used of '%v' but got '%v'`, where, storedPath.WhereUsed[0])
	}
	data, _ := json.Marshal(macros)
	t.Logf(string(data))
}

func TestBatchChangeMacroSetting(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	newSetting := filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`)
	changed, err := proj.BatchChangeMacroSettings(`Tag with Sets.yxmc`, newSetting, nil, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if changed != 1 {
		t.Fatalf(`expected 1 changed macro but got %v`, changed)
	}
	doc, err := proj.RetrieveDocument(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`))
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	var nodes = doc.ReadMappedNodes()
	if setting := nodes[18].ReadMacro().StoredPath; setting != newSetting {
		t.Fatalf(`expected stored macro path to be '%v' but got '%v'`, newSetting, setting)
	}
}

func TestBatchChangeMacroSettingOnlyMatchingFoundPaths(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	relativeSetting := filepath.Join(`macros`, `Tag with Sets.yxmc`)
	absSetting := filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`)
	changed, err := proj.BatchChangeMacroSettings(`Tag with Sets.yxmc`, absSetting, []string{`not\found\here`}, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if changed != 0 {
		t.Fatalf(`expected 0 changed macros but got %v`, changed)
	}
	doc, err := proj.RetrieveDocument(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`))
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	var nodes = doc.ReadMappedNodes()
	if setting := nodes[18].ReadMacro().StoredPath; setting != relativeSetting {
		t.Fatalf(`expected stored macro path to be '%v' but got '%v'`, relativeSetting, setting)
	}
}

func TestBatchChangeMacroSettingOnlyMatchingStoredPaths(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	proj, _ := ryxproject.Open(baseFolder)
	relativeSetting := filepath.Join(`macros`, `Tag with Sets.yxmc`)
	absSetting := filepath.Join(baseFolder, `macros`, `Tag with Sets.yxmc`)
	changed, err := proj.BatchChangeMacroSettings(`Tag with Sets.yxmc`, absSetting, nil, []string{`not\found\here`})
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if changed != 0 {
		t.Fatalf(`expected 0 changed macros but got %v`, changed)
	}
	doc, err := proj.RetrieveDocument(filepath.Join(baseFolder, `01 SETLEAF Equations Completed.yxmd`))
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	var nodes = doc.ReadMappedNodes()
	if setting := nodes[18].ReadMacro().StoredPath; setting != relativeSetting {
		t.Fatalf(`expected stored macro path to be '%v' but got '%v'`, relativeSetting, setting)
	}
}

func generateAbsPath(path ...string) (string, error) {
	return filepath.Abs(filepath.Join(path...))
}
