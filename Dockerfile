FROM golang:1.9-alpine AS golang


WORKDIR /go/src/github.com/segmentio/terraform-docs 
COPY . /go/src/github.com/segmentio/terraform-docs/ 

RUN go build && \
  go test ./...

FROM docker-cd.artifactory.corp.code42.com/c42/cloud-workstation:1.2.0-rc.11 AS scripts
RUN git clone ssh://stash.corp.code42.com:7999/cd/version42.git /opt/code42/versions 

#FROM docker-cd.artifactory.corp.code42.com/c42/cloud-workstation:1.2.0-rc.11
FROM alpine:3.6
#FROM local/cloud-workstation
COPY --from=golang /go/src/github.com/segmentio/terraform-docs/terraform-docs /usr/local/bin
WORKDIR /workspace

ENTRYPOINT [ "/usr/local/bin/terraform-docs" ]
CMD [ "md", "." ]
