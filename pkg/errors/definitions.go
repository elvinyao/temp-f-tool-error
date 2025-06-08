package errors

// =============================================================================
// Error Categories
// =============================================================================

// Constants for error categories, corresponding to `category_name` in errors.yaml.
const (
	CategoryValidation      = "validation"
	CategoryAuthentication  = "authentication"
	CategoryResource        = "resource"
	CategoryRateLimiting    = "rate_limiting"
	CategoryMethod          = "method"
	CategoryExternalService = "external_service"
	CategoryInternal        = "internal"
)

// =============================================================================
// Error Names
// =============================================================================

// Constants for specific error names, corresponding to `name` in errors.yaml.
const (
	// Validation Errors
	ErrNameInvalidParameter    = "invalid_parameter"
	ErrNameMissingParameter    = "missing_parameter"
	ErrNameParameterFormat     = "parameter_format_error"
	ErrNameParameterOutOfRange = "parameter_out_of_range"

	// Authentication Errors
	ErrNameUnauthorized = "unauthorized"
	ErrNameForbidden    = "forbidden"

	// Resource Errors
	ErrNameNotFound      = "not_found"
	ErrNameBoardNotFound = "board_not_found"
	ErrNameCardNotFound  = "card_not_found"

	// Rate Limiting Errors
	ErrNameRateLimitExceeded = "rate_limit_exceeded"
	ErrNameRequestTimeout    = "request_timeout"

	// Method Errors
	ErrNameMethodNotAllowed = "method_not_allowed"

	// External Service Errors
	ErrNameFocalboardAPIError = "focalboard_api_error"
	ErrNameMattermostAPIError = "mattermost_api_error"

	// Internal Errors
	ErrNameInternalError = "internal_error"
)

// =============================================================================
// Parameter Keys
// =============================================================================

// Constants for keys used in the AppError.Params map for structured logging.
const (
	ParamKeyParam      = "param"
	ParamKeyReason     = "reason"
	ParamKeyResourceID = "resource_id"
	ParamKeyOperation  = "operation"
)
