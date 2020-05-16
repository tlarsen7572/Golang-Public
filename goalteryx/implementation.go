package main

import (
	"encoding/xml"
	"fmt"
	"github.com/tlarsen7572/Golang-Public/goalteryx/recordinfo"
	"unsafe"
)

type MyNewPlugin struct {
	ToolId int
	Field  string
}

type ConfigXml struct {
	Field string `xml:"Field"`
}

func (plugin *MyNewPlugin) Init(toolId int, config string) bool {
	plugin.ToolId = toolId
	OutputMessage(plugin.ToolId, 1, fmt.Sprintf(`Tool configuration: %v`, config))
	var c ConfigXml
	err := xml.Unmarshal([]byte(config), &c)
	if err != nil {
		OutputMessage(toolId, 3, err.Error())
		return false
	}
	plugin.Field = c.Field
	return true
}

func (plugin *MyNewPlugin) PushAllRecords(recordLimit int) bool {
	return true
}

func (plugin *MyNewPlugin) Close(hasErrors bool) {

}

func (plugin *MyNewPlugin) AddIncomingConnection(connectionType string, connectionName string) IncomingInterface {
	return &MyNewIncomingInterface{Parent: plugin}
}

func (plugin *MyNewPlugin) AddOutgoingConnection(connectionName string) bool {
	OutputMessage(plugin.ToolId, 1, fmt.Sprintf(`Add outgoing connection: %v`, connectionName))
	return true
}

type MyNewIncomingInterface struct {
	Parent *MyNewPlugin
	inInfo recordinfo.RecordInfo
}

func (ii *MyNewIncomingInterface) Init(recordInfoIn string) bool {
	var err error
	ii.inInfo, err = recordinfo.FromXml(recordInfoIn)
	if err != nil {
		OutputMessage(ii.Parent.ToolId, 3, err.Error())
		return false
	}
	OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`Incoming record info: %v`, recordInfoIn))
	return true
}

func (ii *MyNewIncomingInterface) PushRecord(record unsafe.Pointer) bool {
	var value interface{}
	var isNull bool
	var err error
	value, isNull, err = ii.inInfo.GetInterfaceValueFrom(ii.Parent.Field, record)
	if err != nil {
		OutputMessage(ii.Parent.ToolId, 3, err.Error())
		return false
	}
	if isNull {
		OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`[%v] is null`, ii.Parent.Field))
	} else {
		OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`[%v] is %v`, ii.Parent.Field, value))
	}
	return true
}

func (ii *MyNewIncomingInterface) UpdateProgress(percent float64) {

}

func (ii *MyNewIncomingInterface) Close() {

}

func (ii *MyNewIncomingInterface) Free() {

}
