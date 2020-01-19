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
	data *txml.Node
}

type Position struct {
	X      int
	Y      int
	Width  float64
	Height float64
}

type MacroPath struct {
	StoredPath string
	FoundPath  string
	RelativeTo string
}

func (ryxNode *RyxNode) ReadCategory() Category {
	dll := ryxNode.data.First(`EngineSettings`).Attributes[`EngineDll`]
	macro := ryxNode.data.First(`EngineSettings`).Attributes[`Macro`]
	plugin := ryxNode.data.First(`GuiSettings`).Attributes[`Plugin`]
	cosmeticPlugins := &h.StringArray{`AlteryxGuiToolkit.TextBox.TextBox`, `AlteryxGuiToolkit.HtmlBox.HtmlBox`}

	if dll != "" {
		return Tool
	}
	if macro != "" {
		return Macro
	}
	if cosmeticPlugins.Contains(plugin) {
		return Cosmetic
	}
	if plugin == `AlteryxGuiToolkit.ToolContainer.ToolContainer` {
		return Container
	}
	return Invalid
}

func (ryxNode *RyxNode) ReadId() (int, error) {
	return strconv.Atoi(ryxNode.data.Attributes[`ToolID`])
}

func (ryxNode *RyxNode) ReadPosition() (Position, error) {
	gui := ryxNode.data.First(`GuiSettings`).First(`Position`)
	x, err := strconv.Atoi(gui.Attributes[`x`])
	if err != nil {
		return Position{}, err
	}
	y, err := strconv.Atoi(gui.Attributes[`y`])
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
	return ryxNode.data.First(`GuiSettings`).Attributes[`Plugin`]
}

func (ryxNode *RyxNode) SetPosition(x int, y int) {
	setting := ryxNode.data.First(`GuiSettings`).First(`Position`)
	setting.Attributes = map[string]string{`x`: strconv.Itoa(x), `y`: strconv.Itoa(y)}
}

func (ryxNode *RyxNode) ReadMacro(macroPaths ...string) MacroPath {
	stored := ryxNode.data.First(`EngineSettings`).Attributes[`Macro`]
	if stored == `` {
		return MacroPath{StoredPath: ``, FoundPath: ``}
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
	setting := ryxNode.data.First(`EngineSettings`)
	setting.Attributes = map[string]string{`Macro`: winMacro}
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
	for _, child := range ryxNode.data.First(`ChildNodes`).Nodes {
		node := New(child)
		list = append(list, node)
		list = append(list, node.ReadChildren()...)
	}
	return list
}

func (ryxNode *RyxNode) RemoveChildren(ids ...int) {
	currentIndex := 0
	container := ryxNode.data.First(`ChildNodes`)
	for _, childXml := range container.Nodes {
		childNode := New(childXml)
		if !childNode.MatchesIds(ids...) {
			container.Nodes[currentIndex] = childXml
			childNode.RemoveChildren(ids...)
			currentIndex += 1
		}
	}
	container.Nodes = container.Nodes[0:currentIndex]
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

func (ryxNode *RyxNode) ReadData() *txml.Node {
	return ryxNode.data
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

func (ryxNode *RyxNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(ryxNode.data, start)
}

func (ryxNode *RyxNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	data := &txml.Node{}
	err := d.DecodeElement(data, &start)
	if err == nil {
		ryxNode.data = data
	}
	return err
}

func GenerateNodeFromXml(xmlString string) (*RyxNode, error) {
	ayxNode := &RyxNode{}
	err := xml.Unmarshal([]byte(xmlString), ayxNode)
	return ayxNode, err
}

func New(node *txml.Node) *RyxNode {
	return &RyxNode{node}
}

func NewMacroXml(id int, path string, x int, y int) *txml.Node {
	path = strings.Replace(path, string(os.PathSeparator), `\`, -1)

	return &txml.Node{
		Name: `Node`,
		Attributes: map[string]string{
			`ToolID`: strconv.Itoa(id),
		},
		Nodes: []*txml.Node{
			{
				Name: `GuiSettings`,
				Nodes: []*txml.Node{
					{
						Name: `Position`,
						Attributes: map[string]string{
							`x`: strconv.Itoa(x),
							`y`: strconv.Itoa(y),
						},
					},
				},
			},
			{
				Name: `Properties`,
				Nodes: []*txml.Node{
					{
						Name: `Configuration`,
					},
					{
						Name:       `Annotation`,
						Attributes: map[string]string{`DisplayMode`: `0`},
						Nodes: []*txml.Node{
							{Name: `Name`},
							{Name: `DefaultAnnotationText`},
							{Name: `Left`, Attributes: map[string]string{`value`: `False`}},
						},
					},
				},
			},
			{
				Name: `EngineSettings`,
				Attributes: map[string]string{
					`Macro`: path,
				},
			},
		},
	}
}
