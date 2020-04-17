# Makefile to help with builds
#

VERSION    := $(shell git describe --tags --dirty=-dirty-$$(date +"%s"))
BUILD_SHA  := $(shell git rev-parse HEAD)
GO_VERSION := $(shell go version | sed 's/go version //')
BIN_NAME   := sigcli
GH_OWNER   := timoguin
GH_REPO    := sigcli

# Use the above vars as build flags
LDFLAGS=-ldflags "-X \"main.version=${VERSION}\" -X \"main.buildSha=${BUILD_SHA}\" -X \"main.goVersion=${GO_VERSION}\""

default: help

clean: ## Clean local binaries and cache
	@echo "Cleaning local binaries and caches"
	@go clean
	@rm -rf dist

build: ## Build binary for the current platform
	@echo "Building binary for your current environment"
	go build ${LDFLAGS}

build-all: ## Build multi-platform binaries
	@echo "Tidying and vendoring go modules"
	go mod tidy
	go mod vendor
	@echo "Building multi-platform binaries using gox"
	gox ${LDFLAGS} -output "dist/${BIN_NAME}-${VERSION}-{{.OS}}-{{.Arch}}"

release: gen-checksums ## Publish release binaries to GitHub
	@echo "Publishing release binaries to GitHub"
	@if [ "$$(git diff --stat)" ]; then echo "Git tree is dirty! Aborting"; exit 1; fi
	ghr -soft ${VERSION} dist/

release-draft: gen-checksums ## Publish draft release binaries to GitHub
	@echo "Publishing draft release binaries to GitHub"
	ghr -replace -draft -name ${VERSION} ${VERSION} dist/

gen-checksums: ## Generate sha256 checksum file
	@echo "Generating sha256 sums for binaries"
	sha256sum dist/* > dist/${BIN_NAME}-${VERSION}.sha256sums

tools: ## Install build and release tools
	@echo "Installing build and release tools"
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr

debug: ## Show make debug output
	@echo "VERSION: ${VERSION}"
	@echo "BUILD_SHA: ${BUILD_SHA}"
	@echo "GO_VERSION: ${GO_VERSION}"
	@echo "BIN_NAME: ${BIN_NAME}"
	@echo "GH_OWNER: ${GH_OWNER}"
	@echo "GH_REPO: ${GH_REPO}"
	@echo "LDFLAGS: ${LDFLAGS}"

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
