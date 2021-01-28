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
set -o nounset
set -o pipefail

PWD=$(cd "$(dirname "$0")" && pwd -P)

# Find closest tag
CLOSEST_VERSION=$(git describe --tags --abbrev=0)

# Install git-chglog binary
if ! command -v git-chglog >/dev/null ; then
    make git-chglog
fi

# Generate Changelog
git-chglog --config "${PWD}"/../../scripts/chglog/config-release-note.yml --tag-filter-pattern v[0-9]+.[0-9]+.[0-9]+$ --output "${PWD}"/../../CURRENT-RELEASE-CHANGELOG.md "${CLOSEST_VERSION}"

cat "${PWD}"/../../CURRENT-RELEASE-CHANGELOG.md
