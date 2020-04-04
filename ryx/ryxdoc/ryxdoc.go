package ryxdoc

import (
	"encoding/xml"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
	"github.com/tlarsen7572/Golang-Public/txml"
	"io/ioutil"
)

type RyxDoc struct {
	XMLName xml.Name           `xml:"AlteryxDocument"`
	YxmdVer string             `xml:"yxmdVer,attr"`
	Nodes   []*ryxnode.RyxNode `xml:"Nodes>Node"` // Be careful using Nodes directly.  This field does not
	// give you all nodes in a document, only the root nodes.
	// Nodes can themselves contain nodes.  To iterate through all
	// nodes in a document, use ReadMappedNodes()
	Connections []*RyxConn `xml:"Connections>Connection"`
	Properties  *txml.Node `xml:"Properties"`
	nextId      int
}

type RyxConn struct {
	Name       string
	FromId     int
	ToId       int
	FromAnchor string
	ToAnchor   string
	Wireless   bool
}

func ReadFile(path string) (*RyxDoc, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ReadBytes(content)
}

func ReadBytes(content []byte) (*RyxDoc, error) {
	workflow := &RyxDoc{}
	err := xml.Unmarshal(content, workflow)
	if err != nil {
		return nil, err
	}
	maxId := 0
	for id := range workflow.ReadMappedNodes() {
		if id > maxId {
			maxId = id
		}
	}
	workflow.nextId = maxId + 1
	return workflow, nil
}

func (ryxDoc *RyxDoc) ReadMappedNodes() map[int]*ryxnode.RyxNode {
	nodes := map[int]*ryxnode.RyxNode{}
	for _, node := range ryxDoc.Nodes { // It is ok to use RyxDoc.Nodes here
		addNodeToMap(node, nodes)
		for _, child := range node.ReadChildren() {
			addNodeToMap(child, nodes)
		}
	}
	return nodes
}

func (ryxDoc *RyxDoc) RemoveNodes(nodeIds ...int) {
	currentIndex := 0
	for _, node := range ryxDoc.Nodes { // It is ok to use RyxDoc.Nodes here
		if !node.MatchesIds(nodeIds...) {
			ryxDoc.Nodes[currentIndex] = node // It is ok to use RyxDoc.Nodes here
			node.RemoveChildren(nodeIds...)
			currentIndex += 1
		}
	}
	ryxDoc.Nodes = ryxDoc.Nodes[0:currentIndex] // It is ok to use RyxDoc.Nodes here
}

func (ryxDoc *RyxDoc) RemoveConnectionsBetween(toolIds ...int) {
	var toDelete []*RyxConn
	for _, connection := range ryxDoc.Connections {
		matchesFrom := intsContain(toolIds, connection.FromId)
		matchesTo := intsContain(toolIds, connection.ToId)
		if matchesFrom && matchesTo {
			toDelete = append(toDelete, connection)
		}
	}
	var keep []*RyxConn
	for _, conn := range ryxDoc.Connections {
		matches := false
		for _, toRemove := range toDelete {
			if conn.FromId == toRemove.FromId &&
				conn.ToId == toRemove.ToId &&
				conn.FromAnchor == toRemove.FromAnchor &&
				conn.ToAnchor == toRemove.ToAnchor &&
				conn.Name == toRemove.Name {
				matches = true
				break
			}
		}
		if !matches {
			keep = append(keep, conn)
		}
	}
	ryxDoc.Connections = keep
}

func (ryxDoc *RyxDoc) AddMacroAt(path string, x float64, y float64) *ryxnode.RyxNode {
	id := ryxDoc.grabNextIdAndIncrement()
	macro := ryxnode.NewMacro(id, path, x, y)
	ryxDoc.Nodes = append(ryxDoc.Nodes, macro) // It is ok to use RyxDoc.Nodes here
	return macro
}

func (ryxDoc *RyxDoc) AddConnection(connection *RyxConn) {
	ryxDoc.Connections = append(ryxDoc.Connections, connection)
}

func (ryxDoc *RyxDoc) RenameMacroNodes(macroAbsPaths []string, newPaths []string, macroPaths ...string) int {
	changer := func(node *ryxnode.RyxNode, macroAbsPathIndex int) error {
		node.SetMacro(newPaths[macroAbsPathIndex])
		return nil
	}
	return ryxDoc._changeMacrosIfMatches(macroAbsPaths, changer, macroPaths...)
}

func (ryxDoc *RyxDoc) MakeAllMacrosAbsolute(macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.ReadMappedNodes() {
		err := node.MakeMacroAbsolute(macroPaths...)
		if err == nil {
			changed++
		}
	}
	return changed
}

func (ryxDoc *RyxDoc) MakeMacrosAbsolute(macroAbsPaths []string, macroPaths ...string) int {
	changer := func(node *ryxnode.RyxNode, macroAbsPathIndex int) error {
		return node.MakeMacroAbsolute(macroPaths...)
	}
	return ryxDoc._changeMacrosIfMatches(macroAbsPaths, changer, macroPaths...)
}

func (ryxDoc *RyxDoc) MakeAllMacrosRelative(relativeTo string, macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.ReadMappedNodes() {
		err := node.MakeMacroRelative(relativeTo, macroPaths...)
		if err == nil {
			changed++
		}
	}
	return changed
}

func (ryxDoc *RyxDoc) MakeMacrosRelative(macroAbsPaths []string, relativeTo string, macroPaths ...string) int {
	changer := func(node *ryxnode.RyxNode, macroAbsPathIndex int) error {
		return node.MakeMacroRelative(relativeTo, macroPaths...)
	}
	return ryxDoc._changeMacrosIfMatches(macroAbsPaths, changer, macroPaths...)
}

func (ryxDoc *RyxDoc) Save(path string) error {
	data, err := xml.MarshalIndent(ryxDoc, ``, `  `)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func addNodeToMap(node *ryxnode.RyxNode, nodes map[int]*ryxnode.RyxNode) {
	id, err := node.ReadId()
	if err != nil {
		return
	}
	nodes[id] = node
}

func (ryxDoc *RyxDoc) grabNextIdAndIncrement() int {
	id := 0
	id, ryxDoc.nextId = ryxDoc.nextId, ryxDoc.nextId+1
	return id
}

func (ryxDoc *RyxDoc) _changeMacrosIfMatches(macroAbsPaths []string, changer func(node *ryxnode.RyxNode, macroAbsPathIndex int) error, macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.ReadMappedNodes() {
		macro := node.ReadMacro(macroPaths...)
		for macroAbsPathIndex, macroAbsPath := range macroAbsPaths {
			if macro.FoundPath == macroAbsPath {
				err := changer(node, macroAbsPathIndex)
				if err == nil {
					changed++
				}
				break
			}
		}
	}
	return changed
}
