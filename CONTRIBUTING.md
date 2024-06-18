# Contributing to terraform-docs

Welcome, and thank you for considering contributing to terraform-docs. We
encourage you to help out by raising issues, improving documentation, fixing
bugs, or adding new features.

If you're interested in contributing please start by reading this document. If
you have any questions at all, or don't know where to start, please reach out to
us on [Slack].

## Development Requirements

- [Go] 1.22+
- [goimports]
- [golangci-lint]

For releasing:

- [goreleaser]

You can install all the required tools with `make tools` or individually with
their own make targets. (please refer to Makefile for more details)

## Contributing Code

To contribute bug fixes or features to terraform-docs:

1. Communicate your intent.
1. Make your changes.
1. Test your changes.
1. Update documentation and examples.
1. Open a Pull Request (PR).

Communicating your intent lets the terraform-docs maintainers know that you intend
to contribute, and how. This sets you up for success - you can avoid duplicating
an effort that may already be underway, adding a feature that may be rejected,
or heading down a path that you would be steered away from at review time. The
best way to communicate your intent is via a detailed GitHub issue. Take a look
first to see if there's already an issue relating to the thing you'd like to
contribute. If there isn't, please raise a new one! Let us know what you'd like
to work on, and why.

Be sure to practice [good git commit hygiene] as you make your changes. All but
the smallest changes should be broken up into a few commits that tell a story.
Use your git commits to provide context for the folks who will review PR, and
the folks who will be spelunking the codebase in the months and years to come.
Ensure each of your commits is signed-off in compliance with the [Developer
Certificate of Origin] by using `git commit -s`.

All codes must be formatted properly, we use `goimports` which wrappes around
[gofmt] to keep the code in unified format. You can use `make goimports` to install
and `make fmt` to format your code. Make sure your code doesn't have issues with
`make checkfmt` and `make lint` before submission.

Formatter tests are separated into different groups in order to be able to improve
and maintain the code coverage and also prevent duplicating and running redundant
test cases. For example there are `OnlyXXX` tests which makes sure all sections are
hidden and only one is visible. This removes the need to test for `NotXXX` (which is
no deprecated and removed). Please make sure when a new feature is added or an issue
is fixed the corresponding `.golden` files and test cases are also update accordingly.

Once your change is written, tested, and documented the final step is to have it
reviewed! You'll be presented with a template and a small checklist when you
open a PR. Please read the template and fill out the checklist. Please make all
PR request changes in subsequent commits. This allows your reviewers to see what
has changed as you address their comments. Be mindful of your commit history as
you do this - avoid commit messages like "Address review feedback" if possible.
If doing so is difficult a good alternative is to rewrite your commit history to
clean them up after your PR is approved but before it is merged.

In summary, please:

* Discuss your change in a GitHub issue before you start.
* Use your Git commit messages to communicate your intent to your reviewers.
* Sign-off on all Git commits by running `git commit -s`
* Add or update tests for all changes.
* Preempt common [code review comments] and [test review comments].
* Update all relevant documentation and examples.
* Don't force push to address review feedback. Your commits should tell a story.
* If necessary, tidy up your git commit history once your PR is approved.

Thank you for reading through our contributing guide! We appreciate you taking
the time to ensure your contributions are high quality and easy for our community
to review and accept.

[Slack]: https://slack.terraform-docs.io/
[Go]: https://golang.org/
[goimports]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
[golangci-lint]: https://github.com/golangci/golangci-lint
[goreleaser]: https://goreleaser.com/
[good git commit hygiene]: https://www.futurelearn.com/info/blog/telling-stories-with-your-git-history
[Developer Certificate of Origin]: https://github.com/apps/dco
[gofmt]: https://blog.golang.org/go-fmt-your-code
[code review comments]: https://github.com/golang/go/wiki/CodeReviewComments
[test review comments]: https://github.com/golang/go/wiki/TestComments
