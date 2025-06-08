package service

import (
	"focalboard-tool/pkg/errors"
)

// ValidateStringParam 验证字符串参数
func ValidateStringParam(name string, value interface{}) (string, error) {
	strValue, ok := value.(string)
	if !ok {
		return "", errors.ConfigInvalidParam(name, "must be a string", nil)
	}

	if strValue == "" {
		return "", errors.ConfigMissingParam(name)
	}

	return strValue, nil
}

// ValidateIntParam 验证整数参数
func ValidateIntParam(name string, value interface{}) (int, error) {
	intValue, ok := value.(int)
	if !ok {
		return 0, errors.ConfigInvalidParam(name, "must be an integer", nil)
	}

	return intValue, nil
}

// ValidateBoolParam 验证布尔参数
func ValidateBoolParam(name string, value interface{}) (bool, error) {
	boolValue, ok := value.(bool)
	if !ok {
		return false, errors.ConfigInvalidParam(name, "must be a boolean", nil)
	}

	return boolValue, nil
}
