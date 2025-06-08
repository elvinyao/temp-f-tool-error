package errors

import (
	"fmt"
)

type ErrorType string

const (
	ErrorTypeBusiness ErrorType = "business_error"
	ErrorTypeSystem   ErrorType = "system_error"
)

// AppError defines the application error structure
type AppError struct {
	Type       ErrorType              `json:"type"`             // error type: business_error or system_error
	Code       string                 `json:"code"`             // error code, e.g. "invalid_parameter"
	Message    string                 `json:"message"`          // user-friendly error message
	Detail     string                 `json:"detail,omitempty"` // detailed error information
	SystemID   string                 `json:"system_id"`        // system identifier
	HTTPStatus int                    `json:"-"`                // HTTP status code
	Cause      error                  `json:"-"`                // original error
	Params     map[string]interface{} `json:"params,omitempty"` // additional parameters
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Detail)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Cause
}
