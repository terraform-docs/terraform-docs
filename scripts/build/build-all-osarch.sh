#!/usr/bin/env bash

set -e

BUILD_DIR="${1:-bin}"
NAME="${2:-terraform-docs}"
VERSION="$3"
GOOS="${4:-"linux darwin windows freebsd"}"
GOARCH="${5:-"amd64 arm"}"
GOLDFLAGS="$6"

if [ -z "${VERSION}" ]; then
    echo "Error: VERSION is missing. e.g. ./compress.sh <build_dir> <name> <version> <build_os_list> <build_arch_list> <build_ldflag>"
    exit 1
fi
if [ -z "${GOLDFLAGS}" ]; then
    echo "Error: GOLDFLAGS is missing. e.g. ./compress.sh <build_dir> <name> <version> <build_os_list> <build_arch_list> <build_ldflag>"
    exit 1
fi

PWD=$(cd "$(dirname "$0")" && pwd -P)
BUILD_DIR="${PWD}/../../${BUILD_DIR}"

CGO_ENABLED=0 gox \
    -verbose \
    -ldflags "${GOLDFLAGS}" \
    -gcflags=-trimpath="$(go env GOPATH)" \
    -os="${GOOS}" \
    -arch="${GOARCH}" \
    -osarch="!darwin/arm" \
    -output="${BUILD_DIR}/{{.OS}}-{{.Arch}}/{{.Dir}}" "${PWD}"/../../

printf "\033[36m==> Finalize binary\033[0m\n"

for platform in $(find "${BUILD_DIR}" -mindepth 1 -maxdepth 1 -type d | sort -u); do
    OSARCH=$(basename "${platform}")
    FULLNAME="${NAME}-${VERSION}-${OSARCH}"

    case "${OSARCH}" in
    "windows"*)
        mv "${platform}/${NAME}.exe" "${BUILD_DIR}/${FULLNAME}.exe"
        printf -- "--> %15s: bin/%s\n" "${OSARCH}" "${FULLNAME}.exe"
        ;;
    *)
        mv "${platform}/${NAME}" "${BUILD_DIR}/${FULLNAME}"
        printf -- "--> %15s: bin/%s\n" "${OSARCH}" "${FULLNAME}"
        ;;
    esac
done

cd "${BUILD_DIR}"
touch "${NAME}-${VERSION}.sha256sum"

for binary in $(find . -mindepth 1 -maxdepth 1 -type f | grep -v "${NAME}-${VERSION}.sha256sum" | sort); do
    binary=$(basename "${binary}")

    if command -v sha256sum >/dev/null; then
        sha256sum "${binary}" >>"${NAME}-${VERSION}.sha256sum"
    elif command -v shasum >/dev/null; then
        shasum -a256 "${binary}" >>"${NAME}-${VERSION}.sha256sum"
    fi
done

cd - >/dev/null 2>&1
printf -- "\n--> %15s: bin/%s\n" "sha256sum" "${NAME}-${VERSION}.sha256sum"
