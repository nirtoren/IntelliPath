# Makefile for IntelliPath project

# Go source files
SRC := $(shell find . -name '*.go')

# Binary name
BINARY := intellipath

# Compressed archive name
TARFILE := intellipath.tar.gz

# Default target
.PHONY: all
all: build

# Build the Go project
.PHONY: build
build: $(BINARY)

$(BINARY): $(SRC)
	go build -o $(BINARY) .

# Run tests
.PHONY: test
test:
	go test ./...

# Run tests, if they pass, build the binary and create a tar.gz file
.PHONY: release
release: test build
	tar -czf $(TARFILE) $(BINARY)

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY) $(TARFILE)

# Install dependencies (if any)
.PHONY: deps
deps:
	go mod tidy
