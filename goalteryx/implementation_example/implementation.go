package main

/*
#include "implementation.h"
*/
import "C"
import (
	"encoding/xml"
	"fmt"
	"github.com/tlarsen7572/Golang-Public/goalteryx/api"
	"github.com/tlarsen7572/Golang-Public/goalteryx/recordinfo"
	"unsafe"
)

func main() {}

//export AlteryxGoPlugin
func AlteryxGoPlugin(toolId C.int, xmlProperties unsafe.Pointer, engineInterface unsafe.Pointer, pluginInterface unsafe.Pointer) C.long {
	myPlugin := &MyNewPlugin{}
	return C.long(api.ConfigurePlugin(myPlugin, int(toolId), xmlProperties, engineInterface, pluginInterface))
}

type MyNewPlugin struct {
	ToolId int
	Field  string
}

type ConfigXml struct {
	Field string `xml:"Field"`
}

func (plugin *MyNewPlugin) Init(toolId int, config string) bool {
	plugin.ToolId = toolId
	api.OutputMessage(plugin.ToolId, 1, fmt.Sprintf(`Tool configuration: %v`, config))
	var c ConfigXml
	err := xml.Unmarshal([]byte(config), &c)
	if err != nil {
		api.OutputMessage(toolId, 3, err.Error())
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

func (plugin *MyNewPlugin) AddIncomingConnection(connectionType string, connectionName string) api.IncomingInterface {
	return &MyNewIncomingInterface{Parent: plugin}
}

func (plugin *MyNewPlugin) AddOutgoingConnection(connectionName string) bool {
	api.OutputMessage(plugin.ToolId, 1, fmt.Sprintf(`Add outgoing connection: %v`, connectionName))
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
		api.OutputMessage(ii.Parent.ToolId, 3, err.Error())
		return false
	}
	for _, connection := range []string{`Output1`, `Blah`} {
		err = api.InitOutput(connection, ii.inInfo)
		if err != nil {
			api.OutputMessage(ii.Parent.ToolId, 3, err.Error())
			return false
		}
	}
	api.OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`Incoming record info: %v`, recordInfoIn))
	return true
}

func (ii *MyNewIncomingInterface) PushRecord(record unsafe.Pointer) bool {
	var value interface{}
	var isNull bool
	var err error
	value, isNull, err = ii.inInfo.GetInterfaceValueFrom(ii.Parent.Field, record)
	if err != nil {
		api.OutputMessage(ii.Parent.ToolId, 3, err.Error())
		return false
	}
	if isNull {
		api.OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`[%v] is null`, ii.Parent.Field))
	} else {
		api.OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`[%v] is %v`, ii.Parent.Field, value))
	}
	for _, connection := range []string{`Output1`, `Blah`} {
		err = api.PushRecord(connection, record)
		if err != nil {
			api.OutputMessage(ii.Parent.ToolId, 3, err.Error())
		}
	}
	return true
}

func (ii *MyNewIncomingInterface) UpdateProgress(percent float64) {

}

func (ii *MyNewIncomingInterface) Close() {

}

func (ii *MyNewIncomingInterface) Free() {

}
