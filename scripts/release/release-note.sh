#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PWD=$(cd $(dirname "$0") && pwd -P)

# Find closest tag
CLOSEST_VERSION=$(git describe --tags --abbrev=0)

# Install git-chglog binary
if ! command -v git-chglog >/dev/null ; then
    make git-chglog
fi

# Generate Changelog
git-chglog --config ${PWD}/../../scripts/chglog/config-release-note.yml --tag-filter-pattern v[0-9]+.[0-9]+.[0-9]+$ --output ${PWD}/../../CURRENT-RELEASE-CHANGELOG.md ${CLOSEST_VERSION}

cat ${PWD}/../../CURRENT-RELEASE-CHANGELOG.md
