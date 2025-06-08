package middleware

import (
	"focalboard-tool/internal/conf"
	"focalboard-tool/library/log"
	"focalboard-tool/pkg/errors"
	"focalboard-tool/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorResponse is the unified error response format
type ErrorResponse struct {
	Success bool             `json:"success"`
	Error   *errors.AppError `json:"error,omitempty"`
	Data    interface{}      `json:"data,omitempty"`
}

// SuccessResponse is the unified success response format
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorHandler is the middleware to handle errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// call next handler
		c.Next()

		// if there is an error, handle it
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			handleError(c, err)
		}
	}
}

// handle different types of errors
func handleError(c *gin.Context, err error) {
	// get request id
	requestID, exists := c.Get(HeaderXRequestID)
	if !exists {
		requestID = "unknown"
	}

	// get request related information, for enhanced logging
	method := c.Request.Method
	path := c.Request.URL.Path
	clientIP := c.ClientIP()
	userAgent := c.Request.UserAgent()

	// default is 500 internal error
	httpStatus := http.StatusInternalServerError
	var appErr *errors.AppError

	// try to convert error to AppError
	if e, ok := err.(*errors.AppError); ok {
		appErr = e
		httpStatus = e.HTTPStatus
	} else {
		// unknown error to system error
		appErr = errors.NewSystemError(
			errors.ErrNameInternalError,
			"Unknown system error",
			err,
		)
	}

	logFields := []zap.Field{
		zap.String("request_id", requestID.(string)),
		zap.String("method", method),
		zap.String("path", path),
		zap.String("client_ip", clientIP),
		zap.String("user_agent", userAgent),
		zap.String("error_code", appErr.Code),
		zap.String("error_message", appErr.Message),
		zap.Int("http_status", httpStatus),
	}

	if appErr.Type == errors.ErrorTypeSystem {
		// Clean sensitive information from error details for logging
		sanitizedError := utils.SanitizeErrorMessage(appErr.Cause)

		// add system error specific fields
		systemLogFields := append(logFields,
			zap.String("error_detail", appErr.Detail),
			zap.String("error", sanitizedError),
			zap.String("error_type", "system_error"),
		)

		log.Error("system error occurred", systemLogFields...)
	} else {
		// business error only record as information, add business error identifier
		businessLogFields := append(logFields,
			zap.String("error_type", "business_error"),
		)

		log.Info("business error occurred", businessLogFields...)
	}

	// hide system error details, avoid leaking sensitive information
	if appErr.Type == errors.ErrorTypeSystem {
		// in production environment, hide detailed error information from clients
		responseErr := *appErr
		responseErr.Detail = ""                      // clear detailed error information
		responseErr.SystemID = conf.Conf.App.AppName // set system ID

		// add request id for technical support
		if responseErr.Params == nil {
			responseErr.Params = make(map[string]interface{})
		}
		responseErr.Params["request_id"] = requestID
		responseErr.Params["timestamp"] = time.Now().Unix()

		c.JSON(httpStatus, ErrorResponse{
			Success: false,
			Error:   &responseErr,
		})
	} else {
		// business error can return complete information
		appErr.SystemID = conf.Conf.App.AppName // set system ID
		if appErr.Params == nil {
			appErr.Params = make(map[string]interface{})
		}
		appErr.Params["request_id"] = requestID
		appErr.Params["timestamp"] = time.Now().Unix()

		c.JSON(httpStatus, ErrorResponse{
			Success: false,
			Error:   appErr,
		})
	}
}

// RespondSuccess is the helper function for success response, can be used in controllers
func RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}
