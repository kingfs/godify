# Dify Golang SDK Makefile

.PHONY: build test test-coverage lint format clean install deps examples

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Build
build:
	$(GOBUILD) -v ./...

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	./test_coverage.sh

# Lint code
lint:
	$(GOLINT) run ./...

# Format code
format:
	$(GOFMT) -w .

# Clean
clean:
	$(GOCLEAN)
	rm -rf coverage/

# Install
install: deps
	$(GOCMD) install

# Run examples
example-service:
	$(GOCMD) run examples/service/service_api_example.go

example-web:
	$(GOCMD) run examples/web/web_api_example.go

example-console:
	$(GOCMD) run examples/console/console_api_example.go

example-workflow:
	$(GOCMD) run examples/workflow/complete_workflow_example.go

example-streaming:
	$(GOCMD) run examples/stream/streaming_example.go

# Run all examples
examples: example-service example-web example-console example-workflow example-streaming

# Development setup
dev-setup: deps
	chmod +x test_coverage.sh

# Check Go modules
mod-check:
	$(GOMOD) verify

# Update dependencies
update-deps:
	$(GOMOD) get -u ./...
	$(GOMOD) tidy

# Release preparation
release-prep: clean deps test lint
	@echo "Release preparation complete"

# Help
help:
	@echo "Available targets:"
	@echo "  build         - Build the project"
	@echo "  deps          - Install dependencies"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  lint          - Run linter"
	@echo "  format        - Format code"
	@echo "  clean         - Clean build artifacts"
	@echo "  install       - Install the package"
	@echo "  examples      - Run all examples"
	@echo "  example-*     - Run specific example"
	@echo "  dev-setup     - Setup development environment"
	@echo "  mod-check     - Verify Go modules"
	@echo "  update-deps   - Update dependencies"
	@echo "  release-prep  - Prepare for release"
	@echo "  help          - Show this help"