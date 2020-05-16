package recordinfo

import (
	"fmt"
	"github.com/tlarsen7572/Golang-Public/goalteryx/convert_strings"
	"strconv"
	"time"
	"unsafe"
)

func (info *recordInfo) GetByteValueFrom(fieldName string, record unsafe.Pointer) (value byte, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return byte(0), isNull, err
	}
	return *((*byte)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetBoolValueFrom(fieldName string, record unsafe.Pointer) (value bool, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return false, isNull, err
	}
	return *((*bool)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetInt16ValueFrom(fieldName string, record unsafe.Pointer) (value int16, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return 0, isNull, err
	}
	return *((*int16)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetInt32ValueFrom(fieldName string, record unsafe.Pointer) (value int32, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return 0, isNull, err
	}
	return *((*int32)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetInt64ValueFrom(fieldName string, record unsafe.Pointer) (value int64, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return 0, isNull, err
	}
	return *((*int64)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetFixedDecimalValueFrom(fieldName string, record unsafe.Pointer) (value float64, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return 0, isNull, err
	}
	valueStr := convert_strings.CToString(unsafe.Pointer(uintptr(record) + field.location))
	value, err = strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, false, fmt.Errorf(`error converting '%v' to double in '%v' field`, value, fieldName)
	}
	return value, false, nil
}

func (info *recordInfo) GetFloatValueFrom(fieldName string, record unsafe.Pointer) (value float32, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return 0, isNull, err
	}
	return *((*float32)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetDoubleValueFrom(fieldName string, record unsafe.Pointer) (value float64, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return 0, isNull, err
	}
	return *((*float64)(unsafe.Pointer(uintptr(record) + field.location))), false, nil
}

func (info *recordInfo) GetStringValueFrom(fieldName string, record unsafe.Pointer) (value string, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return ``, isNull, err
	}
	return convert_strings.CToString(unsafe.Pointer(uintptr(record) + field.location)), false, nil
}

func (info *recordInfo) GetWStringValueFrom(fieldName string, record unsafe.Pointer) (value string, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return ``, isNull, err
	}
	return convert_strings.WideCToString(unsafe.Pointer(uintptr(record) + field.location)), false, nil
}

func (info *recordInfo) GetDateValueFrom(fieldName string, record unsafe.Pointer) (value time.Time, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return zeroDate, isNull, err
	}
	dateStr := convert_strings.CToString(unsafe.Pointer(uintptr(record) + field.location))
	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return zeroDate, false, fmt.Errorf(`error converting date '%v' in GetDateValueFrom for field [%v], use format yyyy-MM-dd`, dateStr, fieldName)
	}
	return date, false, nil
}

func (info *recordInfo) GetDateTimeValueFrom(fieldName string, record unsafe.Pointer) (value time.Time, isNull bool, err error) {
	returnEarly, isNull, err, field := info.shouldReturnEarlyWith(fieldName, record)
	if returnEarly {
		return zeroDate, isNull, err
	}
	dateStr := convert_strings.CToString(unsafe.Pointer(uintptr(record) + field.location))
	date, err := time.Parse(dateTimeFormat, dateStr)
	if err != nil {
		return zeroDate, false, fmt.Errorf(`error converting datetime '%v' in GetDateValueFrom for field [%v], use format yyyy-MM-dd hh:mm:ss`, dateStr, fieldName)
	}
	return date, false, nil
}

func (info *recordInfo) shouldReturnEarlyWith(fieldName string, record unsafe.Pointer) (returnEarly bool, isNull bool, err error, field *fieldInfoEditor) {
	field, err = info.getFieldInfo(fieldName)
	if err != nil {
		return true, false, err, nil
	}
	if isValueNull(field, record) {
		return true, true, nil, field
	}
	return false, false, nil, field
}

var nullByteTypes = []string{
	ByteType,
	Int16Type,
	Int32Type,
	Int64Type,
	FixedDecimalType,
	FloatType,
	DoubleType,
	StringType,
	WStringType,
	DateType,
	DateTimeType,
}

func isValueNull(field *fieldInfoEditor, record unsafe.Pointer) bool {
	for _, nullByteType := range nullByteTypes {
		if nullByteType == field.Type {
			nullByte := *((*byte)(unsafe.Pointer(uintptr(record) + field.location + field.fixedLen)))
			return nullByte == byte(1)
		}
	}
	if field.Type == BoolType {
		nullByte := *((*byte)(unsafe.Pointer(uintptr(record) + field.location)))
		return nullByte == byte(2)
	}
	return false
}
