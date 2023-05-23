BUILD_DIR ?= ./build
VERSION := $(shell git rev-parse --short HEAD)
SERVICE_NAME = notifier

# env
include .env
export $(shell sed 's/=.*//' .env)

# clean
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

# build
.PHONY: build
build: clean
	@echo "Building version ${VERSION}"
	@mkdir -p $(BUILD_DIR)
	@go build \
		-ldflags="-X main.Version=${VERSION}" \
		-o "${BUILD_DIR}/${SERVICE_NAME}" \
		apps/main/main.go
	@echo "Build placed -> ${BUILD_DIR}/${SERVICE_NAME}"

# run
.PHONY: run
run:
	@go run \
		-ldflags="-X main.Version=${VERSION}" \
		apps/main/main.go