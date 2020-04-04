package traffic_cop

import (
	"errors"
	"fmt"
)

func _parseStringList(parameters map[string]interface{}, param string) ([]string, error) {
	valueList, listOk := parameters[param].([]interface{})
	valueStringList := []string{}
	valuesOk := true
	var valueString string
	for _, value := range valueList {
		valueString, valuesOk = value.(string)
		if !valuesOk {
			break
		}
		valueStringList = append(valueStringList, valueString)
	}
	if !listOk || !valuesOk {
		return nil, errors.New(fmt.Sprintf(`the %v parameter was not included or was not a list of strings`, param))
	}
	return valueStringList, nil
}

func _errorResponse(err error) FunctionResponse {
	return FunctionResponse{
		Err:      err,
		Response: nil,
	}
}

func _validResponse(value interface{}) FunctionResponse {
	return FunctionResponse{Err: nil, Response: value}
}

func _stringParamErr(param string) error {
	return errors.New(fmt.Sprintf(`the %v parameter was not included or was not a string`, param))
}
