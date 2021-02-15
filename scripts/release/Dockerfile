# Copyright 2021 The terraform-docs Authors.
#
# Licensed under the MIT license (the "License"); you may not
# use this file except in compliance with the License.
#
# You may obtain a copy of the License at the LICENSE file in
# the root directory of this source tree.

# Use empty base image
FROM scratch

# Copy static executable for terraform-docs
COPY terraform-docs /usr/local/bin/terraform-docs

# Set entrypoint
ENTRYPOINT ["terraform-docs"]
