# Copyright 2024 The terraform-docs Authors.
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
COMMIT_HASH  ?= $(shell git rev-parse --short HEAD 2>/dev/null)
CUR_VERSION  ?= $(shell git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-$(COMMIT_HASH)")

###########
##@ Release

.PHONY: contributors
contributors: OLD_VERSION ?= ""
contributors: NEW_VERSION ?= ""
contributors:  ## generate contributors list
	@ $(MAKE) --no-print-directory log-$@
	@ ./contributors.sh "$(OLD_VERSION)" "$(NEW_VERSION)" "1"

PATTERN =

# if the last relase was alpha, beta or rc, 'release' target has to used with current
# cycle release. For example if latest tag is v0.8.0-rc.2 and v0.8.0 GA needs to get
# released the following should be executed: "make release version=0.8.0"
.PHONY: release
release: VERSION ?= $(shell echo $(CUR_VERSION) | sed 's/^v//' | awk -F'[ .]' '{print $(PATTERN)}')
release:   ## Prepare release
	@ $(MAKE) --no-print-directory log-$@
	@ ./release.sh "$(VERSION)" "$(CUR_VERSION)" "1"

.PHONY: patch
patch: PATTERN = '\$$1\".\"\$$2\".\"\$$3+1'
patch: release   ## Prepare Patch release

.PHONY: minor
minor: PATTERN = '\$$1\".\"\$$2+1\".0\"'
minor: release   ## Prepare Minor release

.PHONY: major
major: PATTERN = '\$$1+1\".0.0\"'
major: release   ## Prepare Major release

########################################################################
## Self-Documenting Makefile Help                                     ##
## https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html ##
########################################################################

########
##@ Help

.PHONY: help
help:   ## Display this help
	@awk -v "col=\033[36m" -v "nocol=\033[0m" ' BEGIN { FS = ":.*##" ; printf "Usage:\n  make %s<target>%s\n", col, nocol } /^[a-zA-Z_-]+:.*?##/ { printf "  %s%-12s%s %s\n", col, $$1, nocol, $$2 } /^##@/ { printf "\n%s%s%s\n", nocol, substr($$0, 5), nocol } ' $(MAKEFILE_LIST)

log-%:
	@grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN { FS = ":.*?## " }; { printf "\033[36m==> %s\033[0m\n", $$2 }'
