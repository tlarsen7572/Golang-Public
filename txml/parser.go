package txml

import (
	"encoding/xml"
	"sort"
)

// Parse processes the provided XML string and returns a generic data structure representing the root element and its children.
// If there are multiple root elements, only the first element is returned.
func Parse(xmlString string) (*Node, error) {
	var xmlParsed = &Node{}
	err := xml.Unmarshal([]byte(xmlString), xmlParsed)
	return xmlParsed, err
}

// The Node struct contains basic data elements necessary to represent XML data structures
type Node struct {
	Name       string            `xml:"-"`
	Attributes map[string]string `xml:"-"`
	InnerText  string            `xml:",chardata"`
	Nodes      []*Node           `xml:",any"`
}

func NilNode() *Node {
	return &Node{Name: ``, Attributes: map[string]string{}, InnerText: ``, Nodes: []*Node{}}
}

func (node *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	node.Attributes = make(map[string]string)
	for _, attr := range start.Attr {
		node.Attributes[attr.Name.Local] = attr.Value
	}
	node.Name = start.Name.Local
	type nodePointer Node
	err := d.DecodeElement((*nodePointer)(node), &start)
	if err == nil && len(node.Nodes) > 0 {
		node.InnerText = ``
	}
	return err
}

// FirstChild returns the first instance of the specified element name, or a nilNode if none is found
func (node *Node) First(findName string) *Node {
	for _, child := range node.Nodes {
		if child.Name == findName {
			return child
		}
	}
	return NilNode()
}

// AllNodes returns all elements with the specified name
func (node *Node) AllNodes(findName string) []*Node {
	var matchingNodes []*Node
	for _, child := range node.Nodes {
		if child.Name == findName {
			matchingNodes = append(matchingNodes, child)
		}
	}
	return matchingNodes
}

func (node *Node) ReplaceFirst(findName string, replaceWith *Node) {
	for index := range node.Nodes {
		if node.Nodes[index].Name == findName {
			node.Nodes[index] = replaceWith
			return
		}
	}
}

func (node *Node) ReplaceWith(findName string, callback func(*Node) *Node) {
	for index := range node.Nodes {
		currentNode := node.Nodes[index]
		if currentNode.Name == findName {
			node.Nodes[index] = callback(currentNode)
		}
	}
}

func (node *Node) RemoveFirst(findName string) {
	var index int
	for index = range node.Nodes {
		if node.Nodes[index].Name == findName {
			break
		}
	}
	node.Nodes = append(node.Nodes[:index], node.Nodes[index+1:]...)
}

func (node *Node) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var sortedKeys = sortKeys(node.Attributes)
	for _, key := range sortedKeys {
		var value = node.Attributes[key]
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: key}, Value: value})
	}
	start.Name = xml.Name{Local: node.Name}
	type nodePointer Node
	return e.EncodeElement((*nodePointer)(node), start)
}

// ToXml converts the node back to its XML representation
func (node *Node) ToXml(indent string) (string, error) {
	value, err := xml.MarshalIndent(node, ``, indent)
	return string(value), err
}

func sortKeys(dict map[string]string) []string {
	var keys []string
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
