.PHONY: build test lint clean install fmt help

BINARY_NAME=ollama-gateway
BUILD_DIR=build
MAIN_PATH=cmd/gateway/main.go

VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

test:
	@echo "Running tests..."
	go test -v -race ./...

test-cover:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	@echo "Running linter..."
	golangci-lint run

fmt:
	@echo "Formatting code..."
	go fmt ./...

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/ || cp $(BUILD_DIR)/$(BINARY_NAME) $$HOME/go/bin/

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

dev:
	@echo "Running in development mode..."
	go run $(MAIN_PATH)

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

help:
	@echo "Available commands:"
	@echo "  build      - Build the binary"
	@echo "  test       - Run tests"
	@echo "  test-cover - Run tests with coverage"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  clean      - Clean build artifacts"
	@echo "  install    - Install binary to GOPATH/bin"
	@echo "  run        - Build and run the binary"
	@echo "  dev        - Run in development mode"
	@echo "  deps       - Download and tidy dependencies"
	@echo "  help       - Show this help message"