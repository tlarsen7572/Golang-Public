package main

/*
#include "plugins.h"
*/
import "C"
import "unicode/utf16"
import "unsafe"
import "github.com/mattn/go-pointer"

func main() {

}

var MyPlugin Plugin

type Plugin interface {
	PushAllRecords(recordLimit int) int
}

//export PiPushAllRecords
func PiPushAllRecords(handle unsafe.Pointer, recordLimit C.__int64) C.long {
	alteryxPlugin := pointer.Restore(handle).(Plugin)
	return C.long(alteryxPlugin.PushAllRecords(int(recordLimit)))
}

//export GetPlugin
func GetPlugin() unsafe.Pointer {
	return pointer.Save(MyPlugin)
}

func UTF16ToString(s []uint16) string {
	for i, v := range s {
		if v == 0 {
			s = s[0:i]
			break
		}
	}
	return string(utf16.Decode(s))
}
