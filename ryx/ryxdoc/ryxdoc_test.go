package ryxdoc_test

import (
	"encoding/xml"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	r "github.com/tlarsen7572/Golang-Public/ryx/testdocbuilder"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var yxmd = filepath.Join(`..`, `testdocs`, `01 SETLEAF Equations Completed.yxmd`)
var baseFolder = filepath.Join(`..`, `testdocs`)

func TestLoadDocFromFile(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	doc, err := ryxdoc.ReadFile(yxmd)
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	allNodes := doc.ReadMappedNodes()
	if len(allNodes) != 16 {
		t.Fatalf(`expected 16 Nodes but got %v`, len(allNodes))
	}
}

func TestRemoveNodeFromDoc(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	doc, _ := ryxdoc.ReadFile(yxmd)
	doc.RemoveNodes(1, 21)
	allNodes := doc.ReadMappedNodes()
	if len(allNodes) != 14 {
		t.Fatalf(`expected 14 Nodes but got %v`, len(allNodes))
	}
}

func TestAddMacro(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	doc, _ := ryxdoc.ReadFile(yxmd)
	node := doc.AddMacroAt(`.\macro.yxmc`, 5, 10)
	if id, _ := node.ReadId(); id != 24 {
		t.Fatalf(`expected new macro ID to be 24 but got %v`, id)
	}
	if macro := node.ReadMacro(); macro.StoredPath != `.\macro.yxmc` {
		t.Fatalf(`expected macro path of '.\macro.yxmc' but got '%v'`, macro.StoredPath)
	}
	if count := len(doc.ReadMappedNodes()); count != 17 {
		t.Fatalf(`expected 17 Nodes but got %v`, count)
	}
}

func TestRenameMacroNodes(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	docFolder, _ := generateAbsPath(`..`, `testdocs`)
	macroPath, _ := generateAbsPath(`..`, `testdocs`, `macros`, `Tag with Sets.yxmc`)
	newPath, _ := generateAbsPath(`..`, `testdocs`, `macros`, `Tag.yxmc`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	doc.RenameMacroNodes([]string{macroPath}, []string{newPath}, docFolder)
	expectedPath := strings.Replace(newPath, string(os.PathSeparator), `\`, -1)
	if macro := doc.ReadMappedNodes()[18].ReadMacro(); macro.StoredPath != expectedPath {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath, macro.StoredPath)
	}
}

func TestMakeAllMacrosAbsolute(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	docFolder, _ := generateAbsPath(`..`, `testdocs`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	changed := doc.MakeAllMacrosAbsolute(docFolder)
	newPath1, _ := generateAbsPath(`..`, `testdocs`, `macros`, `Tag with Sets.yxmc`)
	newPath2, _ := generateAbsPath(`..`, `testdocs`, `Calculate Filter Expression.yxmc`)
	expectedPath1 := strings.Replace(newPath1, string(os.PathSeparator), `\`, -1)
	expectedPath2 := strings.Replace(newPath2, string(os.PathSeparator), `\`, -1)
	if changed != 2 {
		t.Fatalf(`expected 2 changed Nodes but got %v`, changed)
	}
	if macro := doc.ReadMappedNodes()[18].ReadMacro(); macro.StoredPath != expectedPath1 {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath1, macro.StoredPath)
	}
	if macro := doc.ReadMappedNodes()[12].ReadMacro(); macro.StoredPath != expectedPath2 {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath2, macro.StoredPath)
	}
}

func TestMakeAllMacrosRelative(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	docFolder, _ := generateAbsPath(`..`, `testdocs`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	doc.MakeAllMacrosAbsolute(docFolder)
	changed := doc.MakeAllMacrosRelative(docFolder, docFolder)
	expectedPath1 := `macros\Tag with Sets.yxmc`
	expectedPath2 := `Calculate Filter Expression.yxmc`
	if changed != 2 {
		t.Fatalf(`expected 2 changed Nodes but got %v`, changed)
	}
	if macro := doc.ReadMappedNodes()[18].ReadMacro(); macro.StoredPath != expectedPath1 {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath1, macro.StoredPath)
	}
	if macro := doc.ReadMappedNodes()[12].ReadMacro(); macro.StoredPath != expectedPath2 {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath2, macro.StoredPath)
	}
}

func TestMakeSpecificMacroAbsolute(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	docFolder, _ := generateAbsPath(`..`, `testdocs`)
	macro, _ := generateAbsPath(`..`, `testdocs`, `Calculate Filter Expression.yxmc`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	changed := doc.MakeMacrosAbsolute([]string{macro}, docFolder)
	unchangedPath := `macros\Tag with Sets.yxmc`
	expectedPath := strings.Replace(macro, string(os.PathSeparator), `\`, -1)
	if macro := doc.ReadMappedNodes()[12].ReadMacro(); macro.StoredPath != expectedPath {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath, macro.StoredPath)
	}
	if macro := doc.ReadMappedNodes()[18].ReadMacro(); macro.StoredPath != unchangedPath {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, unchangedPath, macro.StoredPath)
	}
	if changed != 1 {
		t.Fatalf(`expected 1 changed node but got %v`, changed)
	}
}

func TestMakeSpecificMacroRelative(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	docFolder, _ := generateAbsPath(`..`, `testdocs`)
	macro, _ := generateAbsPath(`..`, `testdocs`, `Calculate Filter Expression.yxmc`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	doc.MakeAllMacrosAbsolute(docFolder)
	changed := doc.MakeMacrosRelative([]string{macro}, docFolder, docFolder)
	newPath1, _ := generateAbsPath(`..`, `testdocs`, `macros`, `Tag with Sets.yxmc`)
	expectedPath1 := strings.Replace(newPath1, string(os.PathSeparator), `\`, -1)
	expectedPath2 := `Calculate Filter Expression.yxmc`
	if changed != 1 {
		t.Fatalf(`expected 1 changed node but got %v`, changed)
	}
	if macro := doc.ReadMappedNodes()[18].ReadMacro(); macro.StoredPath != expectedPath1 {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath1, macro.StoredPath)
	}
	if macro := doc.ReadMappedNodes()[12].ReadMacro(); macro.StoredPath != expectedPath2 {
		t.Fatalf(`expected macro path of '%v' but got '%v'`, expectedPath2, macro.StoredPath)
	}
}

func TestExtractAbsoluteMacro(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	newMacroPath, _ := generateAbsPath(`..`, `testdocs`, `new.yxmc`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	err := doc.ExtractMacro(newMacroPath, ``, 13, 14, 15, 16, 17)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if _, err := os.Stat(newMacroPath); err != nil {
		t.Fatalf(err.Error())
	}
	newMacro, err := ryxdoc.ReadFile(newMacroPath)
	if err != nil {
		t.Fatalf(`expected no error reading new macro but got: %v`, err.Error())
	}
	if count := len(newMacro.ReadMappedNodes()); count != 8 {
		t.Fatalf(`expected 8 Nodes in new macro but got %v`, count)
	}
	if count := len(newMacro.Connections); count != 7 {
		t.Fatalf(`expected 7 Connections in new macro but got %v`, count)
	}
	if count := len(doc.ReadMappedNodes()); count != 12 {
		t.Fatalf(`expected 12 Nodes in yxmd but got %v`, count)
	}
	if count := len(doc.Connections); count != 7 {
		t.Fatalf(`expected 7 Connections in yxmd but got %v`, count)
	}
}

func TestExtractMacroWithHole(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	newMacroPath, _ := generateAbsPath(`..`, `testdocs`, `new.yxmc`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	err := doc.ExtractMacro(newMacroPath, ``, 13, 15, 16, 17)
	if err == nil {
		t.Fatalf(`expected an error extracting macro but got none`)
	}
	t.Logf(err.Error())
}

func TestExtractMacroMultipleConnections(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	newMacroPath, _ := generateAbsPath(`..`, `testdocs`, `new.yxmc`)
	doc, _ := ryxdoc.ReadFile(yxmd)
	err := doc.ExtractMacro(newMacroPath, ``, 14, 15, 16)
	if err != nil {
		t.Fatalf(`expected no error reading new macro but got: %v`, err.Error())
	}
	newMacro, err := ryxdoc.ReadFile(newMacroPath)
	if err != nil {
		t.Fatalf(`expected no error reading new macro but got: %v`, err.Error())
	}
	if count := len(newMacro.ReadMappedNodes()); count != 7 {
		t.Fatalf(`expected 7 Nodes in new macro but got %v`, count)
	}
	if count := len(newMacro.Connections); count != 5 {
		t.Fatalf(`expected 5 Connections in new macro but got %v`, count)
	}
	macroConns := newMacro.Connections
	if !listHasConnection(macroConns, 16, `Output`, 20, `Input`) {
		t.Fatalf(`new macro is missing connection to macro output`)
	}
	if !listHasConnection(macroConns, 18, `Output`, 14, `Input`) {
		t.Fatalf(`new macro is missing connection from first macro input`)
	}
	if !listHasConnection(macroConns, 19, `Output`, 15, `Input`) {
		t.Fatalf(`new macro is missing connection from second macro input`)
	}
	docConns := doc.Connections
	if !listHasConnection(docConns, 13, `True`, 24, `Input18`) {
		t.Fatalf(`original doc is missing connection to first macro input`)
	}
	if !listHasConnection(docConns, 13, `False`, 24, `Input19`) {
		t.Fatalf(`original doc is missing connection to second macro input`)
	}
	if !listHasConnection(docConns, 24, `Output20`, 17, `Input`) {
		t.Fatalf(`original doc is missing connection from macro output`)
	}
	err = doc.Save(yxmd)
	if err != nil {
		t.Fatalf(`unexpected error occurred saving the doc: %v`, err.Error())
	}
}

func TestHiddenConnections(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	file, _ := generateAbsPath(`..`, `testdocs`, `01 SETLEAF Equations Completed.yxmd`)
	doc, _ := ryxdoc.ReadFile(file)
	if !doc.Connections[0].Wireless {
		t.Fatalf(`expected the first connection to have a wireless connection but it did not`)
	}
	if doc.Connections[1].Wireless {
		t.Fatalf(`expected the second connection to have a wired connection but it was wireless`)
	}
}

func TestMarshallIndent(t *testing.T) {
	r.RebuildTestdocs(baseFolder)
	defer r.RebuildTestdocs(baseFolder)

	file, _ := generateAbsPath(`..`, `testdocs`, `01 SETLEAF Equations Completed.yxmd`)
	doc, _ := ryxdoc.ReadFile(file)
	marshalled, err := xml.MarshalIndent(doc, ``, `  `)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	t.Logf(string(marshalled))
}

func listHasConnection(conns []*ryxdoc.RyxConn, fromId int, fromAnchor string, toId int, toAnchor string) bool {
	connFound := false
	for _, conn := range conns {
		if conn.FromId == fromId && conn.FromAnchor == fromAnchor && conn.ToId == toId && conn.ToAnchor == toAnchor {
			connFound = true
			break
		}
	}
	return connFound
}

func generateAbsPath(path ...string) (string, error) {
	return filepath.Abs(filepath.Join(path...))
}
