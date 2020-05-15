package recordinfo

import (
	"fmt"
	"time"
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

	GetByteValueFrom(fieldName string, record unsafe.Pointer) (value byte, isNull bool, err error)
	GetBoolValueFrom(fieldName string, record unsafe.Pointer) (value bool, isNull bool, err error)
	GetInt16ValueFrom(fieldName string, record unsafe.Pointer) (value int16, isNull bool, err error)
	GetInt32ValueFrom(fieldName string, record unsafe.Pointer) (value int32, isNull bool, err error)
	GetInt64ValueFrom(fieldName string, record unsafe.Pointer) (value int64, isNull bool, err error)
	GetFixedDecimalValueFrom(fieldName string, record unsafe.Pointer) (value float64, isNull bool, err error)
	GetFloatValueFrom(fieldName string, record unsafe.Pointer) (value float32, isNull bool, err error)
	GetDoubleValueFrom(fieldName string, record unsafe.Pointer) (value float64, isNull bool, err error)
	GetStringValueFrom(fieldName string, record unsafe.Pointer) (value string, isNull bool, err error)
	GetWStringValueFrom(fieldName string, record unsafe.Pointer) (value string, isNull bool, err error)
	GetDateValueFrom(fieldName string, record unsafe.Pointer) (value time.Time, isNull bool, err error)
	GetDateTimeValueFrom(fieldName string, record unsafe.Pointer) (value time.Time, isNull bool, err error)
}

type recordInfo struct {
	currentLen uintptr
	numFields  int
	fields     []FieldInfo
	fieldNames map[string]int
}

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

func (info *recordInfo) getFieldInfo(fieldName string) (FieldInfo, error) {
	index, ok := info.fieldNames[fieldName]
	if !ok {
		return FieldInfo{}, fmt.Errorf(`field '%v' does not exist`, fieldName)
	}
	field := info.fields[index]
	return field, nil
}

func (info *recordInfo) checkFieldName(name string) string {
	_, exists := info.fieldNames[name]
	for exists {
		name = name + `2`
		_, exists = info.fieldNames[name]
	}
	return name
}
