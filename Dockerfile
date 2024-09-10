# Copyright 2021 The terraform-docs Authors.
#
# Licensed under the MIT license (the "License"); you may not
# use this file except in compliance with the License.
#
# You may obtain a copy of the License at the LICENSE file in
# the root directory of this source tree.

FROM docker.io/library/golang:1.23.1-alpine AS builder

RUN apk add --update --no-cache make

WORKDIR /go/src/terraform-docs

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

################

FROM docker.io/library/alpine:3.20.3

# Mitigate CVE-2023-5363
RUN apk add --no-cache --upgrade "openssl>=3.1.4-r1"

COPY --from=builder /go/src/terraform-docs/bin/linux-*/terraform-docs /usr/local/bin/

ENTRYPOINT ["terraform-docs"]
