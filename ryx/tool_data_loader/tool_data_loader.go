package tool_data_loader

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/tlarsen7572/Golang-Public/ryx/ini_reader"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"github.com/tlarsen7572/Golang-Public/txml"
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type DllToolData struct {
	Error    string
	ToolData []ToolData
}

type ToolData struct {
	Plugin  string
	Inputs  []string
	Outputs []string
	Icon    string
}

func LoadAll(programFiles string, programData string) ([]ToolData, error) {
	javascriptPluginFolders := []string{
		filepath.Join(programData, `Tools`),
		filepath.Join(programFiles, `bin`, `HtmlPlugins`),
	}
	macroPluginsFolder := filepath.Join(programFiles, `Settings`, `AddOnData`, `Macros`)
	runtimeMacrosFolder := filepath.Join(programFiles, `bin`, `RuntimeData`, `Macros`)

	tools, dllErr := LoadDllTools()
	javascriptPluginTools, javascriptErr := LoadJavascriptPluginTools(javascriptPluginFolders...)
	macroPluginTools, macroPluginErr := LoadMacroPluginTools(macroPluginsFolder)
	runtimeMacroTools, runtimeMacroErr := LoadMacroTools(runtimeMacrosFolder)
	tools = append(tools, javascriptPluginTools...)
	tools = append(tools, macroPluginTools...)
	tools = append(tools, runtimeMacroTools...)
	err := chooseErr(dllErr, javascriptErr, macroPluginErr, runtimeMacroErr)
	return tools, err
}

func chooseErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadDllTools() ([]ToolData, error) {
	cmd := exec.Command(path.Join(`.`, `IconLoaderApp.exe`))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	response := &DllToolData{}
	err = json.Unmarshal(out.Bytes(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != `` {
		err = errors.New(response.Error)
	}

	return response.ToolData, err
}

func LoadJavascriptPluginTools(javascriptPluginFolders ...string) ([]ToolData, error) {
	tools := make([]ToolData, 0)
	var finalError error
	for _, toolFolder := range javascriptPluginFolders {
		things, err := ioutil.ReadDir(toolFolder)
		if err != nil {
			finalError = err
			continue
		}
		for _, thing := range things {
			if thing.IsDir() {
				tool, err := readJavascriptPluginTool(filepath.Join(toolFolder, thing.Name()))
				if err == nil {
					tools = append(tools, tool)
				}
			}
		}
	}
	return tools, finalError
}

func LoadMacroPluginTools(pluginMacroInisPath string) ([]ToolData, error) {
	fileInfos, err := ioutil.ReadDir(pluginMacroInisPath)
	if err != nil {
		return nil, err
	}
	folders := make([]string, 0)
	for _, fileInfo := range fileInfos {
		if filepath.Ext(fileInfo.Name()) != `.ini` {
			continue
		}
		filePath := filepath.Join(pluginMacroInisPath, fileInfo.Name())
		mapped := ini_reader.LoadIni(filePath)
		folder, ok := mapped[`Path`]
		if ok {
			folders = append(folders, folder)
		}
	}
	return LoadMacroTools(folders...)
}

func LoadMacroTools(macroPaths ...string) ([]ToolData, error) {
	tools := make([]ToolData, 0)
	var returnErr error
	for _, macroPath := range macroPaths {
		folderTools, err := processFolder(macroPath, macroPath)
		tools = append(tools, folderTools...)
		if err != nil {
			returnErr = err
		}
	}
	return tools, returnErr
}

func readJavascriptPluginTool(dir string) (ToolData, error) {
	folderName := filepath.Base(dir)
	configFileName := filepath.Join(dir, folderName+`Config.xml`)
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return ToolData{}, err
	}
	xml, err := txml.Parse(string(content))
	if err != nil {
		return ToolData{}, err
	}
	iconFile := xml.First(`GuiSettings`).Attributes[`Icon`]

	iconBytes, err := ioutil.ReadFile(filepath.Join(dir, iconFile))
	var iconStr string
	if err == nil {
		iconStr = base64.StdEncoding.EncodeToString(iconBytes)
	}

	inputs := make([]string, 0)
	outputs := make([]string, 0)
	for _, node := range xml.First(`GuiSettings`).First(`InputConnections`).Nodes {
		inputs = append(inputs, node.Attributes[`Name`])
	}
	for _, node := range xml.First(`GuiSettings`).First(`OutputConnections`).Nodes {
		outputs = append(outputs, node.Attributes[`Name`])
	}
	toolData := ToolData{
		Plugin:  folderName,
		Inputs:  inputs,
		Outputs: outputs,
		Icon:    iconStr,
	}
	return toolData, nil
}

func processFolder(path string, macroPath string) ([]ToolData, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	returnTools := make([]ToolData, 0)
	var returnErr error
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		subPath := filepath.Join(path, name)
		if fileInfo.IsDir() {
			tools, err := processFolder(subPath, macroPath)
			if err != nil {
				returnErr = err
			}
			returnTools = append(returnTools, tools...)
			continue
		}
		if filepath.Ext(name) != `.yxmc` {
			continue
		}
		tool, err := ReadSingleMacro(subPath, macroPath)
		if err == nil {
			returnTools = append(returnTools, tool)
			continue
		}
		returnErr = err
	}
	return returnTools, returnErr
}

func ReadSingleMacro(path string, macroPath string) (ToolData, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return ToolData{}, err
	}
	macroXml, err := txml.Parse(string(contents))
	if err != nil {
		return ToolData{}, err
	}
	macro, err := ryxdoc.ReadFile(path)
	if err != nil {
		return ToolData{}, err
	}
	plugin, _ := filepath.Rel(macroPath, path)
	if macroPath == `` {
		plugin = path
	}
	icon := macroXml.First(`Properties`).First(`RuntimeProperties`).First(`MacroImage`).InnerText
	icon = strings.Replace(icon, "\n", ``, -1)
	inputs := make([]string, 0)
	outputs := make([]string, 0)
	for _, node := range macro.ReadNodes() {
		if node.ReadPlugin() == `AlteryxBasePluginsGui.MacroOutput.MacroOutput` {
			name := node.ReadData().First(`Properties`).First(`Configuration`).First(`Name`).InnerText
			outputs = append(outputs, name)
		}
		if node.ReadPlugin() == `AlteryxBasePluginsGui.MacroInput.MacroInput` {
			name := node.ReadData().First(`Properties`).First(`Configuration`).First(`Name`).InnerText
			inputs = append(inputs, name)
		}
	}
	return ToolData{
		Plugin:  plugin,
		Inputs:  inputs,
		Outputs: outputs,
		Icon:    icon,
	}, nil
}
