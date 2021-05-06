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
    CURRENT_VERSION=$(git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.1-$(COMMIT_HASH)")
fi

if [ "v${RELEASE_VERSION}" == "${CURRENT_VERSION}" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) already exists."
    exit 1
fi

if [ "$(git describe --tags "v${RELEASE_VERSION}" 2>/dev/null)" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) already exists."
    exit 1
fi

PWD=$(cd "$(dirname "$0")" && pwd -P)

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

# Bump the released version in README and version.go
if [[ $RELEASE_VERSION != *"-alpha"* && $RELEASE_VERSION != *"-beta"* && $RELEASE_VERSION != *"-rc"* ]]; then
    sed -i -E "s|${CLOSEST_VERSION}|${RELEASE_VERSION}|g" README.md
    sed -i -E "s|${CLOSEST_VERSION}|${RELEASE_VERSION}|g" docs/user-guide/installation.md
    git add README.md docs/user-guide/installation.md
fi

sed -i -E "s|v${RELEASE_VERSION}-alpha|v${RELEASE_VERSION}|g" internal/version/version.go
git add internal/version/version.go

# Commit changes
printf "\033[36m==> %s\033[0m\n" "Commit changes for release version v${RELEASE_VERSION}"
git commit -m "Release version v${RELEASE_VERSION}"

printf "\033[36m==> %s\033[0m\n" "Push commits for v${RELEASE_VERSION}"
git push origin master

# Tag the release
printf "\033[36m==> %s\033[0m\n" "Tag release v${RELEASE_VERSION}"
git tag --annotate --message "v${RELEASE_VERSION} Release" "v${RELEASE_VERSION}"

printf "\033[36m==> %s\033[0m\n" "Push tag release v${RELEASE_VERSION}"
git push origin "v${RELEASE_VERSION}"
