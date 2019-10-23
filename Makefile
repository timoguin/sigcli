# Makefile to help with builds
#

VERSION     := 0.1.0
BUILD_SHA   := $(shell git rev-parse HEAD)
BUILD_TIME  := $(shell TZ=utc date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION  := $(shell go version | sed 's/go version //')
BIN_NAME    := sigcli

# Use the above vars as build flags
LDFLAGS=-ldflags "-X \"main.version=${VERSION}\" -X \"main.buildSha=${BUILD_SHA}\" -X \"main.buildTime=${BUILD_TIME}\" -X \"main.goVersion=${GO_VERSION}\""

help:
	@echo
	@echo "--------------------------------------------------------------"
	@echo
	@echo "Helper commands for building and deploying the project"
	@echo
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
	@echo "--------------------------------------------------------------"
	@echo

clean: ## Clean local binaries and cache
	@echo "Cleaning local binaries and caches"
	@go clean
	@rm -rf dist

build: ## Build binary for the current platform
	@echo "Building binary for your current environment"
	go build ${LDFLAGS}

build_all: ## Build multi-platform binaries
	@echo "Building multi-platform binaries using gox"
	gox ${LDFLAGS} -output "dist/${BIN_NAME}_{{.OS}}_{{.Arch}}"
