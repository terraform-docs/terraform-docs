FROM golang:1.9-alpine AS golang


WORKDIR /go/src/github.com/segmentio/terraform-docs 
COPY . /go/src/github.com/segmentio/terraform-docs/ 

RUN go build && \
  go test ./...

FROM alpine:3.6
COPY --from=golang /go/src/github.com/segmentio/terraform-docs/terraform-docs /usr/local/bin
WORKDIR /workspace

ENTRYPOINT [ "/usr/local/bin/terraform-docs" ]
CMD [ "md", "." ]
