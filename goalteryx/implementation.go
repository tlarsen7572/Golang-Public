package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

type MyNewPlugin struct {
	ToolId int
}

func (plugin *MyNewPlugin) Init(toolId int, config string) bool {
	plugin.ToolId = toolId
	OutputMessage(plugin.ToolId, 1, fmt.Sprintf(`Tool configuration: %v`, config))
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
}

func (ii *MyNewIncomingInterface) Init(recordInfoIn string) bool {
	OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`Incoming record info: %v`, recordInfoIn))
	return true
}

func (ii *MyNewIncomingInterface) PushRecord(record unsafe.Pointer) bool {
	ptr := uintptr(record)
	recordBytes := make([]byte, 0)
	for index := 0; index < 1850; index++ {
		singleByte := *((*byte)(unsafe.Pointer(ptr)))
		recordBytes = append(recordBytes, singleByte)
		ptr += 1
	}
	OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`Record bytes: %v`, recordBytes))
	OutputMessage(ii.Parent.ToolId, 1, fmt.Sprintf(`First field: %v`, binary.LittleEndian.Uint64(recordBytes[0:8])))
	return true
}

func (ii *MyNewIncomingInterface) UpdateProgress(percent float64) {

}

func (ii *MyNewIncomingInterface) Close() {

}

func (ii *MyNewIncomingInterface) Free() {

}
