package ryxnode

import (
	"encoding/xml"
	"errors"
	h "github.com/tlarsen7572/Golang-Public/helpers"
	"github.com/tlarsen7572/Golang-Public/txml"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type RyxNode struct {
	ToolId         string
	GuiSettings    *txml.Node
	Properties     *Properties
	EngineSettings *txml.Node
	ChildNodes     []*RyxNode
}

type Properties struct {
	Configuration Configuration `xml:"Configuration"`
	Annotation    *txml.Node    `xml:"Annotation"`
}

type Configuration struct {
	InnerXml string `xml:",innerxml"`
}

type Position struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type MacroPath struct {
	StoredPath string
	FoundPath  string
	RelativeTo string
}

func (ryxNode *RyxNode) ReadCategory() Category {
	dll := ryxNode.EngineSettings.Attributes[`EngineDll`]
	if dll != "" {
		return Tool
	}

	macro := ryxNode.EngineSettings.Attributes[`Macro`]
	if macro != "" {
		return Macro
	}

	cosmeticPlugins := &h.StringArray{`AlteryxGuiToolkit.TextBox.TextBox`, `AlteryxGuiToolkit.HtmlBox.HtmlBox`}
	plugin := ryxNode.GuiSettings.Attributes[`Plugin`]
	if cosmeticPlugins.Contains(plugin) {
		return Cosmetic
	}

	if plugin == `AlteryxGuiToolkit.ToolContainer.ToolContainer` {
		return Container
	}

	return Invalid
}

func (ryxNode *RyxNode) ReadId() (int, error) {
	return strconv.Atoi(ryxNode.ToolId)
}

func (ryxNode *RyxNode) ReadPosition() (Position, error) {
	gui := ryxNode.GuiSettings.First(`Position`)
	x, err := strconv.ParseFloat(gui.Attributes[`x`], 64)
	if err != nil {
		return Position{}, err
	}
	y, err := strconv.ParseFloat(gui.Attributes[`y`], 64)
	if err != nil {
		return Position{}, err
	}
	width, err := strconv.ParseFloat(gui.Attributes[`width`], 64)
	if err != nil {
		width = 60
	}
	height, err := strconv.ParseFloat(gui.Attributes[`height`], 64)
	if err != nil {
		height = 60
	}
	return Position{X: x, Y: y, Width: width, Height: height}, nil
}

func (ryxNode *RyxNode) ReadPlugin() string {
	return ryxNode.GuiSettings.Attributes[`Plugin`]
}

func (ryxNode *RyxNode) SetPosition(x float64, y float64) {
	setting := ryxNode.GuiSettings.First(`Position`)
	setting.Attributes = map[string]string{`x`: h.DblToStr(x, 0), `y`: h.DblToStr(y, 0)}
}

func (ryxNode *RyxNode) ReadMacro(macroPaths ...string) MacroPath {
	stored := ryxNode.EngineSettings.Attributes[`Macro`]
	if stored == `` {
		return MacroPath{StoredPath: ``, FoundPath: ``, RelativeTo: ``}
	}

	osStored := strings.Replace(stored, `\`, string(os.PathSeparator), -1)
	if _, err := os.Stat(osStored); err == nil {
		return MacroPath{StoredPath: stored, FoundPath: osStored}
	}
	for _, macroPath := range macroPaths {
		absolute := filepath.Join(macroPath, osStored)
		if _, err := os.Stat(absolute); err == nil {
			return MacroPath{StoredPath: stored, FoundPath: absolute, RelativeTo: macroPath}
		}
	}
	return MacroPath{StoredPath: stored, FoundPath: ``}
}

func (ryxNode *RyxNode) SetMacro(macro string) {
	winMacro := strings.Replace(macro, string(os.PathSeparator), `\`, -1)
	ryxNode.EngineSettings.Attributes = map[string]string{`Macro`: winMacro}
}

func (ryxNode *RyxNode) MakeMacroAbsolute(macroPaths ...string) error {
	path := ryxNode.ReadMacro(macroPaths...)
	if path.FoundPath != `` {
		ryxNode.SetMacro(path.FoundPath)
		return nil
	}
	return errors.New(`no valid macro path was found`)
}

func (ryxNode *RyxNode) MakeMacroRelative(to string, macroPaths ...string) error {
	path := ryxNode.ReadMacro(macroPaths...)
	if path.FoundPath != `` {
		relPath, err := filepath.Rel(to, path.FoundPath)
		if err != nil {
			return err
		}
		ryxNode.SetMacro(relPath)
		return nil
	}
	return errors.New(`no valid macro path was found`)
}

func (ryxNode *RyxNode) ReadChildren() []*RyxNode {
	var list []*RyxNode
	for _, child := range ryxNode.ChildNodes {
		list = append(list, child)
		list = append(list, child.ReadChildren()...)
	}
	return list
}

func (ryxNode *RyxNode) RemoveChildren(ids ...int) {
	currentIndex := 0
	for _, childNode := range ryxNode.ChildNodes {
		if !childNode.MatchesIds(ids...) {
			ryxNode.ChildNodes[currentIndex] = childNode
			childNode.RemoveChildren(ids...)
			currentIndex += 1
		}
	}
	ryxNode.ChildNodes = ryxNode.ChildNodes[0:currentIndex]
}

func (ryxNode *RyxNode) MatchesIds(ids ...int) bool {
	nodeId, err := ryxNode.ReadId()
	if err != nil {
		return false
	}
	for _, id := range ids {
		if nodeId == id {
			return true
		}
	}
	return false
}

type Category int

const (
	Invalid   Category = 0
	Tool      Category = 1
	Cosmetic  Category = 2
	Macro     Category = 3
	Container Category = 4
)

var categoryNames = []string{`Invalid`, `Tool`, `Cosmetic`, `Macro`, `Container`}

func (cat Category) String() string {
	return categoryNames[cat]
}

func NewMacro(id int, path string, x float64, y float64) *RyxNode {
	return &RyxNode{
		ToolId: strconv.Itoa(id),
		GuiSettings: &txml.Node{
			Name: "GuiSettings",
			Nodes: []*txml.Node{
				{
					Name: `Position`,
					Attributes: map[string]string{
						`x`: h.DblToStr(x, 0),
						`y`: h.DblToStr(y, 0),
					},
				},
			},
		},
		Properties: &Properties{
			Configuration: Configuration{InnerXml: "<Configuration />"},
			Annotation: &txml.Node{
				Name:       `Annotation`,
				Attributes: map[string]string{`DisplayMode`: `0`},
				Nodes: []*txml.Node{
					{Name: `Name`},
					{Name: `DefaultAnnotationText`},
					{Name: `Left`, Attributes: map[string]string{`value`: `False`}},
				},
			},
		},
		EngineSettings: &txml.Node{
			Name: `EngineSettings`,
			Attributes: map[string]string{
				`Macro`: path,
			},
		},
	}
}

func GenerateNodeFromXml(xmlString string) (*RyxNode, error) {
	ayxNode := &RyxNode{}
	err := xml.Unmarshal([]byte(xmlString), ayxNode)
	return ayxNode, err
}
