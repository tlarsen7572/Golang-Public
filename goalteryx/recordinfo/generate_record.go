package recordinfo

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func (info *recordInfo) GenerateRecord() (unsafe.Pointer, error) {
	data := make([]byte, 0)
	for _, field := range info.fields {
		getBytes := field.generator
		if getBytes == nil {
			return nil, fmt.Errorf(`field '%v' does not have a byte generator`, field.Name)
		}
		fieldData, err := getBytes(field)
		if err != nil {
			return nil, err
		}
		data = append(data, fieldData...)
	}
	return unsafe.Pointer(&data[0]), nil
}

func generateByte(field *fieldInfoEditor) ([]byte, error) {
	nullByte := byte(0)
	if field.value == nil {
		nullByte = byte(1)
	}
	return []byte{field.value.(byte), nullByte}, nil
}

func generateBool(field *fieldInfoEditor) ([]byte, error) {
	if field.value == nil {
		return []byte{byte(2)}, nil
	}
	if field.value.(bool) == true {
		return []byte{byte(1)}, nil
	}
	return []byte{byte(0)}, nil
}

func generateInt16(field *fieldInfoEditor) ([]byte, error) {
	nullByte := byte(0)
	if field.value == nil {
		nullByte = byte(1)
	}
	value := field.value.(int16)
	data := make([]byte, 3)
	binary.LittleEndian.PutUint16(data, uint16(value))
	data[2] = nullByte
	return data, nil
}
