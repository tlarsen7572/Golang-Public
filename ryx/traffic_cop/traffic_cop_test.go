package traffic_cop_test

import (
	"encoding/json"
	"github.com/tlarsen7572/Golang-Public/ryx/config"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxfolder"
	"github.com/tlarsen7572/Golang-Public/ryx/testdocbuilder"
	"github.com/tlarsen7572/Golang-Public/ryx/tool_data_loader"
	cop "github.com/tlarsen7572/Golang-Public/ryx/traffic_cop"
	"os"
	"path/filepath"
	"testing"
)

var workFolder, _ = filepath.Abs(filepath.Join(`..`, `testdocs`))

type params map[string]interface{}

func TestChannelInvalidProject(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{Project: `blah blah blah`, Function: `GetStructure`, Out: out, Config: &config.Config{}}
	response := <-out
	if response.Err == nil {
		t.Fatalf(`expected an error but none was returned`)
	}
	t.Logf(response.Err.Error())
}

func TestInvalidFunction(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{Project: workFolder, Function: `Invalid`, Out: out, Config: &config.Config{}}
	response := <-out
	if response.Err == nil {
		t.Fatalf(`expected an error but none was returned`)
	}
	t.Logf(response.Err.Error())
}

func TestGetProjectStructure(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{Project: workFolder, Function: `GetProjectStructure`, Out: out, Config: &config.Config{}}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	if response.Response == nil {
		t.Fatalf(`expected a non-nil response but got nil`)
	}
	switch v := response.Response.(type) {
	case *ryxfolder.RyxFolder:
		break
	default:
		t.Fatalf(`unexpected type %T`, v)
	}
}

func TestGetDocumentStructure(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `01 SETLEAF Equations Completed.yxmd`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Project:    workFolder,
		Function:   `GetDocumentStructure`,
		Parameters: params{`FilePath`: doc},
		Out:        out,
		Config: &config.Config{
			ToolData: []tool_data_loader.ToolData{},
		},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	if response.Response == nil {
		t.Fatalf(`expected a non-nil response but got nil`)
	}
	switch v := response.Response.(type) {
	case cop.DocumentStructure:
		break
	default:
		t.Fatalf(`unexpected type %T`, v)
	}
	encoded, _ := json.Marshal(response)
	t.Logf(string(encoded))
}

func TestGetDocStructureHasMacroToolData(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `01 SETLEAF Equations Completed.yxmd`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Project:    workFolder,
		Function:   `GetDocumentStructure`,
		Parameters: params{`FilePath`: doc},
		Out:        out,
		Config: &config.Config{
			ToolData: []tool_data_loader.ToolData{},
		},
	}
	response := <-out
	structure := response.Response.(cop.DocumentStructure)
	if count := len(structure.MacroToolData); count != 2 {
		t.Fatalf(`expected 2 macro tool data entries but got %v`, count)
	}
	if structure.MacroToolData[0].Plugin == `` {
		t.Fatalf(`expected a non-empty plugin but it was empty`)
	}
}

func TestGetDocumentStructureExcludesInvalidNodes(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `Calculate Filter Expression.yxmc`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{Project: workFolder, Function: `GetDocumentStructure`, Parameters: map[string]interface{}{`FilePath`: doc}, Out: out, Config: &config.Config{}}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err)
	}
	structure := response.Response.(cop.DocumentStructure)
	if count := len(structure.Nodes); count != 5 {
		t.Fatalf(`expected 5 nodes but got %v`, count)
	}
}

func TestGetRootFolders(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    ``,
		Function:   `BrowseFolder`,
		Parameters: params{`FolderPath`: ``},
		Config:     &config.Config{BrowseFolderRoots: []string{`C:`, `D:`}},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err)
	}
	folders, ok := response.Response.([]string)
	if !ok {
		t.Fatalf(`expected a list of strings but got something else`)
	}
	if count := len(folders); count != 2 {
		t.Fatalf(`expected 2 subfolders but got %v`, count)
	}
	if folders[0] != `C:` {
		t.Fatalf(`expected first folder to be 'C:' but got '%v'`, folders[0])
	}

}

func TestBrowseFoldersWithoutFolderPath(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    ``,
		Function:   `BrowseFolder`,
		Parameters: params{},
		Config:     &config.Config{BrowseFolderRoots: []string{`C:`, `D:`}},
	}
	response := <-out
	if response.Err == nil {
		t.Fatalf(`expected an error but got none`)
	}
}

func TestGetIcons(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    ``,
		Function:   `GetToolData`,
		Parameters: params{},
		Config: &config.Config{
			ToolData: []tool_data_loader.ToolData{
				{
					Plugin:  "Test",
					Inputs:  []string{},
					Outputs: []string{},
					Icon:    "I am a picture in base64 encoding!",
				},
			},
		},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	toolData := response.Response.([]tool_data_loader.ToolData)
	if len(toolData) != 1 {
		t.Fatalf(`expected 1 tool but got %v`, len(toolData))
	}
	if toolData[0].Plugin != `Test` {
		t.Fatalf(`expected tool plugin of 'Test' but got '%v'`, toolData[0].Plugin)
	}
	if toolData[0].Icon != `I am a picture in base64 encoding!` {
		t.Fatalf(`expected tool icon of 'I am a picture in base64 encoding!' but got '%v'`, toolData[0].Icon)
	}
}

func TestGetEmptyWorkflow(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `New Workflow 1.yxwz`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{Project: workFolder, Function: `GetDocumentStructure`, Parameters: map[string]interface{}{`FilePath`: doc}, Out: out, Config: &config.Config{}}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err)
	}
	structure := response.Response.(cop.DocumentStructure)
	if structure.Connections == nil {
		t.Fatalf(`expected empty list of connections, not nil`)
	}
	if structure.Nodes == nil {
		t.Fatalf(`expected empty list of nodes, not nil`)
	}
}

func TestWhereUsed(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `Calculate Filter Expression.yxmc`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{Project: workFolder, Function: `WhereUsed`, Parameters: params{`FilePath`: doc}, Out: out, Config: &config.Config{}}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	whereUsed := response.Response.([]string)
	if len(whereUsed) != 1 {
		t.Fatalf(`expected 1 where used but got %v`, len(whereUsed))
	}
}

func TestRenameFiles(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	from, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `Calculate Filter Expression.yxmc`))
	to, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `macros`, `Calculate Filter Expression.yxmc`))
	fromFiles := []interface{}{from}
	toFiles := []interface{}{to}
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   `RenameFiles`,
		Parameters: params{`From`: fromFiles, `To`: toFiles},
		Config:     &config.Config{},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	_, err := os.Stat(from)
	if !os.IsNotExist(err) {
		t.Fatalf(`expected 'not exists' error but got: %v`, err)
	}
	_, err = os.Stat(to)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
}

func TestMoveFiles(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	file1, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `Calculate Filter Expression.yxmc`))
	file2, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `MultiInOut.yxmc`))
	files := []interface{}{
		file1,
		file2,
	}
	moveTo, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `macros`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   `MoveFiles`,
		Parameters: params{`Files`: files, `MoveTo`: moveTo},
		Config:     &config.Config{},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	errs := response.Response.([]string)
	if count := len(errs); count > 0 {
		t.Fatalf(`expected no errors but got %v`, count)
	}
}

func TestMoveFilesInvalidFilesParam(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	files := []interface{}{1, 2}
	moveTo, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `macros`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   `MoveFiles`,
		Parameters: params{`Files`: files, `MoveTo`: moveTo},
		Config:     &config.Config{},
	}
	response := <-out
	if response.Err == nil {
		t.Fatalf(`expected an error but got none`)
	}
}

func TestAnchorOrder(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `MultiInOut.yxmd`))
	var macro, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `MultiInOut.yxmc`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	for i := 0; i < 100; i++ {
		in <- cop.FunctionCall{
			Project:    workFolder,
			Function:   `GetDocumentStructure`,
			Parameters: params{`FilePath`: doc},
			Out:        out,
			Config: &config.Config{
				ToolData: []tool_data_loader.ToolData{},
			},
		}
		response := <-out
		if response.Err != nil {
			t.Fatalf(`expected no error but got: %v`, response.Err.Error())
		}
		structure := response.Response.(cop.DocumentStructure)
		found := 0
		for _, data := range structure.MacroToolData {
			if data.Plugin != macro {
				continue
			}
			found++
			if ins := len(data.Inputs); ins != 2 {
				t.Fatalf(`expected 2 inputs on iteration %v but got %v`, i, ins)
			}
			if outs := len(data.Inputs); outs != 2 {
				t.Fatalf(`expected 2 outputs on iteration %v but got %v`, i, outs)
			}
			if data.Inputs[0] != `Input2` {
				t.Fatalf(`expected the first input to be 'Input2' on iteration %v but got '%v'`, i, data.Inputs[0])
			}
			if data.Inputs[1] != `Input3` {
				t.Fatalf(`expected the second input to be 'Input3' on iteration %v but got '%v'`, i, data.Inputs[1])
			}
			if data.Outputs[0] != `Output4` {
				t.Fatalf(`expected the first output to be 'Output4' on iteration %v but got '%v'`, i, data.Outputs[0])
			}
			if data.Outputs[1] != `Output5` {
				t.Fatalf(`expected the second output to be 'Output5' on iteration %v but got '%v'`, i, data.Outputs[1])
			}
		}
		if found == 0 {
			t.Fatalf(`expected to find MultiInOut.yxmc on iteration %v but did not`, i)
		}
	}
}

func TestMakeFilesAbsolute(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var file, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `MultiInOut.yxmc`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   "MakeFilesAbsolute",
		Parameters: params{"Files": []interface{}{file}},
		Config:     &config.Config{},
	}

	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	changed := response.Response.(int)
	if changed != 1 {
		t.Fatalf(`expected 1 document to be changed but got %v`, changed)
	}
}

func TestMakeFilesAbsoluteMissingMacro(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   "MakeFilesAbsolute",
		Parameters: params{},
		Config:     &config.Config{},
	}

	response := <-out
	if response.Err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	t.Logf(response.Err.Error())
}

func TestMakeFilesRelative(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	var file, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `MultiInOut.yxmc`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   "MakeFilesRelative",
		Parameters: params{"Files": []interface{}{file}},
		Config:     &config.Config{},
	}

	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	changed := response.Response.(int)
	if changed != 1 {
		t.Fatalf(`expected 1 document to be changed but got %v`, changed)
	}
}

func TestMakeFilesRelativeMissingFiles(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   "MakeFilesRelative",
		Parameters: params{},
		Config:     &config.Config{},
	}

	response := <-out
	if response.Err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	t.Logf(response.Err.Error())
}

func TestMakeAllFilesRelative(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   "MakeAllFilesRelative",
		Parameters: params{},
		Config:     &config.Config{},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	changed := response.Response.(int)
	if changed != 2 {
		t.Fatalf(`expected 2 changed documents but got %v`, changed)
	}
}

func TestMakeAllFilesAbsolute(t *testing.T) {
	rebuildTestDocs()
	defer rebuildTestDocs()

	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Out:        out,
		Project:    workFolder,
		Function:   "MakeAllFilesAbsolute",
		Parameters: params{},
		Config:     &config.Config{},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	changed := response.Response.(int)
	if changed != 2 {
		t.Fatalf(`expected 2 changed documents but got %v`, changed)
	}
}

func TestInterfaceNodesDocumentStructure(t *testing.T) {
	var doc, _ = filepath.Abs(filepath.Join(`..`, `testdocs`, `Interface.yxmc`))
	in := make(chan cop.FunctionCall)
	out := make(chan cop.FunctionResponse)
	go cop.StartTrafficCop(in)

	in <- cop.FunctionCall{
		Project:    workFolder,
		Function:   `GetDocumentStructure`,
		Parameters: params{`FilePath`: doc},
		Out:        out,
		Config: &config.Config{
			ToolData: []tool_data_loader.ToolData{},
		},
	}
	response := <-out
	if response.Err != nil {
		t.Fatalf(`expected no error but got: %v`, response.Err.Error())
	}
	if response.Response == nil {
		t.Fatalf(`expected a non-nil response but got nil`)
	}
	structure := response.Response.(cop.DocumentStructure)
	if count := len(structure.Nodes); count != 40 {
		t.Fatalf(`expected 40 nodes but got %v`, count)
	}
	encoded, _ := json.Marshal(response)
	t.Logf(string(encoded))
}

func rebuildTestDocs() {
	testdocbuilder.RebuildTestdocs(filepath.Join(`..`, `testdocs`))
}
