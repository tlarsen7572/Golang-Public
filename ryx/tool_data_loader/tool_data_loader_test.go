package tool_data_loader_test

import (
	"encoding/json"
	loader "github.com/tlarsen7572/Golang-Public/ryx/tool_data_loader"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestLoadDllTools(t *testing.T) {
	if runtime.GOOS != `windows` {
		t.Skipf(`test must be run on Windows with Alteryx installed`)
	}

	toolData, err := loader.LoadDllTools()
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if len(toolData) == 0 {
		t.Fatalf(`expected 1 or more tools but got none`)
	}
	t.Logf(`count: %v`, len(toolData))
}

func TestLoadJavascriptPluginTools(t *testing.T) {
	if runtime.GOOS != `windows` {
		t.Skipf(`test must be run on Windows with Alteryx installed`)
	}

	toolData, err := loader.LoadJavascriptPluginTools(`C:\ProgramData\Alteryx\Tools`, `C:\Program Files\Alteryx\bin\HtmlPlugins`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if len(toolData) == 0 {
		t.Fatalf(`expected 1 or more tools but got none`)
	}
	t.Logf(`count: %v`, len(toolData))
}

func TestLoadMacroTools(t *testing.T) {
	macroPath := filepath.Join(`..`, `testdocs`)
	toolData, err := loader.LoadMacroTools(macroPath)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if len(toolData) == 0 {
		t.Fatalf(`expected 1 or more tools but got none`)
	}
	t.Logf(`count: %v`, len(toolData))
}

func TestLoadMacroPlugins(t *testing.T) {
	if runtime.GOOS != `windows` {
		t.Skipf(`test must be run on Windows with Alteryx installed`)
	}

	toolData, err := loader.LoadMacroPluginTools(`C:\Program Files\Alteryx\Settings\AddOnData\Macros`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if len(toolData) == 0 {
		t.Fatalf(`expected 1 or more tools but got none`)
	}
	t.Logf(`count: %v`, len(toolData))
}

func TestCombine(t *testing.T) {
	if runtime.GOOS != `windows` {
		t.Skipf(`test must be run on Windows with Alteryx installed`)
	}

	tools, err := loader.LoadAll(`C:\Program Files\Alteryx`, `C:\ProgramData\Alteryx`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if len(tools) == 0 {
		t.Fatalf(`expected 1 or more tools but got none`)
	}

	content, _ := json.Marshal(tools)
	err = ioutil.WriteFile("combined.json", content, 0644)
	if err != nil {
		t.Fatalf(`expected no error writing combined.json but got: %v`, err.Error())
	}
}

func TestLoadMacroWithNewlinesInImage(t *testing.T) {
	path := filepath.Join(`..`, `testdocs`, `Calculate Filter Expression.yxmc`)
	toolData, err := loader.ReadSingleMacro(path, ``)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if strings.Contains(toolData.Icon, "\n") {
		t.Fatalf(`Icon contains newline characters, but these should be removed`)
	}
}

func TestLoadMacroAnchorOrder(t *testing.T) {
	path := filepath.Join(`..`, `testdocs`, `MultiInOut.yxmc`)

	for i := 0; i < 100; i++ {
		data, err := loader.ReadSingleMacro(path, ``)
		if err != nil {
			t.Fatalf(`expected no error but got: %v`, err.Error())
		}
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
}
