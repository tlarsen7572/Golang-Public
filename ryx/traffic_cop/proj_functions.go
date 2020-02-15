package traffic_cop

import (
	"errors"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
	"github.com/tlarsen7572/Golang-Public/ryx/tool_data_loader"
	"path/filepath"
)

type NodeStructure struct {
	ToolId      int
	X           float64
	Y           float64
	Width       float64
	Height      float64
	Plugin      string
	StoredMacro string
	FoundMacro  string
	Category    string
}

type DocumentStructure struct {
	Nodes         []NodeStructure
	Connections   []*ryxdoc.RyxConn
	MacroToolData []tool_data_loader.ToolData
}

const getProjectStructureFunc = `GetProjectStructure`
const getDocumentStructureFunc = `GetDocumentStructure`
const whereUsedFunc = `WhereUsed`
const renameFileFunc = `RenameFile`
const makeMacroAbsoluteFunc = `MakeMacroAbsolute`
const makeMacroRelativeFunc = `MakeMacroRelative`
const makeAllRelativeFunc = `MakeAllMacrosRelative`
const makeAllAbsoluteFunc = `MakeAllMacrosAbsolute`
const invalidProjFunc = `invalid project function`

func handleProjFunction(call FunctionCall, data *TrafficCopData) FunctionResponse {
	switch call.Function {
	case getProjectStructureFunc:
		return getProjectStructure(data)
	case getDocumentStructureFunc:
		return getDocumentStructure(call, data)
	case whereUsedFunc:
		return whereUsed(call, data)
	case renameFileFunc:
		return renameFile(call, data)
	case makeMacroAbsoluteFunc:
		return makeMacroAbsolute(call, data)
	case makeMacroRelativeFunc:
		return makeMacroRelative(call, data)
	case makeAllRelativeFunc:
		return makeAllRelative(data)
	case makeAllAbsoluteFunc:
		return makeAllAbsolute(data)
	default:
		return FunctionResponse{errors.New(invalidProjFunc), nil}
	}
}

func getProjectStructure(data *TrafficCopData) FunctionResponse {
	structure, err := data.Project.Structure()
	if err != nil {
		return FunctionResponse{Err: err, Response: nil}
	}
	return FunctionResponse{nil, structure}
}

func getDocumentStructure(call FunctionCall, data *TrafficCopData) FunctionResponse {
	var filePath string
	var ok bool
	if filePath, ok = call.Parameters[`FilePath`]; !ok {
		return FunctionResponse{errors.New(`the FilePath parameter was not included`), nil}
	}

	doc, err := data.Project.RetrieveDocument(filePath)
	if err != nil {
		return FunctionResponse{err, nil}
	}

	folderPath := filepath.Dir(filePath)
	macroPaths := append(data.MacroPaths, folderPath)

	nodes := []NodeStructure{}
	toolData := []tool_data_loader.ToolData{}
	for _, node := range doc.Nodes {
		id, err := node.ReadId()
		if err != nil {
			continue
		}
		plugin := node.ReadPlugin()
		if plugin == `AlteryxGuiToolkit.Questions.Tab.Tab` {
			continue
		}
		position, err := node.ReadPosition()
		if err != nil {
			position = ryxnode.Position{X: 0, Y: 0, Width: 0, Height: 0}
		}
		macro := node.ReadMacro(macroPaths...)
		category := node.ReadCategory().String()
		nodes = append(nodes, NodeStructure{
			ToolId:      id,
			X:           position.X,
			Y:           position.Y,
			Width:       position.Width,
			Height:      position.Height,
			Plugin:      plugin,
			StoredMacro: macro.StoredPath,
			FoundMacro:  macro.FoundPath,
			Category:    category,
		})
		if plugin := macro.StoredPath; plugin != `` && macro.FoundPath != `` {
			needsToolData := true
			for _, existing := range call.Config.ToolData {
				if plugin == existing.Plugin {
					needsToolData = false
					break
				}
			}
			if needsToolData {
				for _, alreadyGotIt := range toolData {
					if alreadyGotIt.Plugin == macro.FoundPath {
						needsToolData = false
						break
					}
				}
			}
			if needsToolData {
				data, err := tool_data_loader.ReadSingleMacro(macro.FoundPath, ``)
				if err == nil {
					toolData = append(toolData, data)
				}
			}
		}
	}

	connections := doc.Connections
	if connections == nil {
		connections = []*ryxdoc.RyxConn{}
	}

	docStructure := DocumentStructure{
		Nodes:         nodes,
		Connections:   connections,
		MacroToolData: toolData,
	}
	return FunctionResponse{nil, docStructure}
}

func whereUsed(call FunctionCall, data *TrafficCopData) FunctionResponse {
	path, ok := call.Parameters[`FilePath`]
	if !ok {
		return FunctionResponse{errors.New(`the FilePath parameter was not included`), nil}
	}
	whereUsed := data.Project.WhereUsed(path)
	return FunctionResponse{nil, whereUsed}
}

func makeMacroAbsolute(call FunctionCall, data *TrafficCopData) FunctionResponse {
	macro, ok := call.Parameters[`Macro`]
	if !ok {
		return FunctionResponse{
			Err:      errors.New(`the Macro parameter was not included`),
			Response: nil,
		}
	}
	result := data.Project.MakeMacroAbsolute(macro)
	return FunctionResponse{
		Err:      nil,
		Response: result,
	}
}

func makeMacroRelative(call FunctionCall, data *TrafficCopData) FunctionResponse {
	macro, ok := call.Parameters[`Macro`]
	if !ok {
		return FunctionResponse{
			Err:      errors.New(`the Macro parameter was not included`),
			Response: nil,
		}
	}
	result := data.Project.MakeMacroRelative(macro)
	return FunctionResponse{
		Err:      nil,
		Response: result,
	}
}

func makeAllRelative(data *TrafficCopData) FunctionResponse {
	result := data.Project.MakeAllMacrosRelative()
	return FunctionResponse{
		Err:      nil,
		Response: result,
	}
}

func makeAllAbsolute(data *TrafficCopData) FunctionResponse {
	result := data.Project.MakeAllMacrosAbsolute()
	return FunctionResponse{
		Err:      nil,
		Response: result,
	}
}

func renameFile(call FunctionCall, data *TrafficCopData) FunctionResponse {
	from, ok := call.Parameters[`From`]
	if !ok {
		return FunctionResponse{
			Err:      errors.New(`the From parameter was not included`),
			Response: nil,
		}
	}
	to, ok := call.Parameters[`To`]
	if !ok {
		return FunctionResponse{
			Err:      errors.New(`the To parameter was not included`),
			Response: nil,
		}
	}
	err := data.Project.RenameFile(from, to)
	return FunctionResponse{
		Err:      err,
		Response: nil,
	}
}
