package recordinfo

import "fmt"

type RecordInfo interface {
	NumFields() int
	GetFieldByIndex(index int) (FieldInfo, error)
	AddBoolField(name string, source string) string
	AddInt64Field(name string, source string) string
	AddStringField(name string, source string, size int) string
	AddV_WStringField(name string, source string, size int) string
}

type recordInfo struct {
	numFields  int
	fields     []FieldInfo
	fieldNames map[string]bool
}

var BoolType = `bool`
var Int64Type = `int64`
var StringType = `string`
var V_WStringType = `v_wstring`

type FieldInfo struct {
	Name      string
	Source    string
	Size      int
	Precision int
	Type      string
}

func New() RecordInfo {
	return &recordInfo{fieldNames: map[string]bool{}}
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

func (info *recordInfo) AddBoolField(name string, source string) string {
	name = info.checkFieldName(name)
	info.fields = append(info.fields, FieldInfo{
		Name:      name,
		Source:    source,
		Size:      1,
		Precision: 0,
		Type:      BoolType,
	})
	info.numFields++
	info.fieldNames[name] = true
	return name
}

func (info *recordInfo) AddInt64Field(name string, source string) string {
	name = info.checkFieldName(name)
	info.fields = append(info.fields, FieldInfo{
		Name:      name,
		Source:    source,
		Size:      8,
		Precision: 0,
		Type:      Int64Type,
	})
	info.numFields++
	info.fieldNames[name] = true
	return name
}

func (info *recordInfo) AddStringField(name string, source string, size int) string {
	name = info.checkFieldName(name)
	info.fields = append(info.fields, FieldInfo{
		Name:      name,
		Source:    source,
		Size:      size,
		Precision: 0,
		Type:      StringType,
	})
	info.numFields++
	info.fieldNames[name] = true
	return name
}

func (info *recordInfo) AddV_WStringField(name string, source string, size int) string {
	name = info.checkFieldName(name)
	info.fields = append(info.fields, FieldInfo{
		Name:      name,
		Source:    source,
		Size:      size,
		Precision: 0,
		Type:      V_WStringType,
	})
	info.numFields++
	info.fieldNames[name] = true
	return name
}

func (info *recordInfo) checkFieldName(name string) string {
	_, exists := info.fieldNames[name]
	for exists {
		name = name + `2`
		exists = info.fieldNames[name]
	}
	return name
}
