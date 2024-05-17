---
title: "Installation"
description: "terraform-docs installation guide"
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

or

```bash
brew install terraform-docs/tap/terraform-docs
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

terraform-docs can be run as a container by mounting a directory with `.tf`
files in it and run the following command:

```bash
docker run --rm --volume "$(pwd):/terraform-docs" -u $(id -u) quay.io/terraform-docs/terraform-docs:0.17.0 markdown /terraform-docs
```

If `output.file` is not enabled for this module, generated output can be redirected
back to a file:

```bash
docker run --rm --volume "$(pwd):/terraform-docs" -u $(id -u) quay.io/terraform-docs/terraform-docs:0.17.0 markdown /terraform-docs > doc.md
```

{{< alert type="primary" >}}
Docker tag `latest` refers to _latest_ stable released version and `edge` refers
to HEAD of `master` at any given point in time. And any named version tags are
identical to the official GitHub releases without leading `v`.
{{< /alert >}}

## Pre-compiled Binary

Stable binaries are available on the GitHub [Release] page. To install, download
the file for your platform from "Assets" and place it into your `$PATH`:

```bash
curl -sSLo ./terraform-docs.tar.gz https://terraform-docs.io/dl/v0.17.0/terraform-docs-v0.17.0-$(uname)-amd64.tar.gz
tar -xzf terraform-docs.tar.gz
chmod +x terraform-docs
mv terraform-docs /some-dir-in-your-PATH/terraform-docs
```

{{< alert type="primary" >}}
Windows releases are in `ZIP` format.
{{< /alert >}}

## Go Users

The latest version can be installed using `go install` or `go get`:

```bash
# go1.17+
go install github.com/terraform-docs/terraform-docs@v0.17.0
```

```bash
# go1.16
GO111MODULE="on" go get github.com/terraform-docs/terraform-docs@v0.17.0
```

{{< alert type="warning" >}}
To download any version **before** `v0.9.1` (inclusive) you need to use to
old module namespace (`segmentio`):
{{< /alert >}}

```bash
# only for v0.9.1 and before
GO111MODULE="on" go get github.com/segmentio/terraform-docs@v0.9.1
```

{{< alert type="primary" >}}
Please use the latest Go to do this, minimum `go1.16` is required.
{{< /alert >}}

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

```zsh
terraform-docs completion zsh > /usr/local/share/zsh/site-functions/_terraform-docs
autoload -U compinit && compinit
```

### ohmyzsh

```zsh
terraform-docs completion zsh > ~/.oh-my-zsh/completions/_terraform-docs
omz reload
```

### fish

```fish
terraform-docs completion fish > ~/.config/fish/completions/terraform-docs.fish
```

To make this change permanent, the above commands can be added to `~/.profile` file.

[Chocolatey]: https://www.chocolatey.org
[Homebrew]: https://brew.sh
[Release]: https://github.com/terraform-docs/terraform-docs/releases
[Scoop]: https://scoop.sh/
