NAME        := terraform-docs
VENDOR      := segmentio
DESCRIPTION := Generate docs from Terraform modules
MAINTAINER  := Martin Etmajer <metmajer@getcloudnative.io>
URL         := https://github.com/$(VENDOR)/$(NAME)
LICENSE     := MIT

VERSION     := $(shell cat ./VERSION)

GOBUILD     := go build -ldflags "-X main.version=$(VERSION)"
GOPKGS      := $(shell go list ./... | grep -v /vendor)


.PHONY: all
all: clean deps lint test build

.PHONY: authors
authors:
	git log --all --format='%aN <%cE>' | sort -u | egrep -v noreply > AUTHORS

.PHONY: build
build: authors build-darwin-amd64 build-freebsd-amd64 build-linux-amd64 build-windows-amd64

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(NAME)-v$(VERSION)-darwin-amd64

build-freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 $(GOBUILD) -o bin/$(NAME)-v$(VERSION)-freebsd-amd64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(NAME)-v$(VERSION)-linux-amd64

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(NAME)-v$(VERSION)-windows-amd64.exe

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: lint
lint:
	gometalinter --config gometalinter.json ./...

.PHONY: deps
deps:
	dep ensure

.PHONY: release
release:
	git tag -a v$(VERSION) -m v$(VERSION)
	git push --tags

.PHONY: test
test:
	go test -v $(GOPKGS)
