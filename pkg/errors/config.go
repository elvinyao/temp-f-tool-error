package errors

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// ErrorConfig represents the top-level structure of the error configuration file.
type ErrorConfig struct {
	BusinessErrors []ErrorCategory `yaml:"business_error"`
	SystemErrors   []ErrorCategory `yaml:"system_error"`
}

// ErrorCategory represents a category of errors, such as "validation" or "authentication".
type ErrorCategory struct {
	CategoryName string        `yaml:"category_name"`
	Description  string        `yaml:"description"`
	Errors       []ErrorDetail `yaml:"errors"`
}

// ErrorDetail represents the specific configuration for a single error.
type ErrorDetail struct {
	Name            string `yaml:"name"`
	HTTPStatus      int    `yaml:"http_status"`
	MessageTemplate string `yaml:"message_template"`
}

// ErrorDefinition is a flattened structure that combines details from ErrorDetail
// with its type (business or system) for easier use in the application.
type ErrorDefinition struct {
	Name            string
	HTTPStatus      int
	MessageTemplate string
	Type            ErrorType
}

var (
	// Global error configuration instance.
	errorConfig     *ErrorConfig
	errorConfigOnce sync.Once
	errorConfigLock sync.RWMutex
)

// LoadErrorConfig loads and parses the error configuration from a YAML file.
func LoadErrorConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read error config file '%s': %w", configPath, err)
	}

	var config ErrorConfig
	fileExt := strings.ToLower(filepath.Ext(configPath))

	if fileExt != ".yaml" && fileExt != ".yml" {
		return fmt.Errorf("unsupported config file format: %s. Only .yaml or .yml is supported", fileExt)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse error config file: %w", err)
	}

	errorConfigLock.Lock()
	defer errorConfigLock.Unlock()
	errorConfig = &config

	return nil
}

// GetErrorConfig returns the current error configuration.
// If no configuration has been loaded, it throws an error.
func GetErrorConfig() *ErrorConfig {
	errorConfigOnce.Do(func() {
		if errorConfig == nil {
			// If config is not loaded, throw an error.
			panic("error config not loaded")
		}
	})

	errorConfigLock.RLock()
	defer errorConfigLock.RUnlock()
	return errorConfig
}

// FindErrorConfig searches for a specific error definition by its category and name.
// It searches through both business and system errors and returns a flattened, easy-to-use definition.
func FindErrorConfig(categoryName, errorName string) (ErrorDefinition, bool) {
	config := GetErrorConfig()

	// Search in business errors
	for _, category := range config.BusinessErrors {
		if category.CategoryName == categoryName {
			for _, errDetail := range category.Errors {
				if errDetail.Name == errorName {
					return ErrorDefinition{
						Name:            errDetail.Name,
						HTTPStatus:      errDetail.HTTPStatus,
						MessageTemplate: errDetail.MessageTemplate,
						Type:            ErrorTypeBusiness,
					}, true
				}
			}
			// Assuming category names are unique, if the category is found but the error isn't, we can stop.
			return ErrorDefinition{}, false
		}
	}

	// Search in system errors
	for _, category := range config.SystemErrors {
		if category.CategoryName == categoryName {
			for _, errDetail := range category.Errors {
				if errDetail.Name == errorName {
					return ErrorDefinition{
						Name:            errDetail.Name,
						HTTPStatus:      errDetail.HTTPStatus,
						MessageTemplate: errDetail.MessageTemplate,
						Type:            ErrorTypeSystem,
					}, true
				}
			}
			// Assuming category names are unique, if the category is found but the error isn't, we can stop.
			return ErrorDefinition{}, false
		}
	}

	return ErrorDefinition{}, false
}
