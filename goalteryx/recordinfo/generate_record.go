package recordinfo

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"syscall"
	"time"
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
	data := make([]byte, 3)
	nullByte := 2
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(int16)
	binary.LittleEndian.PutUint16(data, uint16(value))
	return data, nil
}

func generateInt32(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, 5)
	nullByte := 4
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(int32)
	binary.LittleEndian.PutUint32(data, uint32(value))
	return data, nil
}

func generateInt64(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, 9)
	nullByte := 8
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(int64)
	binary.LittleEndian.PutUint64(data, uint64(value))
	return data, nil
}

func generateFixedDecimal(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, field.fixedLen+1)
	nullByte := int(field.fixedLen)
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(float64)
	format := `%` + fmt.Sprintf(`%v.%vf`, field.Size, field.Precision)
	valueStr := strings.TrimSpace(fmt.Sprintf(format, value))
	for index := range valueStr {
		data[index] = valueStr[index]
	}
	return data, nil
}

func generateFloat32(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, 5)
	nullByte := 4
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(float32)
	binary.LittleEndian.PutUint32(data, math.Float32bits(value))
	return data, nil
}

func generateFloat64(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, 9)
	nullByte := 8
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(float64)
	binary.LittleEndian.PutUint64(data, math.Float64bits(value))
	return data, nil
}

func generateString(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, field.fixedLen+1)
	nullByte := field.fixedLen
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(string)
	for index := range value {
		data[index] = value[index]
	}
	return data, nil
}

func generateWString(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, field.fixedLen+1)
	nullByte := field.fixedLen
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(string)
	valueChars, err := syscall.UTF16FromString(value)
	if err != nil {
		return nil, err
	}
	for index := range valueChars {
		char := valueChars[index]
		if char == 0 {
			break
		}
		byteIndex := index * 2
		binary.LittleEndian.PutUint16(data[byteIndex:byteIndex+2], char)
	}
	return data, nil
}

func generateDate(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, field.fixedLen+1)
	nullByte := field.fixedLen
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(time.Time)
	valueStr := value.Format(dateFormat)
	for index := range valueStr {
		data[index] = valueStr[index]
	}
	return data, nil
}

func generateDateTime(field *fieldInfoEditor) ([]byte, error) {
	data := make([]byte, field.fixedLen+1)
	nullByte := field.fixedLen
	if field.value == nil {
		data[nullByte] = byte(1)
		return data, nil
	}
	value := field.value.(time.Time)
	valueStr := value.Format(dateTimeFormat)
	for index := range valueStr {
		data[index] = valueStr[index]
	}
	return data, nil
}
