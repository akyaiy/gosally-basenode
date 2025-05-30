APP_NAME := node
BIN_DIR := bin
GOPATH := $(shell go env GOPATH)
.PHONY: all build run test fmt vet lint check clean

all: build

lint-setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

goimports-setup:
	go install golang.org/x/tools/cmd/goimports@latest

setup: lint-setup goimports-setup
	@mkdir -p $(BIN_DIR)
	@echo "Setup complete. Run 'make build' to compile the application."

build:
	@go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

run: build
	./$(BIN_DIR)/$(APP_NAME)

test:
	@go test ./...

fmt:
	@go fmt ./...
	@$(GOPATH)/bin/goimports -w .

vet:
	@go vet ./...

lint:
	@$(GOPATH)/bin/golangci-lint run

check: fmt vet lint test

clean:
	@rm -rf bin