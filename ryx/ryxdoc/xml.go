package ryxdoc

import (
	"encoding/xml"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
	"github.com/tlarsen7572/Golang-Public/txml"
	"strconv"
)

func (ryxDoc *RyxDoc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tempData := ryxDoc.createTempData()
	ryxDoc.marshalNodesInto(tempData)
	ryxDoc.marshalConnectionsInto(tempData)
	ryxDoc.marshalRemainderInto(tempData)
	e.Indent(``, `  `)
	return e.EncodeElement(tempData, start)
}

func (ryxDoc *RyxDoc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	data := &txml.Node{}
	err := d.DecodeElement(data, &start)
	if err != nil {
		return err
	}
	ryxDoc.data = data
	ryxDoc.collectNodesAndRemoveXml()
	err = ryxDoc.collectConnectionsAndRemoveXml()
	return err
}

func (ryxDoc *RyxDoc) createTempData() *txml.Node {
	newNodes := make([]*txml.Node, len(ryxDoc.nodes)+2)
	newNodes[0] = &txml.Node{
		Name:  `Nodes`,
		Nodes: make([]*txml.Node, len(ryxDoc.nodes)),
	}
	newNodes[1] = &txml.Node{
		Name:  `Connections`,
		Nodes: make([]*txml.Node, len(ryxDoc.connections)),
	}
	return &txml.Node{
		Name:       ryxDoc.data.Name,
		Attributes: ryxDoc.data.Attributes,
		InnerText:  ryxDoc.data.InnerText,
		Nodes:      newNodes,
	}
}

func (ryxDoc *RyxDoc) marshalNodesInto(tempData *txml.Node) {
	for index, node := range ryxDoc.nodes {
		tempData.Nodes[0].Nodes[index] = node.ReadData()
	}
}

func (ryxDoc *RyxDoc) marshalConnectionsInto(tempData *txml.Node) {
	for index, conn := range ryxDoc.connections {
		wireless := `False`
		if conn.Wireless {
			wireless = `True`
		}
		tempData.Nodes[1].Nodes[index] = &txml.Node{
			Name:       `Connection`,
			Attributes: map[string]string{`name`: conn.Name, `Wireless`: wireless},
			Nodes: []*txml.Node{
				{
					Name:       "Origin",
					Attributes: map[string]string{`ToolID`: strconv.Itoa(conn.FromId), `Connection`: conn.FromAnchor},
				},
				{
					Name:       "Destination",
					Attributes: map[string]string{`ToolID`: strconv.Itoa(conn.ToId), `Connection`: conn.ToAnchor},
				},
			},
		}
	}
}

func (ryxDoc *RyxDoc) marshalRemainderInto(tempData *txml.Node) {
	for index, data := range ryxDoc.data.Nodes {
		tempData.Nodes[index+2] = data
	}
}

func (ryxDoc *RyxDoc) collectNodesAndRemoveXml() {
	nodes := ryxDoc.data.First(`Nodes`).Nodes
	nodeCount := len(nodes)
	ryxDoc.nodes = make([]*ryxnode.RyxNode, nodeCount)
	for index, node := range nodes {
		ryxDoc.nodes[index] = ryxnode.New(node)
	}
	ryxDoc.data.RemoveFirst(`Nodes`)
}

func (ryxDoc *RyxDoc) collectConnectionsAndRemoveXml() error {
	connections := ryxDoc.data.First(`Connections`).Nodes
	connCount := len(connections)
	ryxDoc.connections = make([]*RyxConn, connCount)
	for index, conn := range connections {
		name := conn.Attributes[`name`]
		wireless := false
		wirelessAttr := conn.Attributes[`Wireless`]
		if wirelessAttr == `True` {
			wireless = true
		}
		origin := conn.First(`Origin`)
		destination := conn.First(`Destination`)
		fromId, err := strconv.Atoi(origin.Attributes[`ToolID`])
		if err != nil {
			return err
		}
		toId, err := strconv.Atoi(destination.Attributes[`ToolID`])
		if err != nil {
			return err
		}
		fromAnchor := origin.Attributes[`Connection`]
		toAnchor := destination.Attributes[`Connection`]
		ryxDoc.connections[index] = &RyxConn{Name: name, FromId: fromId, ToId: toId, FromAnchor: fromAnchor, ToAnchor: toAnchor, Wireless: wireless}
	}
	ryxDoc.data.RemoveFirst(`Connections`)
	return nil
}
