package ryxnode_test

import (
	"encoding/xml"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const tool = `<RyxNode ToolID="1">
	<GuiSettings Plugin="AlteryxBasePluginsGui.DbFileInput.DbFileInput">
		<Position x="318" y="222" />
	</GuiSettings>
	<Properties>
		<Configuration />
		<Annotation DisplayMode="1">
			<Name>My Annotation</Name>
			<DefaultAnnotationText>Default Annotation</DefaultAnnotationText>
			<Left value="True" />
		</Annotation>
		<MetaInfo />
	</Properties>
	<EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxDbFileInput" />
</RyxNode>`

const macro = `    <RyxNode ToolID="22">
      <GuiSettings>
        <Position x="258" y="846" />
      </GuiSettings>
      <Properties>
        <Configuration />
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <EngineSettings Macro="_Utilities\ChooseText.yxmc" />
    </RyxNode>`

const cosmetic = `    <RyxNode ToolID="1">
      <GuiSettings Plugin="AlteryxGuiToolkit.TextBox.TextBox">
        <Position x="150" y="162" width="100" height="40" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Text />
          <Font name="Arial" size="8.25" style="0" />
          <TextColor name="Black" />
          <FillColor name="White" />
          <Shape shape="0" />
          <Justification Justification="4" />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
    </RyxNode>`

const container = `    <Node ToolID="20">
      <GuiSettings Plugin="AlteryxGuiToolkit.ToolContainer.ToolContainer">
        <Position x="269" y="292" width="405.0632" height="188" />
      </GuiSettings>
      <Properties>
        <Configuration>
          <Caption>Container 20</Caption>
          <Style TextColor="#314c4a" FillColor="#ecf2f2" BorderColor="#314c4a" Transparency="25" Margin="25" />
          <Disabled value="False" />
          <Folded value="False" />
        </Configuration>
        <Annotation DisplayMode="0">
          <Name />
          <DefaultAnnotationText />
          <Left value="False" />
        </Annotation>
      </Properties>
      <ChildNodes>
        <Node ToolID="21">
          <GuiSettings Plugin="AlteryxGuiToolkit.TextBox.TextBox">
            <Position x="294" y="378" width="100" height="40" />
          </GuiSettings>
          <Properties>
            <Configuration>
              <Text>One level in</Text>
              <Font name="Arial" size="8.25" style="0" />
              <TextColor name="Black" />
              <FillColor name="White" />
              <Shape shape="0" />
              <Justification Justification="4" />
            </Configuration>
            <Annotation DisplayMode="0">
              <Name />
              <DefaultAnnotationText />
              <Left value="False" />
            </Annotation>
          </Properties>
        </Node>
        <Node ToolID="22">
          <GuiSettings Plugin="AlteryxGuiToolkit.ToolContainer.ToolContainer">
            <Position x="497" y="341" width="152.0632" height="114" />
          </GuiSettings>
          <Properties>
            <Configuration>
              <Caption>Container 22</Caption>
              <Style TextColor="#314c4a" FillColor="#ecf2f2" BorderColor="#314c4a" Transparency="25" Margin="25" />
              <Disabled value="False" />
              <Folded value="False" />
            </Configuration>
            <Annotation DisplayMode="0">
              <Name />
              <DefaultAnnotationText />
              <Left value="False" />
            </Annotation>
          </Properties>
          <ChildNodes>
            <Node ToolID="23">
              <GuiSettings Plugin="AlteryxGuiToolkit.TextBox.TextBox">
                <Position x="522" y="390" width="100" height="40" />
              </GuiSettings>
              <Properties>
                <Configuration>
                  <Text>Two levels in</Text>
                  <Font name="Arial" size="8.25" style="0" />
                  <TextColor name="Black" />
                  <FillColor name="White" />
                  <Shape shape="0" />
                  <Justification Justification="4" />
                </Configuration>
                <Annotation DisplayMode="0">
                  <Name />
                  <DefaultAnnotationText />
                  <Left value="False" />
                </Annotation>
              </Properties>
            </Node>
          </ChildNodes>
        </Node>
      </ChildNodes>
    </Node>`

func TestInstantiateToolNode(t *testing.T) {
	node, err := ryxnode.GenerateNodeFromXml(tool)

	if err != nil {
		t.Fatalf("GenerateNode returned an error: %s", err.Error())
	}

	if id, _ := node.ReadId(); id != 1 {
		t.Fatalf("expected ReadId of 1 but got %v", id)
	}
	if category := node.ReadCategory(); category != ryxnode.Tool {
		t.Fatalf("expected catagory of '%v' but got '%v'", ryxnode.Tool, category)
	}
}

func TestInstantiateMacroNode(t *testing.T) {
	node, err := ryxnode.GenerateNodeFromXml(macro)

	if err != nil {
		t.Fatalf("GenerateNode returned an error: %s", err.Error())
	}

	if id, _ := node.ReadId(); id != 22 {
		t.Fatalf("expected ReadId of 22 but got %v", id)
	}
	if category := node.ReadCategory(); category != ryxnode.Macro {
		t.Fatalf("expected catagory of '%v' but got '%v'", ryxnode.Macro, category)
	}
}

func TestInstantiateCosmeticNode(t *testing.T) {
	node, err := ryxnode.GenerateNodeFromXml(cosmetic)

	if err != nil {
		t.Fatalf("GenerateNode returned an error: %s", err.Error())
	}

	if id, _ := node.ReadId(); id != 1 {
		t.Fatalf("expected ReadId of 1 but got %v", id)
	}
	if category := node.ReadCategory(); category != ryxnode.Cosmetic {
		t.Fatalf("expected catagory of '%v' but got '%v'", ryxnode.Cosmetic, category)
	}
}

func TestInstantiateContainerNode(t *testing.T) {
	node, err := ryxnode.GenerateNodeFromXml(container)

	if err != nil {
		t.Fatalf("GenerateNode returned an error: %s", err.Error())
	}

	if id, _ := node.ReadId(); id != 20 {
		t.Fatalf("expected ReadId of 20 but got %v", id)
	}
	if category := node.ReadCategory(); category != ryxnode.Container {
		t.Fatalf("expected catagory of '%v' but got '%v'", ryxnode.Container, category)
	}
}

func TestNodeToXml(t *testing.T) {
	oldXml := `<RyxNode ToolID="1"><GuiSettings Plugin="A"><Position x="318" y="222"></Position></GuiSettings><Properties><Configuration></Configuration><Annotation DisplayMode="1"><Name>My Annotation</Name><DefaultAnnotationText>Default Annotation</DefaultAnnotationText><Left value="True"></Left></Annotation></Properties><EngineSettings EngineDll="B.dll" EngineDllEntryPoint="C" Macro=""></EngineSettings><ChildNodes><RyxNode ToolID="4"><GuiSettings Plugin="AlteryxBasePluginsGui.Formula.Formula"><Position x="570" y="90"></Position></GuiSettings><Properties><Configuration></Configuration><Annotation DisplayMode="0"><Name></Name><DefaultAnnotationText></DefaultAnnotationText><Left value="False"></Left></Annotation></Properties><EngineSettings EngineDll="AlteryxBasePluginsEngine.dll" EngineDllEntryPoint="AlteryxFormula" Macro=""></EngineSettings><ChildNodes></ChildNodes></RyxNode></ChildNodes></RyxNode>`

	node, err := ryxnode.GenerateNodeFromXml(oldXml)
	newXmlBytes, err := xml.Marshal(node)
	newXml := string(newXmlBytes)

	if err != nil {
		t.Fatalf(`expected no error but not '%v'`, err.Error())
	}
	if strings.ToUpper(oldXml) != strings.ToUpper(newXml) {
		t.Fatalf(`expected '%v' but got '%v'`, newXml, oldXml)
	}
	t.Logf(newXml)
}

func TestReadPosition(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(tool)
	pos, err := node.ReadPosition()
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	if pos.X != 318 || pos.Y != 222 {
		t.Fatalf(`expected position of 318,222 but got %v,%v`, pos.X, pos.Y)
	}
	if pos.Width != 60 {
		t.Fatalf(`expected width of 60 but got %v`, pos.Width)
	}
	if pos.Height != 60 {
		t.Fatalf(`expected height of 60 but got %v`, pos.Height)
	}
}

func TestReadContainerPosition(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(container)
	pos, err := node.ReadPosition()
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	if pos.Width != 405.0632 {
		t.Fatalf(`expected width of 405.0632 but got %v`, pos.Width)
	}
	if pos.Height != 188 {
		t.Fatalf(`expected height of 188 but got %v`, pos.Height)
	}
}

func TestSetPosition(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(tool)
	node.SetPosition(2, 4)
	pos, err := node.ReadPosition()
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	if pos.X != 2 || pos.Y != 4 {
		t.Fatalf(`expected position of 2,4 but got %v,%v`, pos.X, pos.Y)
	}
}

func TestSetPositionOfInvalidRyxNode(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(`<Node></Node>`)
	node.SetPosition(2, 4)
	if encoded, _ := xml.Marshal(node); string(encoded) != `<Node></Node>` {
		t.Fatalf(`expected no change to <Node></Node> but got %v`, string(encoded))
	}
}

func TestReadMacro(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(macro)
	macro := node.ReadMacro()
	if macro.StoredPath != `_Utilities\ChooseText.yxmc` {
		t.Fatalf(`expected stored path '_Utilities\ChooseText.yxmc' but got '%v'`, macro)
	}
	if macro.FoundPath != `` {
		t.Fatalf(`expected no found file but got '%v'`, macro.FoundPath)
	}
}

func TestSetMacro(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(macro)
	node.SetMacro(`hello world.yxmc`)
	if macro := node.ReadMacro(); macro.StoredPath != `hello world.yxmc` {
		t.Fatalf(`expected macro 'hello world.yxmc' but got '%v'`, macro)
	}
}

func TestReadChildNodes(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(container)
	children := node.ReadChildren()
	if len(children) != 3 {
		t.Fatalf(`expected 3 children but got %v`, len(children))
	}
}

func TestChangeChildNodes(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(container)
	children := node.ReadChildren()
	children[2].SetPosition(1, 2)

	rawBytes, _ := xml.Marshal(node)
	newNode := &ryxnode.RyxNode{}
	err := xml.Unmarshal(rawBytes, newNode)
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	pos, _ := newNode.ReadChildren()[2].ReadPosition()
	if pos.X != 1 || pos.Y != 2 {
		t.Fatalf(`marshalling and unmarshalling did not correctly persist change: expected position 1,2 but got %v,%v`, pos.X, pos.Y)
	}
}

func TestNewMacroXmlNode(t *testing.T) {
	macroXml := ryxnode.NewMacroXml(1, `macro.yxmc`, 5, 10)
	macro := ryxnode.New(macroXml)
	id, idErr := macro.ReadId()
	if idErr != nil {
		t.Fatalf(`expected no error retrieving ID but got %v`, idErr.Error())
	}
	if id != 1 {
		t.Fatalf(`expected id of 1 but got %v`, id)
	}
	if path := macro.ReadMacro(); path.StoredPath != `macro.yxmc` {
		t.Fatalf(`expected path 'macro.yxmc' but got '%v'`, path)
	}
	position, posErr := macro.ReadPosition()
	if posErr != nil {
		t.Fatalf(`expected no error reading position but got %v`, posErr.Error())
	}
	if position.X != 5 || position.Y != 10 {
		t.Fatalf(`expected position 5,10 but got %v,%v`, position.X, position.Y)
	}
}

func TestRemoveChildren(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(container)
	node.RemoveChildren(21, 23)
	children := node.ReadChildren()
	if len(children) != 1 {
		t.Fatalf(`expected 1 children but got %v`, len(children))
	}
}

func TestReadMacroFoundFile(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(`<Node ToolID="1"><GuiSettings><Position x="258" y="846" /></GuiSettings><EngineSettings Macro="macros\Tag with Sets.yxmc" /></Node>`)
	root, _ := filepath.Abs(filepath.Join(`..`, `testdocs`))
	expected, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `macros`, `Tag with Sets.yxmc`))
	macroPath := node.ReadMacro(root)
	if macroPath.FoundPath != expected {
		t.Fatalf(`expected found path '%v' but got '%v'`, expected, macroPath.FoundPath)
	}
	if macroPath.RelativeTo != root {
		t.Fatalf(`expected relative to of '%v' but got '%v'`, root, macroPath.RelativeTo)
	}
}

func TestMakeMacroAbsolute(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(`<Node ToolID="1"><GuiSettings><Position x="258" y="846" /></GuiSettings><EngineSettings Macro="macros\Tag with Sets.yxmc" /></Node>`)
	root, _ := filepath.Abs(filepath.Join(`..`, `testdocs`))
	expected, _ := filepath.Abs(filepath.Join(`..`, `testdocs`, `macros`, `Tag with Sets.yxmc`))
	expected = strings.Replace(expected, string(os.PathSeparator), `\`, -1)
	err := node.MakeMacroAbsolute(root)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if macroPath := node.ReadMacro(); macroPath.StoredPath != expected {
		t.Fatalf(`expected stored path '%v' but got '%v'`, expected, macroPath.StoredPath)
	}
}

func TestMakeMacroRelative(t *testing.T) {
	root, _ := filepath.Abs(filepath.Join(`..`, `testdocs`))
	nodeContent := `<Node ToolID="1"><GuiSettings><Position x="258" y="846" /></GuiSettings><EngineSettings Macro="` +
		strings.Replace(root, string(os.PathSeparator), `\`, -1) +
		`\macros\Tag with Sets.yxmc" /></Node>`
	node, _ := ryxnode.GenerateNodeFromXml(nodeContent)
	expected := `macros\Tag with Sets.yxmc`
	err := node.MakeMacroRelative(root, root)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if macroPath := node.ReadMacro(); macroPath.StoredPath != expected {
		t.Fatalf(`expected stored path '%v' but got '%v'`, expected, macroPath.StoredPath)
	}
}

func TestCategoryStringValue(t *testing.T) {
	categories := []ryxnode.Category{0, 1, 2, 3, 4}
	for _, category := range categories {
		if category.String() == `` {
			t.Fatalf(`expected a non-empty name`)
		}
		t.Logf(category.String())
	}
}

func TestReadPlugin(t *testing.T) {
	node, _ := ryxnode.GenerateNodeFromXml(tool)
	plugin := node.ReadPlugin()
	if plugin != `AlteryxBasePluginsGui.DbFileInput.DbFileInput` {
		t.Fatalf(`expected plugin 'AlteryxBasePluginsGui.DbFileInput.DbFileInput' but got '%v'`, plugin)
	}
}
