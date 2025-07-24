# Dify Golang SDK Makefile

.PHONY: build test test-coverage lint format clean install deps examples benchmark docs

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

# Run benchmarks
benchmark:
	$(GOTEST) -bench=. ./benchmark/
	$(GOTEST) -bench=. ./service/
	$(GOTEST) -bench=. ./web/

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
	rm -rf dist/

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

example-advanced:
	$(GOCMD) run examples/advanced/error_handling_example.go
	$(GOCMD) run examples/advanced/interface_design_example.go
	$(GOCMD) run examples/advanced/custom_errors_example.go

# Run all examples
examples: example-service example-web example-console example-workflow example-streaming example-advanced

# Generate documentation
docs:
	@echo "Generating documentation..."
	@mkdir -p docs/generated
	@$(GOCMD) run ./cmd/gendocs/main.go

# Development setup
dev-setup: deps
	chmod +x test_coverage.sh
	@echo "Development environment setup complete"

# Check Go modules
mod-check:
	$(GOMOD) verify

# Update dependencies
update-deps:
	$(GOMOD) get -u ./...
	$(GOMOD) tidy

# Security scan
security-scan:
	@echo "Running security scan..."
	gosec ./...

# Performance test
perf-test:
	@echo "Running performance tests..."
	$(GOTEST) -bench=Benchmark -benchmem ./benchmark/

# Code quality check
quality-check: lint test-coverage security-scan
	@echo "Code quality check complete"

# Release preparation
release-prep: clean deps test lint benchmark
	@echo "Release preparation complete"

# Docker build
docker-build:
	docker build -t dify-sdk-go .

# Docker test
docker-test:
	docker run --rm dify-sdk-go make test

# Generate changelog
changelog:
	@echo "Generating changelog..."
	@git log --oneline --since="1 month ago" > CHANGELOG.md

# Check for vulnerabilities
vuln-check:
	$(GOCMD) list -json -deps ./... | nancy sleuth

# Run all checks
all: quality-check benchmark examples
	@echo "All checks completed successfully"

# Help
help:
	@echo "Available targets:"
	@echo "  build         - Build the project"
	@echo "  deps          - Install dependencies"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  benchmark     - Run benchmarks"
	@echo "  lint          - Run linter"
	@echo "  format        - Format code"
	@echo "  clean         - Clean build artifacts"
	@echo "  install       - Install the package"
	@echo "  examples      - Run all examples"
	@echo "  example-*     - Run specific example"
	@echo "  docs          - Generate documentation"
	@echo "  dev-setup     - Setup development environment"
	@echo "  mod-check     - Verify Go modules"
	@echo "  update-deps   - Update dependencies"
	@echo "  security-scan - Run security scan"
	@echo "  perf-test     - Run performance tests"
	@echo "  quality-check - Run all quality checks"
	@echo "  release-prep  - Prepare for release"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-test   - Run tests in Docker"
	@echo "  changelog     - Generate changelog"
	@echo "  vuln-check    - Check for vulnerabilities"
	@echo "  all           - Run all checks"
	@echo "  help          - Show this help"