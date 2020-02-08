package ryxdoc

import (
	"encoding/xml"
	"strconv"
)

type _RyxConnXml struct {
	XMLName  xml.Name `xml:"Connection"`
	Name     string   `xml:"name,attr"`
	Wireless bool     `xml:"Wireless,attr"`
	Origin   struct {
		ToolId     string `xml:"ToolID,attr"`
		Connection string `xml:"Connection,attr"`
	} `xml:"Origin"`
	Destination struct {
		ToolId     string `xml:"ToolID,attr"`
		Connection string `xml:"Connection,attr"`
	} `xml:"Destination"`
}

func (ryxConn *RyxConn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = `Connection`
	container := _RyxConnXml{
		Name:     ryxConn.Name,
		Wireless: ryxConn.Wireless,
		Origin: struct {
			ToolId     string `xml:"ToolID,attr"`
			Connection string `xml:"Connection,attr"`
		}{
			ToolId:     strconv.Itoa(ryxConn.FromId),
			Connection: ryxConn.FromAnchor,
		},
		Destination: struct {
			ToolId     string `xml:"ToolID,attr"`
			Connection string `xml:"Connection,attr"`
		}{
			ToolId:     strconv.Itoa(ryxConn.ToId),
			Connection: ryxConn.ToAnchor,
		},
	}
	return e.EncodeElement(container, start)
}

func (ryxConn *RyxConn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	container := &_RyxConnXml{}
	err := d.DecodeElement(container, &start)
	if err != nil {
		return err
	}
	fromId, err := strconv.Atoi(container.Origin.ToolId)
	if err != nil {
		return err
	}
	toId, err := strconv.Atoi(container.Destination.ToolId)
	if err != nil {
		return err
	}

	ryxConn.Name = container.Name
	ryxConn.Wireless = container.Wireless
	ryxConn.FromId = fromId
	ryxConn.FromAnchor = container.Origin.Connection
	ryxConn.ToId = toId
	ryxConn.ToAnchor = container.Destination.Connection
	return nil
}
