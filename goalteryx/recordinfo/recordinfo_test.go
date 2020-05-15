package recordinfo_test

import (
	"github.com/tlarsen7572/Golang-Public/goalteryx/recordinfo"
	"testing"
	"unsafe"
)

var recordInfoXml = `<MetaInfo connection="Output">
	<RecordInfo>
		<Field name="ByteField" source="TextInput:" type="Byte"/>
		<Field name="BoolField" source="Formula: 1" type="Bool"/>
		<Field name="Int16Field" source="Formula: 16" type="Int16"/>
		<Field name="Int32Field" source="Formula: 32" type="Int32"/>
		<Field name="Int64Field" source="Formula: 64" type="Int64"/>
		<Field name="FixedDecimalField" scale="6" size="19" source="Formula: 123.45" type="FixedDecimal"/>
		<Field name="FloatField" source="Formula: 678.9" type="Float"/>
		<Field name="DoubleField" source="Formula: 0.12345" type="Double"/>
		<Field name="StringField" size="64" source="Formula: &quot;A&quot;" type="String"/>
		<Field name="WStringField" size="64" source="Formula: &quot;AB&quot;" type="WString"/>
		<Field name="V_StringShortField" size="1000" source="Formula: &quot;ABC&quot;" type="V_String"/>
		<Field name="V_StringLongField" size="2147483647" source="Formula: PadLeft(&quot;&quot;, 500, &apos;B&apos;)" type="V_String"/>
		<Field name="V_WStringShortField" size="10" source="Formula: &quot;XZY&quot;" type="V_WString"/>
		<Field name="V_WStringLongField" size="1073741823" source="Formula: PadLeft(&quot;&quot;, 500, &apos;W&apos;)" type="V_WString"/>
		<Field name="DateField" source="Formula: &apos;2020-01-01&apos;" type="Date"/>
		<Field name="DateTimeField" source="Formula: &apos;2020-02-03 04:05:06&apos;" type="DateTime"/>
	</RecordInfo>
</MetaInfo>
`

var sampleRecord = unsafe.Pointer(&[]byte{1, 0, 1, 16, 0, 0, 32, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 49, 50, 51, 46, 52, 53, 48, 48, 48, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 154, 185, 41, 68, 0, 124, 242, 176, 80, 107, 154, 191, 63, 0, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 0, 66, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 66, 67, 48, 47, 0, 0, 0, 35, 2, 0, 0, 38, 2, 0, 0, 50, 48, 50, 48, 45, 48, 49, 45, 48, 49, 0, 50, 48, 50, 48, 45, 48, 50, 45, 48, 51, 32, 48, 52, 58, 48, 53, 58, 48, 54, 0, 235, 5, 0, 0, 232, 3, 0, 0, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 13, 88, 0, 90, 0, 89, 0, 208, 7, 0, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0, 87, 0}[0])

func TestInstantiateRecordInfoFromXml(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if count := recordInfo.NumFields(); count != 16 {
		t.Fatalf(`expecpted 16 fields but got %v`, count)
	}
	type expectedStruct struct {
		Name      string
		Size      int
		Scale     int
		FieldType string
	}
	expectedFields := []expectedStruct{
		{`ByteField`, 1, 0, recordinfo.ByteType},
		{`BoolField`, 1, 0, recordinfo.BoolType},
		{`Int16Field`, 2, 0, recordinfo.Int16Type},
		{`Int32Field`, 4, 0, recordinfo.Int32Type},
		{`Int64Field`, 8, 0, recordinfo.Int64Type},
		{`FixedDecimalField`, 19, 6, recordinfo.FixedDecimalType},
		{`FloatField`, 4, 0, recordinfo.FloatType},
		{`DoubleField`, 8, 0, recordinfo.DoubleType},
		{`StringField`, 64, 0, recordinfo.StringType},
		{`WStringField`, 64, 0, recordinfo.WStringType},
		{`V_StringShortField`, 1000, 0, recordinfo.V_StringType},
		{`V_StringLongField`, 2147483647, 0, recordinfo.V_StringType},
		{`V_WStringShortField`, 10, 0, recordinfo.V_WStringType},
		{`V_WStringLongField`, 1073741823, 0, recordinfo.V_WStringType},
		{`DateField`, 10, 0, recordinfo.DateType},
		{`DateTimeField`, 19, 0, recordinfo.DateTimeType},
	}
	for index, expectedField := range expectedFields {
		field, _ := recordInfo.GetFieldByIndex(index)
		if field.Name != expectedField.Name {
			t.Fatalf(`expected name '%v' but got '%v' at field %v`, expectedField.Name, field.Name, index)
		}
		if field.Size != expectedField.Size {
			t.Fatalf(`expected size %v but got %v at field %v`, expectedField.Size, field.Size, index)
		}
		if field.Precision != expectedField.Scale {
			t.Fatalf(`expected scale %v but got %v at field %v`, expectedField.Scale, field.Precision, index)
		}
		if field.Type != expectedField.FieldType {
			t.Fatalf(`expected '%v' but got '%v' at field %v`, expectedField.FieldType, field.Type, index)
		}
	}
}

func TestCorrectlyRetrieveByteValue(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetByteValueFrom(`ByteField`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != byte(1) {
		t.Fatalf(`expected 1 but got %v`, value)
	}
}

func TestCorrectlyRetrieveBoolValue(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetBoolValueFrom(`ByteField`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != true {
		t.Fatalf(`expected true but got %v`, value)
	}
}

func TestCorrectlyRetrieveInt16Value(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetInt16ValueFrom(`Int16Field`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != 16 {
		t.Fatalf(`expected 16 but got %v`, value)
	}
}

func TestCorrectlyRetrieveInt32Value(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetInt32ValueFrom(`Int32Field`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != 32 {
		t.Fatalf(`expected 32 but got %v`, value)
	}
}

func TestCorrectlyRetrieveInt64Value(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetInt64ValueFrom(`Int64Field`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != 64 {
		t.Fatalf(`expected 64 but got %v`, value)
	}
}

func TestCorrectlyRetrieveFixedDecimalValue(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetFixedDecimalValueFrom(`FixedDecimalField`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != 123.450000 {
		t.Fatalf(`expected 123.450000 but got %v`, value)
	}
}

func TestCorrectlyRetrieveFloatValue(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetFloatValueFrom(`FloatField`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != 678.9 {
		t.Fatalf(`expected 678.9 but got %v`, value)
	}
}

func TestCorrectlyRetrieveDoubleValue(t *testing.T) {
	recordInfo, err := recordinfo.FromXml(recordInfoXml)
	if err != nil {
		t.Fatalf(err.Error())
	}

	value, err := recordInfo.GetDoubleValueFrom(`DoubleField`, sampleRecord)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if value != 0.12345 {
		t.Fatalf(`expected 0.12345 but got %v`, value)
	}
}
