package ryxnode

import (
	"encoding/xml"
	"github.com/tlarsen7572/Golang-Public/txml"
)

type _RyxNodeWithChildrenXml struct {
	ToolId         string      `xml:"ToolID,attr"`
	GuiSettings    *txml.Node  `xml:"GuiSettings"`
	Properties     *Properties `xml:"Properties"`
	EngineSettings *txml.Node  `xml:"EngineSettings,omitempty"`
	ChildNodes     []*RyxNode  `xml:"ChildNodes>Node,omitempty"`
}

type _RyxNodeXml struct {
	ToolId         string      `xml:"ToolID,attr"`
	GuiSettings    *txml.Node  `xml:"GuiSettings"`
	Properties     *Properties `xml:"Properties"`
	EngineSettings *txml.Node  `xml:"EngineSettings,omitempty"`
}

func (ryxNode *RyxNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = `Node`
	guiSettings := ryxNode.GuiSettings
	if guiSettings.IsNil() {
		guiSettings = nil
	}
	engineSettings := ryxNode.EngineSettings
	if engineSettings.IsNil() {
		engineSettings = nil
	}
	childNodes := ryxNode.ChildNodes
	if len(childNodes) == 0 {
		container := &_RyxNodeXml{
			ToolId:         ryxNode.ToolId,
			GuiSettings:    guiSettings,
			Properties:     ryxNode.Properties,
			EngineSettings: engineSettings,
		}
		return e.EncodeElement(container, start)
	}
	container := &_RyxNodeWithChildrenXml{
		ToolId:         ryxNode.ToolId,
		GuiSettings:    guiSettings,
		Properties:     ryxNode.Properties,
		EngineSettings: engineSettings,
		ChildNodes:     childNodes,
	}
	return e.EncodeElement(container, start)
}

func (ryxNode *RyxNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	container := &_RyxNodeWithChildrenXml{}
	err := d.DecodeElement(container, &start)
	if err != nil {
		return err
	}

	guiSettings := txml.NilNode()
	if container.GuiSettings != nil {
		guiSettings = container.GuiSettings
	}
	engineSettings := txml.NilNode()
	if container.EngineSettings != nil {
		engineSettings = container.EngineSettings
	}
	ryxNode.ToolId = container.ToolId
	ryxNode.GuiSettings = guiSettings
	ryxNode.EngineSettings = engineSettings
	ryxNode.Properties = container.Properties
	ryxNode.ChildNodes = container.ChildNodes
	return nil
}
