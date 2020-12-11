FROM golang:1.15.2-alpine AS builder

RUN apk add --update --no-cache ca-certificates bash make gcc musl-dev git openssh wget curl

WORKDIR /go/src/terraform-docs

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

################

FROM alpine:3.12.2

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/terraform-docs/bin/linux-amd64/terraform-docs /usr/local/bin/

ENTRYPOINT ["terraform-docs"]
