package recordinfo

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type xmlMetaInfo struct {
	ToolId string      `xml:"MetaInfo"`
	Fields []*xmlField `xml:"RecordInfo>Field"`
}

type xmlField struct {
	Name   string `xml:"name,attr"`
	Source string `xml:"source,attr"`
	Size   string `xml:"size,attr"`
	Type   string `xml:"type,attr"`
}

func FromXml(recordInfoXml string) (RecordInfo, error) {
	var metaInfo xmlMetaInfo
	err := xml.Unmarshal([]byte(recordInfoXml), &metaInfo)
	if err != nil {
		return nil, fmt.Errorf(`error creating RecordInfo from xml: %v`, err.Error())
	}
	recordInfo := New()
	for index, field := range metaInfo.Fields {
		switch strings.ToLower(field.Type) {
		case BoolType:
			recordInfo.AddBoolField(field.Name, field.Source)
		case Int64Type:
			recordInfo.AddInt64Field(field.Name, field.Source)
		case StringType:
			size, err := strconv.Atoi(field.Size)
			if err != nil {
				return nil, fmt.Errorf(`error converting field %v size to an int.  Provided size was %v`, index, field.Size)
			}
			recordInfo.AddStringField(field.Name, field.Source, size)
		case V_WStringType:
			size, err := strconv.Atoi(field.Size)
			if err != nil {
				return nil, fmt.Errorf(`error converting field %v size to an int.  Provided size was %v`, index, field.Size)
			}
			recordInfo.AddV_WStringField(field.Name, field.Source, size)
		default:
			continue
		}
	}
	return recordInfo, nil
}
