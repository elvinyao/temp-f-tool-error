package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	HeaderXRequestID = "X-Request-ID"
)

// RequestID is the middleware to generate and add request id to context and response header
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// first try to get request id from header
		requestID := c.GetHeader(HeaderXRequestID)

		// if no request id in header, generate a new one
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// add to context and response header
		c.Set(HeaderXRequestID, requestID)
		c.Header(HeaderXRequestID, requestID)

		c.Next()
	}
}
