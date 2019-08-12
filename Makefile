# Project variables
NAME        := terraform-docs
VENDOR      := segmentio
DESCRIPTION := Generate docs from Terraform modules
MAINTAINER  := Martin Etmajer <metmajer@getcloudnative.io>
URL         := https://github.com/$(VENDOR)/$(NAME)
LICENSE     := MIT

# Repository variables
PACKAGE     := github.com/$(VENDOR)/$(NAME)

# Build variables
BUILD_DIR   := bin
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE  ?= $(shell date +%FT%T%z)
VERSION     ?= $(shell git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-$(COMMIT_HASH)")

# Go variables
GOCMD       := GO111MODULE=on go
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
GOFILES     ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOPKGS      ?= $(shell $(GOCMD) list $(MODVENDOR) ./... | grep -v /vendor)
MODVENDOR   := -mod=vendor

GOLDFLAGS   :="
GOLDFLAGS   += -X $(PACKAGE)/internal/pkg/version.version=$(VERSION)
GOLDFLAGS   += -X $(PACKAGE)/internal/pkg/version.commitHash=$(COMMIT_HASH)
GOLDFLAGS   += -X $(PACKAGE)/internal/pkg/version.buildDate=$(BUILD_DATE)
GOLDFLAGS   +="

GOBUILD     ?= CGO_ENABLED=0 $(GOCMD) build $(MODVENDOR) -ldflags $(GOLDFLAGS)
GORUN       ?= GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) run $(MODVENDOR)

# Binary versions
GITCHGLOG_VERSION := 0.8.0
GOLANGCI_VERSION  := v1.17.1

.PHONY: all
all: clean deps lint test build

#########################
## Development targets ##
#########################
.PHONY: checkfmt
checkfmt: RESULT = $(shell goimports -l $(GOFILES) | tee >(if [ "$$(wc -l)" = 0 ]; then echo "OK"; fi))
checkfmt: SHELL := /usr/bin/env bash
checkfmt: ## Check formatting of all go files
	@ $(MAKE) --no-print-directory log-$@
	@ echo "$(RESULT)"
	@ if [ "$(RESULT)" != "OK" ]; then exit 1; fi

.PHONY: clean
clean: ## Clean workspace
	@ $(MAKE) --no-print-directory log-$@
	rm -rf ./$(BUILD_DIR)

.PHONY: fmt
fmt: ## Format all go files
	@ $(MAKE) --no-print-directory log-$@
	goimports -w $(GOFILES)

.PHONY: lint
lint: ## Run linter
	@ $(MAKE) --no-print-directory log-$@
	GO111MODULE=on golangci-lint run ./...

.PHONY: test
test: ## Run tests
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) test $(MODVENDOR) -v $(GOPKGS)

.PHONY: vendor
vendor: ## Install 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) mod vendor

.PHONY: verify
verify: ## Verify 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) mod verify

###################
## Build targets ##
###################
.PHONY: build
build: clean ## Build binary for current OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	$(GOBUILD) -o ./$(BUILD_DIR)/$(GOOS)-$(GOARCH)/$(NAME)

.PHONY: build-all
build-all: GOOS      = linux darwin windows freebsd
build-all: GOARCH    = amd64 arm
build-all: clean ## Build binary for all OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	@ ./hack/build/build-all-osarch.sh "$(BUILD_DIR)" "$(NAME)" "$(VERSION)" "$(GOOS)" "$(GOARCH)" $(GOLDFLAGS)

#####################
## Release targets ##
#####################
.PHONY: release patch minor major
PATTERN =

release: version ?= $(shell echo $(VERSION) | sed 's/^v//' | awk -F'[ .]' '{print $(PATTERN)}')
release: ## Prepare release
	@ $(MAKE) --no-print-directory log-$@
	@ ./hack/release/release.sh "$(version)" "$(VERSION)"

patch: PATTERN = '\$$1\".\"\$$2\".\"\$$3+1'
patch: release ## Prepare Patch release

minor: PATTERN = '\$$1\".\"\$$2+1\".0\"'
minor: release ## Prepare Minor release

major: PATTERN = '\$$1+1\".0.0\"'
major: release ## Prepare Major release

####################
## Helper targets ##
####################
.PHONY: authors
authors: ## Generate Authors
	git log --all --format='%aN <%aE>' | sort -u | egrep -v noreply > AUTHORS

.PHONY: changelog
changelog: ## Generate Changelog
	git-chglog -o CHANGELOG.md

.PHONY: goimports
goimports: ## Install goimports
	@ $(MAKE) --no-print-directory log-$@
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

.PHONY: tools
tools: ## Install required tools
	@ $(MAKE) --no-print-directory log-$@
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s  -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)
	curl -sfL https://github.com/git-chglog/git-chglog/releases/download/$(GITCHGLOG_VERSION)/git-chglog_$(shell go env GOOS)_$(shell go env GOARCH) -o $(shell go env GOPATH)/bin/git-chglog && chmod +x $(shell go env GOPATH)/bin/git-chglog
	GO111MODULE=off go get -u github.com/mitchellh/gox

####################################
## Self-Documenting Makefile Help ##
####################################
.PHONY: help
help:
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

log-%:
	@ grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m==> %s\033[0m\n", $$2}'
