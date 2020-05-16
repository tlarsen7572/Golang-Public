package recordinfo

import "time"

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

var zeroDate = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

type generateBytes func(field *fieldInfoEditor) ([]byte, error)
