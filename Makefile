# Makefile for focalboard-tool project

# Variables
APP_NAME = focalboard-tool
GO_VERSION = 1.24.3
BUILD_DIR = ./build
DOCS_DIR = ./docs

# Colors for output
RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[0;33m
BLUE = \033[0;34m
NC = \033[0m # No Color

# Default target
.PHONY: help
help: ## Show this help message
	@echo "$(GREEN)Available commands:$(NC)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(BLUE)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development commands
.PHONY: dev
dev: ## Run the application in development mode
	@echo "$(YELLOW)Starting development server...$(NC)"
	go run main.go -c configs/application.toml

.PHONY: build
build: ## Build the application
	@echo "$(YELLOW)Building $(APP_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "$(GREEN)Build completed: $(BUILD_DIR)/$(APP_NAME)$(NC)"

.PHONY: clean
clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	rm -rf $(BUILD_DIR)
	@echo "$(GREEN)Clean completed$(NC)"

# Swagger/Documentation commands
.PHONY: swag-init
swag-init: ## Initialize swaggo (install swag if not exists)
	@echo "$(YELLOW)Checking swag installation...$(NC)"
	@which $$(go env GOPATH)/bin/swag > /dev/null || (echo "$(RED)Installing swag...$(NC)" && go install github.com/swaggo/swag/cmd/swag@latest)
	@echo "$(GREEN)Swag is ready$(NC)"

.PHONY: docs
docs: swag-init ## Generate swagger documentation
	@echo "$(YELLOW)Generating swagger documentation...$(NC)"
	$$(go env GOPATH)/bin/swag init -g main.go -o $(DOCS_DIR) --parseInternal --parseDependency
	@echo "$(GREEN)Documentation generated in $(DOCS_DIR)$(NC)"

.PHONY: docs-fmt
docs-fmt: swag-init ## Format swagger annotations
	@echo "$(YELLOW)Formatting swagger annotations...$(NC)"
	$$(go env GOPATH)/bin/swag fmt
	@echo "$(GREEN)Swagger annotations formatted$(NC)"

.PHONY: docs-clean
docs-clean: ## Clean generated documentation
	@echo "$(YELLOW)Cleaning documentation...$(NC)"
	rm -f $(DOCS_DIR)/docs.go $(DOCS_DIR)/swagger.json $(DOCS_DIR)/swagger.yaml
	@echo "$(GREEN)Documentation cleaned$(NC)"

.PHONY: docs-serve
docs-serve: docs ## Generate docs and start server to view them
	@echo "$(YELLOW)Starting server with swagger docs...$(NC)"
	@echo "$(GREEN)Swagger UI will be available at http://localhost:8080/swagger/index.html$(NC)"
	go run main.go -c configs/application.toml

# Testing commands
.PHONY: test
test: ## Run tests
	@echo "$(YELLOW)Running tests...$(NC)"
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "$(YELLOW)Running tests with coverage...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

# Dependency management
.PHONY: deps
deps: ## Download dependencies
	@echo "$(YELLOW)Downloading dependencies...$(NC)"
	go mod download
	@echo "$(GREEN)Dependencies downloaded$(NC)"

.PHONY: deps-tidy
deps-tidy: ## Clean up dependencies
	@echo "$(YELLOW)Tidying dependencies...$(NC)"
	go mod tidy
	@echo "$(GREEN)Dependencies tidied$(NC)"

.PHONY: deps-update
deps-update: ## Update dependencies
	@echo "$(YELLOW)Updating dependencies...$(NC)"
	go get -u ./...
	go mod tidy
	@echo "$(GREEN)Dependencies updated$(NC)"

# Code quality commands
.PHONY: fmt
fmt: ## Format code
	@echo "$(YELLOW)Formatting code...$(NC)"
	go fmt ./...
	@echo "$(GREEN)Code formatted$(NC)"

.PHONY: vet
vet: ## Run go vet
	@echo "$(YELLOW)Running go vet...$(NC)"
	go vet ./...
	@echo "$(GREEN)Vet completed$(NC)"

.PHONY: lint
lint: ## Run golangci-lint (requires golangci-lint to be installed)
	@echo "$(YELLOW)Running linter...$(NC)"
	@which golangci-lint > /dev/null || (echo "$(RED)golangci-lint not found. Install it from https://golangci-lint.run/$(NC)" && exit 1)
	golangci-lint run
	@echo "$(GREEN)Linting completed$(NC)"

# All-in-one commands
.PHONY: check
check: fmt vet docs ## Run all checks (format, vet, generate docs)
	@echo "$(GREEN)All checks completed$(NC)"

.PHONY: build-all
build-all: clean deps check test build ## Clean, get deps, check, test, and build
	@echo "$(GREEN)Full build completed$(NC)" 