GO := go

BINARY_NAME := $(shell basename $(CURDIR))

BUILD_DIR := build

TARGET_OS ?= linux
TARGET_ARCH ?= amd64

LDFLAGS := -ldflags="-w -s"

.PHONY: all build clean help

all: build

build:
	@echo "==> Building for $(TARGET_OS)/$(TARGET_ARCH)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "==> Build successful! Binary is at: $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "==> Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "==> Clean complete."

help:
	@echo "Available commands:"
	@echo "  make build    - Build the application for linux/amd64 (default)."
	@echo "  make clean    - Remove build artifacts."
	@echo "  make help     - Show this help message."
	@echo ""
	@echo "You can override the target OS and architecture:"
	@echo "  make build TARGET_OS=windows TARGET_ARCH=amd64"
	@echo "  make build TARGET_OS=darwin TARGET_ARCH=arm64  (for Apple Silicon)"

