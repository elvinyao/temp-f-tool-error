package errors

import (
	"fmt"
	"net/http"
)

// NewConfigError creates a new AppError based on a definition from the configuration.
// It acts as the central factory for all configuration-based errors.
func NewConfigError(category, errorName string, cause error, args []interface{}, params map[string]interface{}) *AppError {
	errorDef, exists := FindErrorConfig(category, errorName)
	if !exists {
		// Fallback to a generic internal error if the definition is not found.
		// This prevents the application from crashing due to a missing error definition.
		return NewSystemError(
			ErrNameInternalError,
			fmt.Sprintf("undefined error: category='%s' name='%s'", category, errorName),
			cause,
		)
	}

	// Format the error message with provided arguments.
	message := errorDef.MessageTemplate
	if len(args) > 0 {
		message = fmt.Sprintf(errorDef.MessageTemplate, args...)
	}

	return &AppError{
		Type:       errorDef.Type,
		Code:       errorDef.Name,
		Message:    message,
		HTTPStatus: errorDef.HTTPStatus,
		Cause:      cause,
		Params:     params,
	}
}

func NewSystemError(code, message string, cause error, params ...map[string]interface{}) *AppError {
	var p map[string]interface{}
	if len(params) > 0 {
		p = params[0]
	}

	return &AppError{
		Type:       ErrorTypeSystem,
		Code:       code,
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
		Cause:      cause,
		Params:     p,
	}
}

func NewBusinessError(code, message string, cause error, params ...map[string]interface{}) *AppError {
	var p map[string]interface{}
	if len(params) > 0 {
		p = params[0]
	}

	return &AppError{
		Type:       ErrorTypeBusiness,
		Code:       code,
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
		Cause:      cause,
		Params:     p,
	}
}

// ConfigInvalidParam creates a 'parameter invalid' error using the configuration.
func ConfigInvalidParam(paramName string, reason string, cause error) *AppError {
	params := map[string]interface{}{
		ParamKeyParam:  paramName,
		ParamKeyReason: reason,
	}
	return NewConfigError(CategoryValidation, ErrNameInvalidParameter, cause, []interface{}{paramName, reason}, params)
}

// ConfigMissingParam creates a 'parameter missing' error using the configuration.
func ConfigMissingParam(paramName string) *AppError {
	params := map[string]interface{}{
		ParamKeyParam: paramName,
	}
	return NewConfigError(CategoryValidation, ErrNameMissingParameter, nil, []interface{}{paramName}, params)
}

// ConfigResourceNotFound creates a 'resource not found' error using the configuration.
func ConfigResourceNotFound(resourceID string, cause error) *AppError {
	params := map[string]interface{}{
		ParamKeyResourceID: resourceID,
	}
	return NewConfigError(CategoryResource, ErrNameNotFound, cause, []interface{}{resourceID}, params)
}

// ConfigUnauthorized creates an 'unauthorized' error using the configuration.
func ConfigUnauthorized(cause error) *AppError {
	return NewConfigError(CategoryAuthentication, ErrNameUnauthorized, cause, nil, nil)
}

// ConfigFocalboardAPIError creates a 'Focalboard API error' using the configuration.
func ConfigFocalboardAPIError(operation string, cause error) *AppError {
	params := map[string]interface{}{
		ParamKeyOperation: operation,
	}
	return NewConfigError(CategoryExternalService, ErrNameFocalboardAPIError, cause, []interface{}{operation}, params)
}

// ConfigForbidden creates a 'forbidden' error using the configuration.
func ConfigForbidden(cause error) *AppError {
	return NewConfigError(CategoryAuthentication, ErrNameForbidden, cause, nil, nil)
}
