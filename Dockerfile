# Copyright 2021 The terraform-docs Authors.
#
# Licensed under the MIT license (the "License"); you may not
# use this file except in compliance with the License.
#
# You may obtain a copy of the License at the LICENSE file in
# the root directory of this source tree.

FROM golang:1.15.6-alpine AS builder

RUN apk add --update --no-cache make

WORKDIR /go/src/terraform-docs

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

################

# Use empty base image
FROM scratch

# Copy static executable for terraform-docs
COPY --from=builder /go/src/terraform-docs/bin/linux-amd64/terraform-docs /usr/local/bin/

# Set entrypoint
ENTRYPOINT ["terraform-docs"]
