# Copyright 2021 The terraform-docs Authors.
#
# Licensed under the MIT license (the "License"); you may not
# use this file except in compliance with the License.
#
# You may obtain a copy of the License at the LICENSE file in
# the root directory of this source tree.

# Project variables
PROJECT_NAME  := terraform-docs
PROJECT_OWNER := terraform-docs
DESCRIPTION   := generate documentation from Terraform modules in various output formats
PROJECT_URL   := https://github.com/$(PROJECT_OWNER)/$(PROJECT_NAME)
LICENSE       := MIT

# Build variables
BUILD_DIR    := bin
COMMIT_HASH  ?= $(shell git rev-parse --short HEAD 2>/dev/null)
CUR_VERSION  ?= $(shell git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-$(COMMIT_HASH)")
COVERAGE_OUT := coverage.out

# Go variables
GO          ?= go
GO_PACKAGE  := github.com/$(PROJECT_OWNER)/$(PROJECT_NAME)
GOOS        ?= $(shell $(GO) env GOOS)
GOARCH      ?= $(shell $(GO) env GOARCH)

GOLDFLAGS   += -X $(GO_PACKAGE)/internal/version.commit=$(COMMIT_HASH)

GOBUILD     ?= CGO_ENABLED=0 $(GO) build -ldflags="$(GOLDFLAGS)"
GORUN       ?= GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run

GOIMPORTS_LOCAL_ARG := -local github.com/terraform-docs

# Docker variables
DEFAULT_TAG  ?= $(shell echo "$(CUR_VERSION)" | tr -d 'v')
DOCKER_IMAGE := quay.io/$(PROJECT_OWNER)/$(PROJECT_NAME)
DOCKER_TAG   ?= $(DEFAULT_TAG)

# Binary versions
GOLANGCI_VERSION  := v1.55.2

.PHONY: all
all: clean verify checkfmt lint test build

###############
##@ Development

.PHONY: checkfmt
checkfmt:   ## Check formatting of all go files
	@ $(MAKE) --no-print-directory log-$@
	@ goimports -l $(GOIMPORTS_LOCAL_ARG) main.go cmd/ internal/ scripts/docs/ && echo "OK"

.PHONY: clean
clean:   ## Clean workspace
	@ $(MAKE) --no-print-directory log-$@
	rm -rf ./$(BUILD_DIR) ./$(COVERAGE_OUT)

.PHONY: fmt
fmt:   ## Format all go files
	@ $(MAKE) --no-print-directory log-$@
	goimports -w $(GOIMPORTS_LOCAL_ARG) main.go cmd/ internal/ scripts/docs/

.PHONY: lint
lint:   ## Run linter
	@ $(MAKE) --no-print-directory log-$@
	golangci-lint run ./...

.PHONY: staticcheck
staticcheck:   ## Run staticcheck
	@ $(MAKE) --no-print-directory log-$@
	$(GO) run honnef.co/go/tools/cmd/staticcheck@2023.1.6 -- ./...

.PHONY: test
test:   ## Run tests
	@ $(MAKE) --no-print-directory log-$@
	$(GO) test -coverprofile=$(COVERAGE_OUT) -covermode=atomic -v ./...

.PHONY: verify
verify:   ## Verify 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	$(GO) mod verify

# removed and gitignoreed 'vendor/', not needed anymore #
.PHONY: vendor deps
vendor:
deps:

#########
##@ Build

.PHONY: build
build: clean ## Build binary for current OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	$(GOBUILD) -o ./$(BUILD_DIR)/$(GOOS)-$(GOARCH)/$(PROJECT_NAME)

.PHONY: docker
docker:   ## Build Docker image
	@ $(MAKE) --no-print-directory log-$@
	docker build --pull --tag $(DOCKER_IMAGE):$(DOCKER_TAG) --file Dockerfile .

.PHONY: push
push:   ## Push Docker image
	@ $(MAKE) --no-print-directory log-$@
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: docs
docs:   ## Generate document of formatter commands
	@ $(MAKE) --no-print-directory log-$@
	$(GORUN) ./scripts/docs/generate.go

###########
##@ Release

PATTERN =

# if the last relase was alpha, beta or rc, 'release' target has to used with current
# cycle release. For example if latest tag is v0.8.0-rc.2 and v0.8.0 GA needs to get
# released the following should be executed: "make release version=0.8.0"
.PHONY: release
release: VERSION ?= $(shell echo $(CUR_VERSION) | sed 's/^v//' | awk -F'[ .]' '{print $(PATTERN)}')
release:   ## Prepare release
	@ $(MAKE) --no-print-directory log-$@
	@ ./scripts/release/release.sh "$(VERSION)" "$(CUR_VERSION)" "1"

.PHONY: patch
patch: PATTERN = '\$$1\".\"\$$2\".\"\$$3+1'
patch: release   ## Prepare Patch release

.PHONY: minor
minor: PATTERN = '\$$1\".\"\$$2+1\".0\"'
minor: release   ## Prepare Minor release

.PHONY: major
major: PATTERN = '\$$1+1\".0.0\"'
major: release   ## Prepare Major release

###########
##@ Helpers

.PHONY: goimports
goimports:   ## Install goimports
ifeq (, $(shell which goimports))
	$(GO) install golang.org/x/tools/cmd/goimports@latest
endif

.PHONY: golangci
golangci:   ## Install golangci
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell $(GO) env GOPATH)/bin $(GOLANGCI_VERSION)
endif

.PHONY: tools
tools:   ## Install required tools
	@ $(MAKE) --no-print-directory log-$@
	@ $(MAKE) --no-print-directory goimports golangci

########################################################################
## Self-Documenting Makefile Help                                     ##
## https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html ##
########################################################################

########
##@ Help

.PHONY: help
help:   ## Display this help
	@awk \
		-v "col=\033[36m" -v "nocol=\033[0m" \
		' \
			BEGIN { \
				FS = ":.*##" ; \
				printf "Usage:\n  make %s<target>%s\n", col, nocol \
			} \
			/^[a-zA-Z_-]+:.*?##/ { \
				printf "  %s%-12s%s %s\n", col, $$1, nocol, $$2 \
			} \
			/^##@/ { \
				printf "\n%s%s%s\n", nocol, substr($$0, 5), nocol \
			} \
		' $(MAKEFILE_LIST)

log-%:
	@grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk \
			'BEGIN { \
				FS = ":.*?## " \
			}; \
			{ \
				printf "\033[36m==> %s\033[0m\n", $$2 \
			}'
