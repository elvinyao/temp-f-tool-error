package service

import (
	"focalboard-tool/pkg/errors"
)

// ConfigValidateStringParam 使用配置的错误类型验证字符串参数
func ConfigValidateStringParam(name string, value interface{}) (string, error) {
	strValue, ok := value.(string)
	if !ok {
		return "", errors.ConfigInvalidParam(name, "必须是字符串类型", nil)
	}

	if strValue == "" {
		return "", errors.ConfigMissingParam(name)
	}

	return strValue, nil
}

// ConfigValidateIntParam 使用配置的错误类型验证整数参数
func ConfigValidateIntParam(name string, value interface{}) (int, error) {
	intValue, ok := value.(int)
	if !ok {
		return 0, errors.ConfigInvalidParam(name, "必须是整数类型", nil)
	}

	return intValue, nil
}

// ConfigValidateBoolParam 使用配置的错误类型验证布尔参数
func ConfigValidateBoolParam(name string, value interface{}) (bool, error) {
	boolValue, ok := value.(bool)
	if !ok {
		return false, errors.ConfigInvalidParam(name, "必须是布尔类型", nil)
	}

	return boolValue, nil
}
