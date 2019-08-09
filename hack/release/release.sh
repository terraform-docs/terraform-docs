#!/usr/bin/env bash

set -e

RELEASE_VERSION="$1"
CURRENT_VERSION="$2"

if [ -z "${RELEASE_VERSION}" ]; then
    echo "Error: missing value for 'version'. e.g. 'make release version=x.y.z'"
    exit 1
fi

if [ -z "${CURRENT_VERSION}" ]; then
    echo "Error: CURRENT_VERSION is missing. e.g. ./release.sh <release_version> <current_version>"
    exit 1
fi

if [ "v${RELEASE_VERSION}" = "${CURRENT_VERSION}" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) exists."
    exit 1
else
    git tag --annotate --message "v${RELEASE_VERSION} Release" v${RELEASE_VERSION}
    echo "Tag v${RELEASE_VERSION} Release"
    git push origin v${RELEASE_VERSION}
    echo "Push v${RELEASE_VERSION} Release"
fi
