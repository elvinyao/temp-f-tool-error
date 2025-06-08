package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// SensitiveParams defines commonly used sensitive parameter names
var SensitiveParams = []string{
	"token", "read_token", "write_token", "access_token", "api_key",
	"password", "secret", "key", "auth", "authorization", "credential",
}

// MaskedValue is the text used to replace sensitive values
const MaskedValue = "MASKED"
const MaskedValueLength = 3

// IsSensitiveParam checks if a parameter name contains sensitive information
func IsSensitiveParam(paramName string) bool {
	lowerName := strings.ToLower(paramName)

	for _, sensitive := range SensitiveParams {
		if strings.Contains(lowerName, sensitive) {
			return true
		}
	}

	return false
}

// SanitizeParamValue masks sensitive parameter values for logging
// It shows first 4 characters and masks the rest
func SanitizeParamValue(value interface{}) interface{} {
	if strVal, ok := value.(string); ok && len(strVal) > MaskedValueLength {
		return strVal[:MaskedValueLength] + MaskedValue
	}
	return value
}

// SanitizeParams processes a map of parameters and masks sensitive values
func SanitizeParams(params map[string]interface{}) map[string]interface{} {
	sanitizedParams := make(map[string]interface{})

	for k, v := range params {
		if IsSensitiveParam(k) {
			sanitizedParams[k] = SanitizeParamValue(v)
		} else {
			sanitizedParams[k] = v
		}
	}

	return sanitizedParams
}

// SanitizeErrorMessage cleans sensitive information from error messages
// Particularly useful for removing tokens from URLs in error messages
func SanitizeErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	errorMsg := err.Error()

	// Use regex to find and replace sensitive URL parameters
	for _, param := range SensitiveParams {
		// Pattern to match parameter=value in URLs
		pattern := regexp.MustCompile(fmt.Sprintf(`(?i)(%s=)[^&\s"]*`, regexp.QuoteMeta(param)))
		errorMsg = pattern.ReplaceAllString(errorMsg, "${1}"+MaskedValue)
	}

	// Additional cleanup for common URL patterns
	// Match URLs and clean query parameters
	urlPattern := regexp.MustCompile(`(https?://[^\s"]*\?)([^"\s]*)`)
	errorMsg = urlPattern.ReplaceAllStringFunc(errorMsg, func(match string) string {
		parts := strings.SplitN(match, "?", 2)
		if len(parts) != 2 {
			return match
		}

		baseURL := parts[0] + "?"
		queryString := parts[1]

		// Parse and clean query parameters
		if parsedQuery, err := url.ParseQuery(queryString); err == nil {
			cleanQuery := url.Values{}
			for key, values := range parsedQuery {
				if IsSensitiveParam(key) {
					cleanQuery[key] = []string{MaskedValue}
				} else {
					cleanQuery[key] = values
				}
			}
			return baseURL + cleanQuery.Encode()
		}

		return match
	})

	return errorMsg
}

// SanitizeString cleans sensitive information from any string
// This is a general-purpose function that can be used for various string sanitization needs
func SanitizeString(text string) string {
	// Use regex to find and replace sensitive patterns
	for _, param := range SensitiveParams {
		// Pattern to match parameter=value patterns
		pattern := regexp.MustCompile(fmt.Sprintf(`(?i)(%s[=:]\s*)[^\s&"']*`, regexp.QuoteMeta(param)))
		text = pattern.ReplaceAllString(text, "${1}"+MaskedValue)
	}

	return text
}
