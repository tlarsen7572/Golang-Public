package ryxdoc

import (
	"errors"
	"fmt"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
	"github.com/tlarsen7572/Golang-Public/txml"
	"path/filepath"
	"strconv"
)

var gridStartPos = 54
var gridSize = 12
var horizontalGap = gridSize * 8
var verticalGap = gridSize * 7

func (ryxDoc *RyxDoc) ExtractMacro(macroAbsPath string, relativeTo string, toolIds ...int) error {
	macroPath := macroAbsPath
	var err error
	if relativeTo != `` {
		macroPath, err = filepath.Rel(relativeTo, macroAbsPath)
		if err != nil {
			return err
		}
	}
	if HasHole(ryxDoc, toolIds...) {
		return errors.New(`there is a hole in the selected tools - ExtractMacro cannot continue`)
	}

	left, top, right, _ := ryxDoc.getBoundingBox(toolIds...)
	newMacroTool := ryxDoc.AddMacroAt(macroPath, left, top)
	newMacro := generateDocFrom(ryxDoc, toolIds...)
	adjustX, _ := normalizeToolPositions(newMacro, left, top)
	generateMacroConnections(newMacro, ryxDoc, newMacroTool, right+adjustX+horizontalGap, toolIds...)
	ryxDoc.RemoveConnectionsBetween(toolIds...)
	ryxDoc.RemoveNodes(toolIds...)

	err = newMacro.Save(macroAbsPath)
	if err != nil {
		return err
	}
	return nil
}

func HasHole(doc *RyxDoc, toolIds ...int) bool {
	var unselected []int
	for toolId := range doc.ReadMappedNodes() {
		if !intsContain(toolIds, toolId) {
			unselected = append(unselected, toolId)
		}
	}

	outputConns, inputConns := readIoConns(doc)
	for _, startingAt := range unselected {

		selectedUpstream := upstreamSelected(inputConns, startingAt, toolIds...)
		selectedDownstream := downstreamSelected(outputConns, startingAt, toolIds...)
		if selectedUpstream && selectedDownstream {
			return true
		}
	}
	return false
}

func upstreamSelected(inputConns map[int][]*RyxConn, startingAt int, toolIds ...int) bool {
	getter := func(conn *RyxConn) int { return conn.FromId }
	return directionSelected(inputConns, startingAt, getter, toolIds...)
}

func downstreamSelected(outputConns map[int][]*RyxConn, startingAt int, toolIds ...int) bool {
	getter := func(conn *RyxConn) int { return conn.ToId }
	return directionSelected(outputConns, startingAt, getter, toolIds...)
}

func directionSelected(connections map[int][]*RyxConn, startingAt int, getId func(*RyxConn) int, toolIds ...int) bool {
	for _, connection := range connections[startingAt] {
		id := getId(connection)
		if intsContain(toolIds, id) {
			return true
		}
		distantlySelected := directionSelected(connections, id, getId, toolIds...)
		if distantlySelected {
			return true
		}
	}
	return false
}

func normalizeToolPositions(doc *RyxDoc, minLeft int, minTop int) (int, int) {
	startX := gridStartPos + horizontalGap
	startY := gridStartPos
	adjustX := startX - minLeft
	adjustY := startY - minTop
	for _, node := range doc.ReadMappedNodes() {
		position, err := node.ReadPosition()
		if err == nil {
			node.SetPosition(position.X+adjustX, position.Y+adjustY)
		}
	}
	return adjustX, adjustY
}

func generateMacroConnections(newMacro *RyxDoc, origDoc *RyxDoc, newMacroTool *ryxnode.RyxNode, outputX int, toolIds ...int) {
	newMacroToolId, err := newMacroTool.ReadId()
	if err != nil {
		return
	}

	questionTabId := newMacro.grabNextIdAndIncrement()
	questionTab := ryxnode.New(newQuestionTabXml(questionTabId))
	newMacro.nodes = append(newMacro.nodes, questionTab)
	tab := addTabQuestion(newMacro, questionTabId)

	inputCount := 0
	outputCount := 0
	for _, connection := range origDoc.connections {
		matchesFrom := intsContain(toolIds, connection.FromId)
		matchesTo := intsContain(toolIds, connection.ToId)
		if matchesFrom && !matchesTo {
			y := gridStartPos + (verticalGap * outputCount)
			outputCount += 1
			outputId := newMacro.grabNextIdAndIncrement()
			output := ryxnode.New(newMacroOutputXml(outputId, outputX, y))
			addQuestionToTab(tab, `MacroOutput`, fmt.Sprintf(`Macro Output (%v)`, outputId), outputId)
			newMacro.nodes = append(newMacro.nodes, output)
			newMacro.AddConnection(&RyxConn{
				Name:       connection.Name,
				FromId:     connection.FromId,
				FromAnchor: connection.FromAnchor,
				ToId:       outputId,
				ToAnchor:   `Input`,
			})
			connection.FromId = newMacroToolId
			connection.FromAnchor = `Output` + strconv.Itoa(outputId)
		}
		if matchesTo && !matchesFrom {
			y := gridStartPos + (verticalGap * inputCount)
			inputCount += 1
			inputId := newMacro.grabNextIdAndIncrement()
			input := ryxnode.New(newMacroInputXml(inputId, gridStartPos, y))
			addQuestionToTab(tab, `MacroInput`, fmt.Sprintf(`Macro Input (%v)`, inputId), inputId)
			newMacro.nodes = append(newMacro.nodes, input)
			newMacro.AddConnection(&RyxConn{
				Name:       connection.Name,
				FromId:     inputId,
				FromAnchor: `Output`,
				ToId:       connection.ToId,
				ToAnchor:   connection.ToAnchor,
			})
			connection.ToId = newMacroToolId
			connection.ToAnchor = `Input` + strconv.Itoa(inputId)
		}
	}
}

func readIoConns(doc *RyxDoc) (outputConns map[int][]*RyxConn, inputConns map[int][]*RyxConn) {
	outputConns = map[int][]*RyxConn{}
	inputConns = map[int][]*RyxConn{}
	for _, conn := range doc.connections {
		outputConns[conn.FromId] = append(outputConns[conn.FromId], conn)
		inputConns[conn.ToId] = append(inputConns[conn.ToId], conn)
	}
	return outputConns, inputConns
}

func generateDocFrom(from *RyxDoc, toolIds ...int) *RyxDoc {
	content := []byte(docXml)
	doc, _ := ReadBytes(content)
	copyNodes(from, doc, toolIds...)
	copyConnections(from, doc, toolIds...)
	doc.nextId = readLargestInt(toolIds) + 1
	return doc
}

func copyNodes(from *RyxDoc, to *RyxDoc, toolIds ...int) {
	for _, node := range from.nodes {
		for _, id := range toolIds {
			if checkId, _ := node.ReadId(); checkId == id {
				to.nodes = append(to.nodes, node)
				break
			}
		}
	}
}

func copyConnections(from *RyxDoc, to *RyxDoc, toolIds ...int) {
	for _, connection := range from.connections {
		matchesFrom := intsContain(toolIds, connection.FromId)
		matchesTo := intsContain(toolIds, connection.ToId)
		if matchesFrom && matchesTo {
			to.AddConnection(connection)
		}
	}
}

func (ryxDoc *RyxDoc) getBoundingBox(toolIds ...int) (left int, top int, right int, bottom int) {
	left = maxInt()
	top = maxInt()
	right = 0
	bottom = 0
	for _, node := range ryxDoc.nodes {
		for _, id := range toolIds {
			if checkId, _ := node.ReadId(); checkId == id {
				if position, err := node.ReadPosition(); err == nil {
					if position.X < left {
						left = position.X
					}
					if position.Y < top {
						top = position.Y
					}
					if position.X > right {
						right = position.X
					}
					if position.Y > bottom {
						bottom = position.Y
					}
				}
				break
			}
		}
	}
	return left, top, right, bottom
}

func intsContain(values []int, check int) bool {
	for _, value := range values {
		if value == check {
			return true
		}
	}
	return false
}

func maxInt() int {
	return int(^uint(0) >> 1)
}

func readLargestInt(values []int) int {
	max := 0
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func addTabQuestion(doc *RyxDoc, tabId int) *txml.Node {
	node := doc.data.First(`Properties`).First(`RuntimeProperties`).First(`Questions`)
	name := fmt.Sprintf(`Tab (%v)`, tabId)
	question := generateQuestion(`Tab`, `Questions`, name, tabId)
	node.Nodes = append(node.Nodes, question)
	return question
}

func addQuestionToTab(tab *txml.Node, qType string, name string, toolId int) {
	questions := tab.First(`Questions`)
	question := generateQuestion(qType, name, name, toolId)
	questions.Nodes = append(questions.Nodes, question)
}

func generateQuestion(qType string, description string, name string, toolId int) *txml.Node {
	return &txml.Node{
		Name: `Question`,
		Nodes: []*txml.Node{
			{
				Name:      `Type`,
				InnerText: qType,
			},
			{
				Name:      `Description`,
				InnerText: description,
			},
			{
				Name:      `Name`,
				InnerText: name,
			},
			{
				Name:       `ToolId`,
				Attributes: map[string]string{`value`: strconv.Itoa(toolId)},
			},
			{
				Name: `Questions`,
			},
		},
	}
}

func newMacroInputXml(id int, x int, y int) *txml.Node {
	return &txml.Node{
		Name: `Node`,
		Attributes: map[string]string{
			`ToolID`: strconv.Itoa(id),
		},
		Nodes: []*txml.Node{
			{
				Name:       `GuiSettings`,
				Attributes: map[string]string{`Plugin`: `AlteryxBasePluginsGui.MacroInput.MacroInput`},
				Nodes: []*txml.Node{
					{
						Name:       `Position`,
						Attributes: map[string]string{`x`: strconv.Itoa(x), `y`: strconv.Itoa(y)},
					},
				},
			},
			{
				Name: `Properties`,
				Nodes: []*txml.Node{
					{
						Name: `Configuration`,
						Nodes: []*txml.Node{
							{
								Name:       `UseFileInput`,
								Attributes: map[string]string{`value`: `False`},
							},
							{
								Name:      `Name`,
								InnerText: `Input` + strconv.Itoa(id),
							},
							{
								Name: `Abbrev`,
							},
							{
								Name:       `ShowFieldMap`,
								Attributes: map[string]string{`value`: `False`},
							},
							{
								Name:       `Optional`,
								Attributes: map[string]string{`value`: `False`},
							},
							{
								Name: `TextInput`,
								Nodes: []*txml.Node{
									{
										Name: `Configuration`,
										Nodes: []*txml.Node{
											{Name: `NumRows`, Attributes: map[string]string{`value`: `0`}},
											{Name: `Fields`},
											{Name: `Data`},
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: `EngineSettings`,
				Attributes: map[string]string{
					`EngineDll`:           `AlteryxBasePluginsEngine.dll`,
					`EngineDllEntryPoint`: `AlteryxMacroInput`,
				},
			},
		},
	}
}

func newMacroOutputXml(id int, x int, y int) *txml.Node {
	return &txml.Node{
		Name: `Node`,
		Attributes: map[string]string{
			`ToolID`: strconv.Itoa(id),
		},
		Nodes: []*txml.Node{
			{
				Name:       `GuiSettings`,
				Attributes: map[string]string{`Plugin`: `AlteryxBasePluginsGui.MacroOutput.MacroOutput`},
				Nodes: []*txml.Node{
					{
						Name:       `Position`,
						Attributes: map[string]string{`x`: strconv.Itoa(x), `y`: strconv.Itoa(y)},
					},
				},
			},
			{
				Name: `Properties`,
				Nodes: []*txml.Node{
					{
						Name: `Configuration`,
						Nodes: []*txml.Node{
							{
								Name:      `Name`,
								InnerText: `Output` + strconv.Itoa(id),
							},
						},
					},
				},
			},
			{
				Name: `EngineSettings`,
				Attributes: map[string]string{
					`EngineDll`:           `AlteryxBasePluginsEngine.dll`,
					`EngineDllEntryPoint`: `AlteryxMacroOutput`,
				},
			},
		},
	}
}

func newQuestionTabXml(id int) *txml.Node {
	return &txml.Node{
		Name: `Node`,
		Attributes: map[string]string{
			`ToolID`: strconv.Itoa(id),
		},
		Nodes: []*txml.Node{
			{
				Name:       `GuiSettings`,
				Attributes: map[string]string{`Plugin`: `AlteryxGuiToolkit.Questions.Tab.Tab`},
				Nodes: []*txml.Node{
					{
						Name:       `Position`,
						Attributes: map[string]string{`x`: `0`, `y`: `0`},
					},
				},
			},
		},
	}
}

var docXml = `<AlteryxDocument yxmdVer="2019.3">
  <Nodes>
  </Nodes>
  <Connections>
  </Connections>
  <Properties>
    <Memory default="True" />
    <GlobalRecordLimit value="0" />
    <TempFiles default="True" />
    <Annotation on="True" includeToolName="False" />
    <ConvErrorLimit value="10" />
    <ConvErrorLimit_Stop value="False" />
    <CancelOnError value="False" />
    <DisableBrowse value="False" />
    <EnablePerformanceProfiling value="False" />
    <DisableAllOutput value="False" />
    <ShowAllMacroMessages value="False" />
    <ShowConnectionStatusIsOn value="True" />
    <ShowConnectionStatusOnlyWhenRunning value="True" />
    <ZoomLevel value="0" />
    <LayoutType>Horizontal</LayoutType>
    <Constants>
    </Constants>
    <MetaInfo>
      <NameIsFileName value="True" />
      <Name></Name>
      <Description />
      <RootToolName />
      <ToolVersion />
      <ToolInDb value="False" />
      <CategoryName />
      <SearchTags />
      <Author />
      <Company />
      <Copyright />
      <DescriptionLink actual="" displayed="" />
      <Example>
        <Description />
        <File />
      </Example>
    </MetaInfo>
    <Events>
      <Enabled value="True" />
    </Events>
    <RuntimeProperties>
      <Actions />
      <Questions>
      </Questions>
      <ModuleType>Macro</ModuleType>
      <MacroCustomHelp value="False" />
      <MacroDynamicOutputFields value="False" />
      <MacroImageStd value="39" />
      <MacroInputs />
      <MacroOutputs />
      <Wiz_CustomHelp value="False" />
      <Wiz_CustomGraphic value="False" />
      <Wiz_ShowOutput value="True" />
      <Wiz_OpenOutputTools>
      </Wiz_OpenOutputTools>
      <Wiz_OutputMessage />
      <Wiz_NoOutputFilesMessage />
      <Wiz_ChainRunWizard />
    </RuntimeProperties>
  </Properties>
</AlteryxDocument>`
