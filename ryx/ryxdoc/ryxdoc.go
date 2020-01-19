package ryxdoc

import (
	"encoding/xml"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
	"github.com/tlarsen7572/Golang-Public/txml"
	"io/ioutil"
)

type RyxDoc struct {
	data        *txml.Node
	nodes       []*ryxnode.RyxNode
	connections []*RyxConn
	nextId      int
}

type RyxConn struct {
	Name       string
	FromId     int
	ToId       int
	FromAnchor string
	ToAnchor   string
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

func (ryxDoc *RyxDoc) ReadNodes() []*ryxnode.RyxNode {
	return ryxDoc.nodes
}

func (ryxDoc *RyxDoc) ReadMappedNodes() map[int]*ryxnode.RyxNode {
	nodes := map[int]*ryxnode.RyxNode{}
	for _, node := range ryxDoc.nodes {
		addNodeToMap(node, nodes)
		for _, child := range node.ReadChildren() {
			addNodeToMap(child, nodes)
		}
	}
	return nodes
}

func (ryxDoc *RyxDoc) ReadAllConnections() []*RyxConn {
	return ryxDoc.connections
}

func (ryxDoc *RyxDoc) RemoveNodes(nodeIds ...int) {
	currentIndex := 0
	for _, node := range ryxDoc.nodes {
		if !node.MatchesIds(nodeIds...) {
			ryxDoc.nodes[currentIndex] = node
			node.RemoveChildren(nodeIds...)
			currentIndex += 1
		}
	}
	ryxDoc.nodes = ryxDoc.nodes[0:currentIndex]
}

func (ryxDoc *RyxDoc) RemoveConnectionsBetween(toolIds ...int) {
	var toDelete []*RyxConn
	for _, connection := range ryxDoc.connections {
		matchesFrom := intsContain(toolIds, connection.FromId)
		matchesTo := intsContain(toolIds, connection.ToId)
		if matchesFrom && matchesTo {
			toDelete = append(toDelete, connection)
		}
	}
	var keep []*RyxConn
	for _, conn := range ryxDoc.connections {
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
	ryxDoc.connections = keep
}

func (ryxDoc *RyxDoc) AddMacroAt(path string, x int, y int) *ryxnode.RyxNode {
	id := ryxDoc.grabNextIdAndIncrement()
	macro := ryxnode.NewMacroXml(id, path, x, y)
	node := ryxnode.New(macro)
	ryxDoc.nodes = append(ryxDoc.nodes, node)
	return node
}

func (ryxDoc *RyxDoc) AddConnection(connection *RyxConn) {
	ryxDoc.connections = append(ryxDoc.connections, connection)
}

func (ryxDoc *RyxDoc) RenameMacroNodes(macroAbsPath string, newPath string, macroPaths ...string) int {
	renamedNodes := 0
	for _, node := range ryxDoc.nodes {
		macro := node.ReadMacro(macroPaths...)
		if macro.FoundPath == macroAbsPath {
			node.SetMacro(newPath)
			renamedNodes++
		}
	}
	return renamedNodes
}

func (ryxDoc *RyxDoc) MakeAllMacrosAbsolute(macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.nodes {
		err := node.MakeMacroAbsolute(macroPaths...)
		if err == nil {
			changed++
		}
	}
	return changed
}

func (ryxDoc *RyxDoc) MakeMacroAbsolute(macroAbsPath string, macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.nodes {
		macro := node.ReadMacro(macroPaths...)
		if macro.FoundPath == macroAbsPath {
			err := node.MakeMacroAbsolute(macroPaths...)
			if err == nil {
				changed++
			}
		}
	}
	return changed
}

func (ryxDoc *RyxDoc) MakeAllMacrosRelative(relativeTo string, macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.nodes {
		err := node.MakeMacroRelative(relativeTo, macroPaths...)
		if err == nil {
			changed++
		}
	}
	return changed
}

func (ryxDoc *RyxDoc) MakeMacroRelative(macroAbsPath string, relativeTo string, macroPaths ...string) int {
	changed := 0
	for _, node := range ryxDoc.nodes {
		macro := node.ReadMacro(macroPaths...)
		if macro.FoundPath == macroAbsPath {
			err := node.MakeMacroRelative(relativeTo, macroPaths...)
			if err == nil {
				changed++
			}
		}
	}
	return changed
}

func (ryxDoc *RyxDoc) Save(path string) error {
	data, err := xml.Marshal(ryxDoc)
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
