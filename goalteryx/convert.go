package main

/*
#cgo darwin LDFLAGS: -liconv
#cgo windows LDFLAGS: -liconv
#include <stdlib.h>
#ifdef __APPLE__
#  define LIBICONV_PLUG 1
#endif
#include <iconv.h>
#include <wchar.h>
#include <string.h>
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

// iconv charset strings
var (
	iconvCharsetWchar = C.CString("wchar_t//TRANSLIT")
	iconvCharsetChar  = C.CString("//TRANSLIT")
	iconvCharsetAscii = C.CString("ascii//TRANSLIT")
	iconvCharsetUtf8  = C.CString("utf-8//TRANSLIT")
)

// Copying and modifying wchar convertGoStringToWcharString function to return an unsafe pointer which will be
// released by Alteryx
func convertGoStringToWcharString(input string) (outputPointer unsafe.Pointer, err error) {
	// open iconv
	iconv, errno := C.iconv_open(iconvCharsetWchar, iconvCharsetUtf8)
	if iconv == nil || errno != nil {
		return nil, fmt.Errorf("Could not open iconv instance: %s", errno)
	}
	defer C.iconv_close(iconv)

	// calculate bufferSizes in bytes for C
	bytesLeftInCSize := C.size_t(len([]byte(input))) // count exact amount of bytes from input
	bytesLeftOutCSize := C.size_t(len(input) * 4)    // wide char seems to be 4 bytes for every single- or multi-byte character. Not very sure though.

	// input for C. makes a copy using C malloc and therefore should be free'd.
	inputCString := C.CString(input)
	defer C.free(unsafe.Pointer(inputCString))

	// output for C
	outputCString := (*C.char)(C.malloc(bytesLeftOutCSize))
	defer C.free(unsafe.Pointer(outputCString))

	// call iconv for conversion of charsets, return on error
	saveInputCString, saveOutputCString := inputCString, outputCString
	_, errno = C.iconv(iconv, &inputCString, &bytesLeftInCSize, &outputCString, &bytesLeftOutCSize)
	if errno != nil {
		return nil, errno
	}
	inputCString, outputCString = saveInputCString, saveOutputCString

	outputLen := len(input)*4 - int(bytesLeftOutCSize)
	outputChars := make([]int8, outputLen)
	C.memcpy(unsafe.Pointer(&outputChars[0]), unsafe.Pointer(outputCString), C.size_t(outputLen))

	// convert []int8 to WcharString
	// create WcharString with same length as input, and one extra position for the null terminator.
	output := make([]uint32, 0, len(input)+1)
	// create buff to convert each outputChar
	wcharAsByteAry := make([]byte, 4)
	// loop for as long as there are output chars
	for len(outputChars) >= 4 {
		// create 4 position byte slice
		wcharAsByteAry[0] = byte(outputChars[0])
		wcharAsByteAry[1] = byte(outputChars[1])
		wcharAsByteAry[2] = byte(outputChars[2])
		wcharAsByteAry[3] = byte(outputChars[3])
		// combine 4 position byte slice into uint32
		wcharAsUint32 := binary.LittleEndian.Uint32(wcharAsByteAry)
		// find null terminator (doing this right?)
		if wcharAsUint32 == 0x0 {
			break
		}
		// append uint32 to outputUint32
		output = append(output, wcharAsUint32)
		// reslice the outputChars
		outputChars = outputChars[4:]
	}
	// Add null terminator
	output = append(output, 0x0)

	return unsafe.Pointer(&output[0]), nil
}
