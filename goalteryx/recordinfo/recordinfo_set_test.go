package recordinfo_test

import (
	"github.com/tlarsen7572/Golang-Public/goalteryx/recordinfo"
	"testing"
)

func TestSetValuesAndGenerateRecord(t *testing.T) {
	recordInfo := generateTestRecordInfo()
	_ = recordInfo.SetByteField(`ByteField`, byte(1))
	_ = recordInfo.SetBoolField(`BoolField`, true)
	_ = recordInfo.SetInt16Field(`Int16Field`, 2)
	/*
		_ = recordInfo.SetInt32Field(`Int32Field`, 3)
		_ = recordInfo.SetInt64Field(`Int64Field`, 4)
		_ = recordInfo.SetFixedDecimalField(`FixedDecimalField`, 123.45)
		_ = recordInfo.SetFloatField(`FloatField`, float32(654.321))
		_ = recordInfo.SetDoubleField(`DoubleField`, 909.33)
		_ = recordInfo.SetStringField(`StringField`, `ABCDEFG`)
		_ = recordInfo.SetWStringField(`WStringField`, `CXVY`)
		_ = recordInfo.SetDateField(`DateField`, time.Date(2020,1,2,0 ,0,0,0,time.Local))
		_ = recordInfo.SetDateTimeField(`DateTimeField`, time.Date(2021, 3, 4, 5, 6, 7, 0, time.Local))
	*/
	record, err := recordInfo.GenerateRecord()
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}

	byteVal, isNull, err := recordInfo.GetByteValueFrom(`ByteField`, record)
	checkExpectedGetValueFrom(t, byteVal, byte(1), isNull, false, err, nil)

	boolVal, isNull, err := recordInfo.GetBoolValueFrom(`BoolField`, record)
	checkExpectedGetValueFrom(t, boolVal, true, isNull, false, err, nil)

	int16Val, isNull, err := recordInfo.GetInt16ValueFrom(`Int16Field`, record)
	checkExpectedGetValueFrom(t, int16Val, int16(2), isNull, false, err, nil)
}

func generateTestRecordInfo() recordinfo.RecordInfo {
	recordInfo := recordinfo.New()
	recordInfo.AddByteField(`ByteField`, ``)
	recordInfo.AddBoolField(`BoolField`, ``)
	recordInfo.AddInt16Field(`Int16Field`, ``)
	/*
		recordInfo.AddInt32Field(`Int32Field`, ``)
		recordInfo.AddInt64Field(`Int64Field`, ``)
		recordInfo.AddFixedDecimalField(`FixedDecimalField`, ``, 19, 2)
		recordInfo.AddFloatField(`FloatField`, ``)
		recordInfo.AddDoubleField(`DoubleField`, ``)
		recordInfo.AddStringField(`StringField`, ``, 10)
		recordInfo.AddWStringField(`WStringField`, ``, 5)
		recordInfo.AddDateField(`DateField`, ``)
		recordInfo.AddDateTimeField(`DateTimeField`, ``)
	*/
	return recordInfo
}
