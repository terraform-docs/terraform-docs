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

PWD=$(cd "$(dirname "$0")" && pwd -P)
PWD="${PWD}/../.."

# Make sure site/ folder does not exist
rm -rf "${PWD}"/site

# Clone the website repository locally
git clone -b main https://github.com/terraform-docs/website "${PWD}"/site

# Update website content
rm -rf "${PWD}"/site/content/
cp -r "${PWD}"/docs/ "${PWD}"/site/content/
