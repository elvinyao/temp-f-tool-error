package errors

import (
	"errors"
)

// IsErrorType checks if an error is of a specific AppError type
func IsErrorType(err error, errorType ErrorType, errorCode string) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Type == errorType && (errorCode == "" || appErr.Code == errorCode)
	}
	return false
}

// IsBusinessError checks if an error is a business error
func IsBusinessError(err error) bool {
	return IsErrorType(err, ErrorTypeBusiness, "")
}

// IsSystemError checks if an error is a system error
func IsSystemError(err error) bool {
	return IsErrorType(err, ErrorTypeSystem, "")
}

// IsNotFoundError checks if an error is a "resource not found" error
func IsNotFoundError(err error) bool {
	return IsErrorType(err, ErrorTypeBusiness, ErrNameNotFound)
}

// IsAuthError checks if an error is an authentication/authorization error
func IsAuthError(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Type == ErrorTypeBusiness &&
			(appErr.Code == ErrNameUnauthorized || appErr.Code == ErrNameForbidden)
	}
	return false
}

// IsInvalidParamError checks if an error is an invalid parameter error
func IsInvalidParamError(err error) bool {
	return IsErrorType(err, ErrorTypeBusiness, ErrNameInvalidParameter)
}

// IsMissingParamError checks if an error is a missing parameter error
func IsMissingParamError(err error) bool {
	return IsErrorType(err, ErrorTypeBusiness, ErrNameMissingParameter)
}
