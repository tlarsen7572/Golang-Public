package main

/*
#include "plugins.h"
*/
import "C"
import (
	"fmt"
	"github.com/mattn/go-pointer"
	"github.com/tlarsen7572/Golang-Public/goalteryx/convert_strings"
	"os"
	"time"
	"unsafe"
)

func main() {

}

var MyPlugin Plugin
var Engine *C.struct_EngineInterface

type Plugin interface {
	Init(toolId int, config string) bool
	PushAllRecords(recordLimit int) bool
	Close(hasErrors bool)
	AddIncomingConnection(connectionType string, connectionName string) IncomingInterface
	AddOutgoingConnection(connectionName string) bool
}

type IncomingInterface interface {
	Init(recordInfoIn string) bool
	PushRecord(record unsafe.Pointer) bool
	UpdateProgress(percent float64)
	Close()
	Free()
}

//export AlteryxGoPlugin
func AlteryxGoPlugin(toolId C.int, pXmlProperties unsafe.Pointer, pEngineInterface *C.struct_EngineInterface, r_pluginInterface *C.struct_PluginInterface) C.long {
	Engine = pEngineInterface
	config := convert_strings.WideCToString(pXmlProperties)
	printLogf(`converted config in AlteryxGoPlugin: %v`, config)
	MyPlugin = &MyNewPlugin{}
	if !MyPlugin.Init(int(toolId), config) {
		return C.long(0)
	}

	r_pluginInterface.handle = GetPlugin()
	r_pluginInterface.pPI_PushAllRecords = C.T_PI_PushAllRecords(C.PiPushAllRecords)
	r_pluginInterface.pPI_Close = C.T_PI_Close(C.PiClose)
	r_pluginInterface.pPI_AddIncomingConnection = C.T_PI_AddIncomingConnection(C.PiAddIncomingConnection)
	r_pluginInterface.pPI_AddOutgoingConnection = C.T_PI_AddOutgoingConnection(C.PiAddOutgoingConnection)
	printLogf(`hooked up PluginInterface`)
	return C.long(1)
}

//export PiPushAllRecords
func PiPushAllRecords(handle unsafe.Pointer, recordLimit C.__int64) C.long {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	if alteryxPlugin.PushAllRecords(int(recordLimit)) {
		return C.long(1)
	}
	return C.long(0)
}

//export PiClose
func PiClose(handle unsafe.Pointer, hasErrors C.bool) {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	alteryxPlugin.Close(bool(hasErrors))
}

//export PiAddIncomingConnection
func PiAddIncomingConnection(handle unsafe.Pointer, connectionType unsafe.Pointer, connectionName unsafe.Pointer, incomingInterface *C.struct_IncomingConnectionInterface) C.long {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	goName := convert_strings.WideCToString(connectionName)
	goType := convert_strings.WideCToString(connectionType)
	goIncomingInterface := alteryxPlugin.AddIncomingConnection(goType, goName)
	iiHandle := pointer.Save(goIncomingInterface)
	incomingInterface.handle = iiHandle
	incomingInterface.pII_Init = C.T_II_Init(C.IiInit)
	incomingInterface.pII_PushRecord = C.T_II_PushRecord(C.IiPushRecord)
	incomingInterface.pII_UpdateProgress = C.T_II_UpdateProgress(C.IiUpdateProgress)
	incomingInterface.pII_Close = C.T_II_Close(C.IiClose)
	incomingInterface.pII_Free = C.T_II_Free(C.IiFree)
	return C.long(1)
}

//export PiAddOutgoingConnection
func PiAddOutgoingConnection(handle unsafe.Pointer, connectionName unsafe.Pointer, incomingConnection *C.struct_IncomingConnectionInterface) C.long {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	goName := convert_strings.WideCToString(connectionName)
	if alteryxPlugin.AddOutgoingConnection(goName) {
		return C.long(1)
	}
	return C.long(0)
}

//export IiInit
func IiInit(handle unsafe.Pointer, recordInfoIn unsafe.Pointer) C.long {
	incomingInterface := pointer.Restore(handle).(IncomingInterface)
	goRecordInfoIn := convert_strings.WideCToString(recordInfoIn)
	if incomingInterface.Init(goRecordInfoIn) {
		return C.long(1)
	}
	return C.long(0)
}

//export IiPushRecord
func IiPushRecord(handle unsafe.Pointer, record unsafe.Pointer) C.long {
	incomingInterface := pointer.Restore(handle).(IncomingInterface)
	if incomingInterface.PushRecord(record) {
		return C.long(1)
	}
	return C.long(0)
}

//export IiUpdateProgress
func IiUpdateProgress(handle unsafe.Pointer, percent C.double) {
	incomingInterface := pointer.Restore(handle).(IncomingInterface)
	incomingInterface.UpdateProgress(float64(percent))
}

//export IiClose
func IiClose(handle unsafe.Pointer) {
	incomingInterface := pointer.Restore(handle).(IncomingInterface)
	incomingInterface.Close()
}

//export IiFree
func IiFree(handle unsafe.Pointer) {
	incomingInterface := pointer.Restore(handle).(IncomingInterface)
	incomingInterface.Free()
}

//export GetPlugin
func GetPlugin() unsafe.Pointer {
	return pointer.Save(MyPlugin)
}

func OutputMessage(toolId int, status int, message string) {
	cMessage, err := convert_strings.StringToWideC(message)
	if err != nil {
		printLogf(`error converting message to wcharstring in OutputMessage: %v`, err.Error())
		return
	}
	if cMessage == nil {
		return
	}

	printLogf(`getting ready to call output message with ` + message)
	C.callEngineOutputMessage(Engine, C.int(toolId), C.int(status), cMessage)
}

func printLogf(message string, args ...interface{}) {
	file, _ := os.OpenFile("C:\\temp\\output.txt", os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	file.WriteString(fmt.Sprintf(time.Now().String()+": "+message+"\n", args...))
}
