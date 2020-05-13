package main

/*
#include "plugins.h"
*/
import "C"
import (
	"fmt"
	"github.com/mattn/go-pointer"
	"github.com/vitaminwater/cgo.wchar"
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
	PushAllRecords(recordLimit int) int
	Close(hasErrors bool)
	AddIncomingConnection(connectionType string, connectionName string) IncomingInterface
	AddOutgoingConnection(connectionName string) bool
}

type IncomingInterface interface {
	Init(recordInfoIn string) bool
}

//export AlteryxGoPlugin
func AlteryxGoPlugin(toolId C.int, pXmlProperties unsafe.Pointer, pEngineInterface *C.struct_EngineInterface, r_pluginInterface *C.struct_PluginInterface) C.long {
	Engine = pEngineInterface
	config, err := wchar.WcharStringPtrToGoString(pXmlProperties)
	if err != nil {
		printLogf(`error converting pXmlProperties to string in AlteryxGoPlugin: %v`, err.Error())
		return C.long(0)
	}
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
	return C.long(alteryxPlugin.PushAllRecords(int(recordLimit)))
}

//export PiClose
func PiClose(handle unsafe.Pointer, hasErrors C.bool) {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	alteryxPlugin.Close(bool(hasErrors))
}

//export PiAddIncomingConnection
func PiAddIncomingConnection(handle unsafe.Pointer, connectionType unsafe.Pointer, connectionName unsafe.Pointer, incomingInterface *C.struct_IncomingConnectionInterface) C.long {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	goName, err := wchar.WcharStringPtrToGoString(connectionName)
	if err != nil {
		printLogf(`error converting connectionName to string in PiAddIncomingConnection: %v`, err.Error())
	}
	goType, err := wchar.WcharStringPtrToGoString(connectionType)
	if err != nil {
		printLogf(`error converting connectionType to string in PiAddIncomingConnection: %v`, err.Error())
	}
	goIncomingInterface := alteryxPlugin.AddIncomingConnection(goType, goName)
	iiHandle := pointer.Save(goIncomingInterface)
	incomingInterface.handle = iiHandle
	incomingInterface.pII_Init = C.T_II_Init(C.IiInit)
	return C.long(1)
}

//export PiAddOutgoingConnection
func PiAddOutgoingConnection(handle unsafe.Pointer, connectionName unsafe.Pointer, incomingConnection *C.struct_IncomingConnectionInterface) C.long {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	goName, err := wchar.WcharStringPtrToGoString(connectionName)
	if err != nil {
		printLogf(`error converting connectionName to string in PiAddOutgoingConnection: %v`, err.Error())
	}
	if alteryxPlugin.AddOutgoingConnection(goName) {
		return C.long(1)
	}
	return C.long(0)
}

//export IiInit
func IiInit(handle unsafe.Pointer, recordInfoIn unsafe.Pointer) C.long {
	incomingInterface := pointer.Restore(handle).(IncomingInterface)
	goRecordInfoIn, err := wchar.WcharStringPtrToGoString(recordInfoIn)
	if err != nil {
		printLogf(`error converting recordInfoIn to string in IiInit: %v`, err.Error())
	}
	if incomingInterface.Init(goRecordInfoIn) {
		return C.long(1)
	}
	return C.long(0)
}

//export GetPlugin
func GetPlugin() unsafe.Pointer {
	return pointer.Save(MyPlugin)
}

func OutputMessage(toolId int, status int, message string) {
	cMessage, err := wchar.FromGoString(message)
	if err != nil {
		printLogf(`error converting message to wcharstring in OutputMessage: %v`, err.Error())
		return
	}
	if cMessage == nil {
		return
	}
	printLogf(`getting ready to call output message`)
	C.callEngineOutputMessage(Engine, C.int(toolId), C.int(status), unsafe.Pointer(&cMessage[0]))
}

func printLogf(message string, args ...interface{}) {
	file, _ := os.OpenFile("C:\\temp\\output.txt", os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	file.WriteString(fmt.Sprintf(time.Now().String()+": "+message+"\n", args...))
}
