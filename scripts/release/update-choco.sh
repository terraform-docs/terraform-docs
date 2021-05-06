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

if [ -n "$(git status --short)" ]; then
    echo "Error: There are untracked/modified changes, commit or discard them before the release."
    exit 1
fi

RELEASE_VERSION=$1

if [ -z "${RELEASE_VERSION}" ]; then
    echo "Error: release version is missing"
    exit 1
fi

PWD=$(cd "$(dirname "$0")" && pwd -P)

# get closest GA tag immediately before the latest one, ignore alpha, beta and rc tags
function getClosestVersion() {
    local latest
    latest=""
    for t in $(git tag --sort=-creatordate); do
        tag="$t"
        if [[ "$tag" == *"-alpha"* ]] || [[ "$tag" == *"-beta"* ]] || [[ "$tag" == *"-rc"* ]]; then
            continue
        fi
        if [ -z "$latest" ]; then
            latest="$t"
            continue
        fi
        break
    done
    echo "${tag//v/}"
}
CLOSEST_VERSION=$(getClosestVersion)

git clone https://github.com/terraform-docs/chocolatey-package "${PWD}/chocolatey-package"

# Bump version in terraform-docs.nuspec
sed -i -E "s|<version>${CLOSEST_VERSION}</version>|<version>${RELEASE_VERSION}</version>|g" "${PWD}/chocolatey-package/terraform-docs.nuspec"

# Bump version and checksum in tools/chocolateyinstall.ps1
CHECKSUM=$(grep windows-amd64.zip "${PWD}/../../dist/terraform-docs-v${RELEASE_VERSION}.sha256sum" | awk '{print $1}')

sed -i -E "s|checksum = '.*$|checksum = '${CHECKSUM}'|g" "${PWD}/chocolatey-package/tools/chocolateyinstall.ps1"
sed -i -E "s|v${CLOSEST_VERSION}|v${RELEASE_VERSION}|g" "${PWD}/chocolatey-package/tools/chocolateyinstall.ps1"

pushd "${PWD}/chocolatey-package/"
git diff
popd
