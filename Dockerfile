# Copyright 2021 The terraform-docs Authors.
#
# Licensed under the MIT license (the "License"); you may not
# use this file except in compliance with the License.
#
# You may obtain a copy of the License at the LICENSE file in
# the root directory of this source tree.

FROM golang:1.18.3-alpine AS builder

RUN apk add --update --no-cache make

WORKDIR /go/src/terraform-docs

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

################

FROM scratch

COPY --from=builder /go/src/terraform-docs/bin/linux-amd64/terraform-docs /usr/local/bin/

ENTRYPOINT ["terraform-docs"]
