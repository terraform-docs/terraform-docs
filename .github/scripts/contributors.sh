#!/usr/bin/env bash
#
# Copyright 2024 The terraform-docs Authors.
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

OLD_VERSION=${1//v/}
NEW_VERSION=${2//v/}
FROM_MAKEFILE=$3

# get closest GA tag, ignore alpha, beta and rc tags
function getClosestVersion() {
    for t in $(git tag --sort=-creatordate); do
        tag="$t"
        if [[ "$tag" == *"-alpha"* ]] || [[ "$tag" == *"-beta"* ]] || [[ "$tag" == *"-rc"* ]]; then
            continue
        fi
        if [ "$tag" == "v${NEW_VERSION}" ]; then
            continue
        fi
        break
    done
    echo "${tag//v/}"
}
CLOSEST_VERSION=$(getClosestVersion)

if [ -z "$OLD_VERSION" ]; then
    OLD_VERSION="${CLOSEST_VERSION}"
fi

if [ -z "$OLD_VERSION" ] || [ -z "$NEW_VERSION" ]; then
    if [ -z "${FROM_MAKEFILE}" ]; then
        echo "Error: refs are missing. e.g. contributors <OLD_VERSION> <NEW_VERSION>"
    else
        echo "Error: refs are missing. e.g. 'make contributors OLD_VERSION=x.y.z NEW_VERSION=a.b.c'"
    fi
    exit 1
fi

touch contributors.list

git log "v${OLD_VERSION}..v${NEW_VERSION}" |
grep ^Author: |
sed 's/ <.*//; s/^Author: //' |
sort |
uniq |
while read -r line; do
    name=$(printf %s "$line" | iconv -f utf-8 -t ascii//translit | jq -sRr @uri)
    handle=$(curl -fsSL "https://api.github.com/search/users?q=in:name%20${name}" | jq -r '.items[0].login')
    if [ "$handle" == "null" ]; then
        echo "- @${name}" >> contributors.list
    else
        echo "- @${handle}" >> contributors.list
    fi
    sleep 5
done
