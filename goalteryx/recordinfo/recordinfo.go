package recordinfo

import (
	"fmt"
	"github.com/tlarsen7572/Golang-Public/goalteryx/convert_strings"
	"strconv"
	"unsafe"
)

type RecordInfo interface {
	NumFields() int
	GetFieldByIndex(index int) (FieldInfo, error)
	AddByteField(name string, source string) string
	AddBoolField(name string, source string) string
	AddInt16Field(name string, source string) string
	AddInt32Field(name string, source string) string
	AddInt64Field(name string, source string) string
	AddFixedDecimalField(name string, source string, size int, precision int) string
	AddFloatField(name string, source string) string
	AddDoubleField(name string, source string) string
	AddStringField(name string, source string, size int) string
	AddWStringField(name string, source string, size int) string
	AddV_StringField(name string, source string, size int) string
	AddV_WStringField(name string, source string, size int) string
	AddDateField(name string, source string) string
	AddDateTimeField(name string, source string) string
	GetByteValueFrom(fieldName string, record unsafe.Pointer) (byte, error)
	GetBoolValueFrom(fieldName string, record unsafe.Pointer) (bool, error)
	GetInt16ValueFrom(fieldName string, record unsafe.Pointer) (int16, error)
	GetInt32ValueFrom(fieldName string, record unsafe.Pointer) (int32, error)
	GetInt64ValueFrom(fieldName string, record unsafe.Pointer) (int64, error)
	GetFixedDecimalValueFrom(fieldName string, record unsafe.Pointer) (float64, error)
	GetFloatValueFrom(fieldName string, record unsafe.Pointer) (float32, error)
	GetDoubleValueFrom(fieldName string, record unsafe.Pointer) (float64, error)
}

type recordInfo struct {
	currentLen uintptr
	numFields  int
	fields     []FieldInfo
	fieldNames map[string]int
}

var ByteType = `byte`
var BoolType = `bool`
var Int16Type = `int16`
var Int32Type = `int32`
var Int64Type = `int64`
var FixedDecimalType = `fixeddecimal`
var FloatType = `float`
var DoubleType = `double`
var StringType = `string`
var WStringType = `wstring`
var V_StringType = `v_string`
var V_WStringType = `v_wstring`
var DateType = `date`
var DateTimeType = `datetime`

type FieldInfo struct {
	Name        string
	Source      string
	Size        int
	Precision   int
	Type        string
	location    uintptr
	fixedLen    uintptr
	nullByteLen uintptr
}

func New() RecordInfo {
	return &recordInfo{fieldNames: map[string]int{}}
}

func (info *recordInfo) NumFields() int {
	return info.numFields
}

func (info *recordInfo) GetFieldByIndex(index int) (FieldInfo, error) {
	if count := len(info.fields); index < 0 || index >= count {
		return FieldInfo{}, fmt.Errorf(`index was not between 0 and %v`, count)
	}
	return info.fields[index], nil
}

func (info *recordInfo) AddByteField(name string, source string) string {
	return info.addField(name, source, 1, 0, ByteType, 1, 1)
}

func (info *recordInfo) AddBoolField(name string, source string) string {
	return info.addField(name, source, 1, 0, BoolType, 1, 0)
}

func (info *recordInfo) AddInt16Field(name string, source string) string {
	return info.addField(name, source, 2, 0, Int16Type, 2, 1)
}

func (info *recordInfo) AddInt32Field(name string, source string) string {
	return info.addField(name, source, 4, 0, Int32Type, 4, 1)
}

func (info *recordInfo) AddInt64Field(name string, source string) string {
	return info.addField(name, source, 8, 0, Int64Type, 8, 1)
}

func (info *recordInfo) AddFixedDecimalField(name string, source string, size int, precision int) string {
	return info.addField(name, source, size, precision, FixedDecimalType, uintptr(size), 1)
}

func (info *recordInfo) AddFloatField(name string, source string) string {
	return info.addField(name, source, 4, 0, FloatType, 4, 1)
}

func (info *recordInfo) AddDoubleField(name string, source string) string {
	return info.addField(name, source, 8, 0, DoubleType, 8, 1)
}

func (info *recordInfo) AddStringField(name string, source string, size int) string {
	return info.addField(name, source, size, 0, StringType, uintptr(size), 1)
}

func (info *recordInfo) AddWStringField(name string, source string, size int) string {
	return info.addField(name, source, size, 0, WStringType, uintptr(size), 1)
}

func (info *recordInfo) AddV_StringField(name string, source string, size int) string {
	return info.addField(name, source, size, 0, V_StringType, 4, 0)
}

func (info *recordInfo) AddV_WStringField(name string, source string, size int) string {
	return info.addField(name, source, size, 0, V_WStringType, 4, 0)
}

func (info *recordInfo) AddDateField(name string, source string) string {
	return info.addField(name, source, 10, 0, DateType, 10, 1)
}

func (info *recordInfo) AddDateTimeField(name string, source string) string {
	return info.addField(name, source, 19, 0, DateTimeType, 19, 1)
}

func (info *recordInfo) GetByteValueFrom(fieldName string, record unsafe.Pointer) (byte, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	return *((*byte)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) GetBoolValueFrom(fieldName string, record unsafe.Pointer) (bool, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return false, err
	}
	return *((*bool)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) GetInt16ValueFrom(fieldName string, record unsafe.Pointer) (int16, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	return *((*int16)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) GetInt32ValueFrom(fieldName string, record unsafe.Pointer) (int32, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	return *((*int32)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) GetInt64ValueFrom(fieldName string, record unsafe.Pointer) (int64, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	return *((*int64)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) GetFixedDecimalValueFrom(fieldName string, record unsafe.Pointer) (float64, error) {
	fieldPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	valueStr := convert_strings.CToString(unsafe.Pointer(fieldPtr))
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, fmt.Errorf(`error converting '%v' to double in '%v' field`, value, fieldName)
	}
	return value, nil
}

func (info *recordInfo) GetFloatValueFrom(fieldName string, record unsafe.Pointer) (float32, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	return *((*float32)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) GetDoubleValueFrom(fieldName string, record unsafe.Pointer) (float64, error) {
	recordPtr, err := info.getFieldLocationPtr(fieldName, record)
	if err != nil {
		return 0, err
	}
	return *((*float64)(unsafe.Pointer(recordPtr))), nil
}

func (info *recordInfo) addField(name string, source string, size int, scale int, fieldType string, fixedLen uintptr, nullByteLen uintptr) string {
	actualName := info.checkFieldName(name)
	info.fields = append(info.fields, FieldInfo{
		Name:        actualName,
		Source:      source,
		Size:        size,
		Precision:   scale,
		Type:        fieldType,
		location:    info.currentLen,
		fixedLen:    fixedLen,
		nullByteLen: nullByteLen,
	})
	info.fieldNames[actualName] = info.numFields
	info.numFields++
	info.currentLen += fixedLen + nullByteLen
	return actualName
}

func (info *recordInfo) getFieldLocationPtr(fieldName string, record unsafe.Pointer) (uintptr, error) {
	index, ok := info.fieldNames[fieldName]
	if !ok {
		return 0, fmt.Errorf(`field '%v' does not exist`, fieldName)
	}
	field := info.fields[index]
	return uintptr(record) + field.location, nil
}

func (info *recordInfo) checkFieldName(name string) string {
	_, exists := info.fieldNames[name]
	for exists {
		name = name + `2`
		_, exists = info.fieldNames[name]
	}
	return name
}
