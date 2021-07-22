#PROJECT_ROOT_DIR hold the absolute path to this project root dir.
PROJECT_ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: build
build:
	@go mod download && go mod verify
	@cd cmd/cart/api && go build
