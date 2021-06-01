---
title: "Installation"
description: "terraform-docs installation guide."
menu:
  docs:
    parent: "user-guide"
weight: 110
toc: true
---

`terraform-docs` is available on Linux, macOS, Windows, and FreeBSD platforms.

## Homebrew

If you are a macOS user, you can use [Homebrew].

```bash
brew install terraform-docs
```

## Windows

If you are a Windows user:

### Scoop

You can use [Scoop].

```bash
scoop bucket add terraform-docs https://github.com/terraform-docs/scoop-bucket
scoop install terraform-docs
```

### Chocolatey

or you can use [Chocolatey].

```bash
choco install terraform-docs
```

## Docker

You also can run `terraform-docs` as a container:

```bash
docker run quay.io/terraform-docs/terraform-docs:0.14.1
```

Docker tag `latest` refers to _latest_ stable released version and `edge`refers
to HEAD of `master` at any given point in time. And any named version tags are
identical to the official GitHub releases without leading `v`.

## Pre-compiled Binary

Stable binaries are available on the GitHub [Release] page. To install, download
the file for your platform from "Assets" and place it into your `$PATH`:

```bash
curl -sSLo ./terraform-docs.tar.gz https://terraform-docs.io/dl/v0.14.1/terraform-docs-v0.14.1-$(uname)-amd64.tar.gz
tar -xzf terraform-docs.tar.gz
chmod +x terraform-docs
mv terraform-docs /some-dir-in-your-PATH/terraform-docs
```

**Note:** Windows releases are in `ZIP` format.

## Go Users

The latest version can be installed using `go get`:

```bash
GO111MODULE="on" go get github.com/terraform-docs/terraform-docs@v0.14.1
```

**NOTE:** to download any version **before** `v0.9.1` (inclusive) you need to use to
old module namespace (`segmentio`):

```bash
# only for v0.9.1 and before
GO111MODULE="on" go get github.com/segmentio/terraform-docs@v0.9.1
```

**NOTE:** please use the latest Go to do this, minimum `go1.16` or greater.

This will put `terraform-docs` in `$(go env GOPATH)/bin`. If you encounter the error
`terraform-docs: command not found` after installation then you may need to either add
that directory to your `$PATH` as shown [here] or do a manual installation by cloning
the repo and run `make build` from the repository which will put `terraform-docs` in:

```bash
$(go env GOPATH)/src/github.com/terraform-docs/terraform-docs/bin/$(uname | tr '[:upper:]' '[:lower:]')-amd64/terraform-docs
```

## Code Completion

The code completion for `bash` or `zsh` can be installed as follow. Note that shell
auto-completion is not available on Windows platform.

### bash

```bash
terraform-docs completion bash > ~/.terraform-docs-completion
source ~/.terraform-docs-completion

# or the one-liner below

source <(terraform-docs completion bash)
```

### zsh

```bash
terraform-docs completion zsh > /usr/local/share/zsh/site-functions/_terraform-docs
autoload -U compinit && compinit
```

To make this change permanent, the above commands can be added to `~/.profile` file.

[Release]: https://github.com/terraform-docs/terraform-docs/releases
[Homebrew]: https://brew.sh
[Scoop]: https://scoop.sh/
[Chocolatey]: https://www.chocolatey.org
