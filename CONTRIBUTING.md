# Contributing

Thank you for your interest in this project!

We use *GitHub Issues* for trackings issues and features. You can make a contribution by:

1. Reporting an issue or making a feature request [here](https://github.com/terraform-docs/terraform-docs/issues).
2. Contributing code to this project by fixing an issue or adding a new feature (see below).

Before contributing a new feature, please discuss its suitability with the project maintainers in an issue first. Thanks!

## Development Requirements

For development:

- [Go](https://golang.org/) 1.15+
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [golangci-lint](https://github.com/golangci/golangci-lint)

For releasing:

- [gox](https://github.com/mitchellh/gox)
- [git-chlog](https://github.com/terraform-docs/git-chglog)

You can install required tools with `make tools` or individually, refer to Makefile for more details.

## Contribution Process

1. Fork and *git clone* [terraform-docs](https://github.com/terraform-docs/terraform-docs).
2. Create a new *git branch* from the master branch where you develop your changes.
3. Create a [Pull Request](https://help.github.com/articles/about-pull-requests/) for your contribution by following the instructions in the pull request template [here](https://github.com/terraform-docs/terraform-docs/pull).
4. Perform a code review with the project maintainers on the pull request. We may suggest changes, improvements or alternatives.
5. Once approved, your code will be merged into `master` and your name will be included in `AUTHORS`.

### Requirements

Pull requests have to meet the following requirements:

1. **Tests**: Code changes need to be tested with code and tests being located in the same folder (see packages [format](https://github.com/terraform-docs/terraform-docs/tree/master/internal/format/) for example). Make sure that your tests pass using `make test`.

2. **Documentation**: Pull requests need to update the [Formats Guide](/docs/FORMATS_GUIDE.md) and if need be the main [README](README.md) together with the code change. You can generate the format guides by using `make docs`.

3. **Commits**: Commits should be as small as possible while ensuring that each commit compiles and passes tests independently. [Write good commit messages](https://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html). Ensure each of your commits is signed-off in compliance with the [Developer Certificate of Origin](https://github.com/apps/dco) by using `git commit -s`. If needed, [squash your commits](https://davidwalsh.name/squash-commits-git) prior to submission.

4. **Code Style**: We use `goimports` which wrappes around [gofmt](https://blog.golang.org/go-fmt-your-code) to keep the code in unified format. You can use `make goimports` to install and `make fmt` to format your code. If useful, include code comments to support your intentions. Make sure your code doesn't have issues with `make checkfmt` and `make lint` before submission.

## Additional Resources

- [Golang Basics: Writing Unit Tests (Alex Ellis)](https://blog.alexellis.io/golang-writing-unit-tests/)
- [Advanced Testing in Go (Mitchell Hashimoto)](https://about.sourcegraph.com/go/advanced-testing-in-go/)
