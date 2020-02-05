package txml_test

import (
	"github.com/tlarsen7572/Golang-Public/txml"
	"testing"
)

const singleChildXml = `<Element id="1" type="Something"><Sub>My Text</Sub></Element>`
const multiChildXml = `<Root><Child id="1"></Child><Child id="2"></Child></Root>`
const multiComplexChildXml = `<Root><Child id="1"></Child><Child id="2"></Child><Something></Something></Root>`
const humanReadable = `<Root>
  <Node id="1"></Node>
  <Node id="2"></Node>
</Root>`

func TestParseSingleChildXml(t *testing.T) {
	parsed, err := txml.Parse(singleChildXml)

	if err != nil {
		t.Fatalf(`expected no error but got '%v'`, err.Error())
	}
	if parsed == nil {
		t.Fatalf(`expected non-nil parsed but got nil`)
	}
	if idAttr := parsed.Attributes[`id`]; idAttr != `1` {
		t.Fatalf(`expected id of 1 but got %v`, idAttr)
	}
	if typeAttr := parsed.Attributes[`type`]; typeAttr != `Something` {
		t.Fatalf(`expected type of 'Something' but got '%v'`, typeAttr)
	}
	if count := len(parsed.Nodes); count != 1 {
		t.Fatalf(`expected 1 child but not %v`, count)
	}
	if innerText := parsed.Nodes[0].InnerText; innerText != `My Text` {
		t.Fatalf(`expected inner text 'My Text' but got '%v'`, innerText)
	}
}

func TestParseMultiChildXml(t *testing.T) {
	parsed, err := txml.Parse(multiChildXml)

	if err != nil {
		t.Fatalf(`expected no error but got '%v'`, err.Error())
	}
	if id := parsed.First(`Child`).Attributes[`id`]; id != `1` {
		t.Fatalf(`expected first child id of 1 but got %v`, id)
	}
	if id := parsed.AllNodes(`Child`)[1].Attributes[`id`]; id != `2` {
		t.Fatalf(`expected second child id of 2 but got '%v'`, id)
	}
}

func TestParseComplexMultipleChildrenXml(t *testing.T) {
	parsed, err := txml.Parse(multiComplexChildXml)

	if err != nil {
		t.Fatalf(`expected no error but got '%v'`, err.Error())
	}
	if count := len(parsed.Nodes); count != 3 {
		t.Fatalf(`expected count of 3 but got '%v'`, count)
	}
	if count := len(parsed.AllNodes(`Child`)); count != 2 {
		t.Fatalf(`expected Child count of 2 but got %v`, count)
	}
	if count := len(parsed.AllNodes(`Something`)); count != 1 {
		t.Fatalf(`expected Something count of 1 but got %v`, count)
	}
}

func TestSingleChildToXml(t *testing.T) {
	node := txml.Node{
		Name:       "Element",
		Attributes: map[string]string{"id": "1", "type": "Something"},
		Nodes: []*txml.Node{
			{Name: "Sub", InnerText: "My Text"},
		},
	}
	xml, err := node.ToXml(``)

	if err != nil {
		t.Fatalf(`expected no error but got '%v'`, err.Error())
	}
	if xml != singleChildXml {
		t.Fatalf(`expected xml '%v' but got '%v'`, singleChildXml, xml)
	}
}

func TestMultiChildToXml(t *testing.T) {
	node := txml.Node{
		Name: "Root",
		Nodes: []*txml.Node{
			{Name: "Child", Attributes: map[string]string{"id": "1"}},
			{Name: "Child", Attributes: map[string]string{"id": "2"}},
		},
	}
	xml, err := node.ToXml(``)

	if err != nil {
		t.Fatalf(`expected no error but got '%v'`, err.Error())
	}
	if xml != multiChildXml {
		t.Fatalf(`expected xml '%v' but got '%v'`, multiChildXml, xml)
	}
}

func TestComplexMultiChildrenToXml(t *testing.T) {
	node := txml.Node{
		Name: "Root",
		Nodes: []*txml.Node{
			{Name: "Child", Attributes: map[string]string{"id": "1"}},
			{Name: "Child", Attributes: map[string]string{"id": "2"}},
			{Name: "Something"},
		},
	}
	xml, err := node.ToXml(``)

	if err != nil {
		t.Fatalf(`expected no error but got '%v'`, err.Error())
	}
	if xml != multiComplexChildXml {
		t.Fatalf(`expected xml '%v' but got '%v'`, multiComplexChildXml, xml)
	}
}

func TestFirstChildThatDoesNotExist(t *testing.T) {
	parsed, _ := txml.Parse(singleChildXml)
	first := parsed.First("NotInData")
	if first == nil {
		t.Fatalf("calling first with invalid name should not be nil")
	}
	if first.Name != `` {
		t.Fatalf("the name of an invalid child should be empty")
	}
}

func TestAttributeThatDoesNotExist(t *testing.T) {
	parsed, _ := txml.Parse(singleChildXml)
	if parsed.Attributes[`InvalidAttribute`] != `` {
		t.Fatalf(`expected empty string but got %v`, parsed.Attributes[`InvalidAttribute`])
	}
}

func TestReplaceFirst(t *testing.T) {
	parsed, _ := txml.Parse(multiChildXml)
	newNode := &txml.Node{Name: `Child`, Attributes: map[string]string{`id`: `10`}}
	parsed.ReplaceFirst(`Child`, newNode)
	if id := parsed.First(`Child`).Attributes[`id`]; id != `10` {
		t.Fatalf(`expected first child id of 10 but got %v`, id)
	}
}

func TestReplaceWith(t *testing.T) {
	parsed, _ := txml.Parse(multiChildXml)
	callback := func(node *txml.Node) *txml.Node {
		return &txml.Node{Name: `NewChild`, Attributes: node.Attributes}
	}
	parsed.ReplaceWith(`Child`, callback)
	newChild := parsed.AllNodes(`NewChild`)
	if len(newChild) != 2 {
		t.Fatalf(`expected 2 NewChild elements but got %v`, len(newChild))
	}
	if id := newChild[0].Attributes[`id`]; id != `1` {
		t.Fatalf(`expected first NewChild id of '1' but got '%v'`, id)
	}
	if id := newChild[1].Attributes[`id`]; id != `2` {
		t.Fatalf(`expected first NewChild id of '2' but got '%v'`, id)
	}
	if count := len(parsed.AllNodes(`Child`)); count > 0 {
		t.Fatalf(`expected 0 Child elements but got %v`, count)
	}
}

func TestRemoveFirst(t *testing.T) {
	parsed, _ := txml.Parse(multiChildXml)
	parsed.RemoveFirst(`Child`)
	if len(parsed.Nodes) != 1 {
		t.Fatalf(`expected 1 child node but got %v`, len(parsed.Nodes))
	}
	if id := parsed.Nodes[0].Attributes[`id`]; id != `2` {
		t.Fatalf(`expected id 2 but got %v`, id)
	}
}

func TestRemoveAll(t *testing.T) {
	parsed, _ := txml.Parse(multiChildXml)
	parsed.RemoveAll(`Child`)
	if count := len(parsed.Nodes); count != 0 {
		t.Fatalf(`expected 0 nodes but got %v`, count)
	}
}

func TestHumanReadable(t *testing.T) {
	parsed, err := txml.Parse(humanReadable)
	if err != nil {
		t.Fatalf(`expected no error parsing but got: %v`, err.Error())
	}
	xml, err := parsed.ToXml(`  `)
	if err != nil {
		t.Fatalf(`expected no error converting to XML but got: %v`, err.Error())
	}
	if xml != humanReadable {
		t.Fatalf(`expected '%v' but got '%v'`, humanReadable, xml)
	}
}
