#!/usr/bin/env bash
#
# Copyright 2021 The terraform-docs Authors.
#
# Licensed under the MIT license (the "License"); you may not
# use this file except in compliance with the License.
#
# You may obtain a copy of the License at the LICENSE file in
# the root directory of this source tree.

set -o errexit
set -o pipefail

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)

if [ -z "${CURRENT_BRANCH}" ] || [ "${CURRENT_BRANCH}" != "master" ]; then
    echo "Error: The current branch is '${CURRENT_BRANCH}', switch to 'master' to do the release."
    exit 1
fi

if [ -n "$(git status --short)" ]; then
    echo "Error: There are untracked/modified changes, commit or discard them before the release."
    exit 1
fi

RELEASE_VERSION=$1
CURRENT_VERSION=$2
FROM_MAKEFILE=$3

if [ -z "${RELEASE_VERSION}" ]; then
    if [ -z "${FROM_MAKEFILE}" ]; then
        echo "Error: VERSION is missing. e.g. ./release.sh <version>"
    else
        echo "Error: missing value for 'version'. e.g. 'make release VERSION=x.y.z'"
    fi
    exit 1
fi

if [ -z "${CURRENT_VERSION}" ]; then
    CURRENT_VERSION=$(git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.1-${COMMIT_HASH}")
fi

if [ "v${RELEASE_VERSION}" == "${CURRENT_VERSION}" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) already exists."
    exit 1
fi

if [ "$(git describe --tags "v${RELEASE_VERSION}" 2>/dev/null)" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) already exists."
    exit 1
fi

# get closest GA tag, ignore alpha, beta and rc tags
function getClosestVersion() {
    for t in $(git tag --sort=-creatordate); do
        tag="$t"
        if [[ "$tag" == *"-alpha"* ]] || [[ "$tag" == *"-beta"* ]] || [[ "$tag" == *"-rc"* ]]; then
            continue
        fi
        break
    done
    echo "${tag//v/}"
}
CLOSEST_VERSION=$(getClosestVersion)

echo "Release Version: ${RELEASE_VERSION}"
echo "Closest Version: ${CLOSEST_VERSION}"

# Bump the released version in README and version.go
if [[ $RELEASE_VERSION != *"-alpha"* && $RELEASE_VERSION != *"-beta"* && $RELEASE_VERSION != *"-rc"* ]]; then
    sed -i -E "s|${CLOSEST_VERSION}|${RELEASE_VERSION}|g" README.md
    sed -i -E "s|${CLOSEST_VERSION}|${RELEASE_VERSION}|g" docs/user-guide/installation.md

    echo "Modified: README.md"
    echo "Modified: docs/user-guide/installation.md"
    git add README.md docs/user-guide/installation.md
fi

sed -i -E "s|coreVersion([[:space:]]*)= \"(.*)\"|coreVersion\1= \"${RELEASE_VERSION}\"|g" internal/version/version.go
sed -i -E "s|prerelease([[:space:]]*)= \"(.*)\"|prerelease\1= \"\"|g" internal/version/version.go

echo "Modified: internal/version/version.go"
git add internal/version/version.go
