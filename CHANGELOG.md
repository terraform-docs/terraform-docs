# CHANGELOG

<a name="unreleased"></a>
## [Unreleased]

### Bug Fixes
- Mark variables not required if default set to null ([#221](https://github.com/segmentio/terraform-docs/issues/221))
- --no-header should not attempt reading main.tf file ([#224](https://github.com/segmentio/terraform-docs/issues/224))
- Fix type conversion for numbers ([#204](https://github.com/segmentio/terraform-docs/issues/204))

### Code Refactoring
- Add Default value types for better marshalling ([#196](https://github.com/segmentio/terraform-docs/issues/196))
- Introduce Format interface and expose to public pkg ([#195](https://github.com/segmentio/terraform-docs/issues/195))
- Add tfconf.Options to load Module with ([#193](https://github.com/segmentio/terraform-docs/issues/193))

### Documentation
- Enhance automatic document generation ([#227](https://github.com/segmentio/terraform-docs/issues/227))
- Add installation guide for Windows users ([#218](https://github.com/segmentio/terraform-docs/issues/218))
- Put reference to usage, cli, etc. in user guide ([#216](https://github.com/segmentio/terraform-docs/issues/216))
- Example git hook to keep module docs up to date ([#214](https://github.com/segmentio/terraform-docs/issues/214))
- Auto generate formats document from examples ([#192](https://github.com/segmentio/terraform-docs/issues/192))

### Enhancements
- Add extensive tests coverage for all the packages ([#208](https://github.com/segmentio/terraform-docs/issues/208))

### Features
- Add support for tfvars hcl and json commands ([#226](https://github.com/segmentio/terraform-docs/issues/226))
- Allow hiding the "Sensitive" column in markdown ([#223](https://github.com/segmentio/terraform-docs/issues/223))
- Add section for module requirements ([#222](https://github.com/segmentio/terraform-docs/issues/222))
- Add support for fetching the module header from any file ([#217](https://github.com/segmentio/terraform-docs/issues/217))
- Add support for XML renderer ([#198](https://github.com/segmentio/terraform-docs/issues/198))
- Show sensitivity of the output value in rendered result ([#207](https://github.com/segmentio/terraform-docs/issues/207))
- Extract and render output values from Terraform ([#191](https://github.com/segmentio/terraform-docs/issues/191))
- Render formatted results with go templates ([#177](https://github.com/segmentio/terraform-docs/issues/177))
- Add support for YAML renderer ([#189](https://github.com/segmentio/terraform-docs/issues/189))


<a name="v0.8.2"></a>
## [v0.8.2] - 2020-02-03

### Bug Fixes
- Do not escape markdown table inside module header ([#186](https://github.com/segmentio/terraform-docs/issues/186))
- Add double space only at the end of paragraph lines ([#185](https://github.com/segmentio/terraform-docs/issues/185))
- Preserve asterisk list in header and fix escaping ([#179](https://github.com/segmentio/terraform-docs/issues/179))
- Add newline between code block and trailing lines ([#184](https://github.com/segmentio/terraform-docs/issues/184))


<a name="v0.8.1"></a>
## [v0.8.1] - 2020-01-21

### Bug Fixes
- Show native map and list as default value in JSON ([#174](https://github.com/segmentio/terraform-docs/issues/174))


<a name="v0.8.0"></a>
## [v0.8.0] - 2020-01-17

### Bug Fixes
- Do not escape any characters of a URL ([#170](https://github.com/segmentio/terraform-docs/issues/170))
- Add double space at the end of multi-lines paragraph ([#169](https://github.com/segmentio/terraform-docs/issues/169))
- Show empty JSON properties, as 'null' for all types ([#166](https://github.com/segmentio/terraform-docs/issues/166))
- Show all JSON properties, empty or null ([#160](https://github.com/segmentio/terraform-docs/issues/160))
- Do not escape strings inside code blocks ([#155](https://github.com/segmentio/terraform-docs/issues/155))
- Read leading module header from main.tf ([#154](https://github.com/segmentio/terraform-docs/issues/154))
- Read leading comment lines if description is not provided ([#151](https://github.com/segmentio/terraform-docs/issues/151))
- Reimplement '--no-sort' to be compatible with Terraform 0.12 configuration ([#141](https://github.com/segmentio/terraform-docs/issues/141))

### Code Refactoring
- Move doc.Doc to tfconf.Module ([#136](https://github.com/segmentio/terraform-docs/issues/136))

### Documentation
- Initial commit of usage documentation ([#162](https://github.com/segmentio/terraform-docs/issues/162))
- Deprecate accepting files as commands param ([#163](https://github.com/segmentio/terraform-docs/issues/163))
- Update Module internal documentaion

### Enhancements
- Rename flag to '--sort-by-required' ([#150](https://github.com/segmentio/terraform-docs/issues/150))
- Mark '--with-aggregate-type-defaults' as deprecated ([#148](https://github.com/segmentio/terraform-docs/issues/148))
- Bump homebrew formula version on release ([#135](https://github.com/segmentio/terraform-docs/issues/135))
- Enable new go linters and fix the existing issues ([#132](https://github.com/segmentio/terraform-docs/issues/132))

### Features
- Add '--no-escape' flag to 'json' command ([#147](https://github.com/segmentio/terraform-docs/issues/147))
- Add flags to not show different sections ([#144](https://github.com/segmentio/terraform-docs/issues/144))
- Add '--no-color' flag to 'pretty' command ([#143](https://github.com/segmentio/terraform-docs/issues/143))
- Show 'providers' information ([#140](https://github.com/segmentio/terraform-docs/issues/140))
- Bump golang to latest v1.13 ([#133](https://github.com/segmentio/terraform-docs/issues/133))
- Support Terraform 0.12.x configuration ([#113](https://github.com/segmentio/terraform-docs/issues/113))

### BREAKING CHANGE

- With Terraform 0.12 ability to generate
output from file has been deprecated in favor of from folder
which contains one or more `.tf` files.

- In the JSON format response, list of "Inputs"
has been renamed to `inputs`.
- In the JSON format response, list of "Outputs" has been renamed
to `outputs`.

- In the JSON format respone, module "Comment" has been renamed to module `header`.

- For simplicity we've decided to
deprecated the old `--sort-inputs-by-required` flag
to the simpler and more generic `--sort--by-required`.
The deprecated flags will get removed second release
from now.

- As of Terraform 0.12, the default value of
input variables are shown in full JSON format (if available)
and `--with-aggregate-type-defaults` is not needed anymore.
The flag is marked as soft deprecated and will get removed in
the second release from now.

- With Terraform 0.12 the information about `providers` being used in the module will be generated by default. This will cause the first generation of documents with the latest release of `terraform-docs` binary be slightly different than before, now there will be `Providers` section in Markdown and `providers` block in JSON. You can ignore this by using new `--no-providers` flag if you choose to.


<a name="v0.7.0"></a>
## [v0.7.0] - 2019-12-12

- Update Changelog
- Release version v0.7.0
- Use Github Actions instead of Circle CI ([#124](https://github.com/segmentio/terraform-docs/issues/124))
- Enhance release scripts
- Generate release note based on the current tag changelog
- Update Installation and Code Completion in README
- Code blocks support for all formats. Single line break support ([#123](https://github.com/segmentio/terraform-docs/issues/123))
- Update Changelog
- target deps was missing (required by all) ([#126](https://github.com/segmentio/terraform-docs/issues/126))
- Enhance Makefile and add editorconfig ([#115](https://github.com/segmentio/terraform-docs/issues/115))
- Add support for controlling the indentation of Markdown headers ([#120](https://github.com/segmentio/terraform-docs/issues/120))
- Refactor Settings for better performance ([#119](https://github.com/segmentio/terraform-docs/issues/119))
- Add --no-escape flag for Markdown printer ([#117](https://github.com/segmentio/terraform-docs/issues/117))
- Update Changelog.
- Use Cobra CLI instead of docopt ([#116](https://github.com/segmentio/terraform-docs/issues/116))
- Update Changelog.
- Escape pipe character when generating Markdown ([#114](https://github.com/segmentio/terraform-docs/issues/114))
- Add appropriate Changelog header.
- Complete development requirements documentation.
- Configure git-chglog to not show git merge commit messages.
- Add Changelog generation via git-chglog. ([#104](https://github.com/segmentio/terraform-docs/issues/104))
- Remove occurrence of gometalinter from CircleCI config.
- Replace dep with Go Modules ([#100](https://github.com/segmentio/terraform-docs/issues/100))
- Replace gometalinter with golangci-lint. ([#103](https://github.com/segmentio/terraform-docs/issues/103))
- Add 'enhancement' section to pull request template ([#101](https://github.com/segmentio/terraform-docs/issues/101))
- Fix typo in options documentation ([#98](https://github.com/segmentio/terraform-docs/issues/98))
- Bump Homebrew formula to 0.6.0.


<a name="v0.6.0"></a>
## [v0.6.0] - 2018-12-13

- Bump version to 0.6.0.
- Unify default values of inputs ([#97](https://github.com/segmentio/terraform-docs/issues/97))
- Unify description text of inputs and outputs ([#96](https://github.com/segmentio/terraform-docs/issues/96))
- Capitalize headings in documentation.
- Fix Markdown lint errors and enhancement in README ([#94](https://github.com/segmentio/terraform-docs/issues/94))
- Update project documentation.
- Move Terraform test configuration to folder 'examples'.
- Capitalize the word 'markdown' in documentation.
- Purge History.md file.
- Add support for rendering Markdown documents ([#81](https://github.com/segmentio/terraform-docs/issues/81))
- Migrate from github.com/tj/docopt to github.com/docopt/docopt-go ([#91](https://github.com/segmentio/terraform-docs/issues/91))
- Fix authors target in Makefile to get 'Author''s email not 'Committer' ([#90](https://github.com/segmentio/terraform-docs/issues/90))
- Add requirement to discuss suitability of a new feature in an issue before submitting a pull request.


<a name="v0.5.0"></a>
## [v0.5.0] - 2018-10-24

- Bump version to 0.5.0.
- Add support to print Markdown files with underscored variable names escaped ([#48](https://github.com/segmentio/terraform-docs/issues/48)) ([#63](https://github.com/segmentio/terraform-docs/issues/63))
- Add CircleCI badge.
- Fix homebrew formula. ([#75](https://github.com/segmentio/terraform-docs/issues/75))
- Add sort by "required" and then name ([#43](https://github.com/segmentio/terraform-docs/issues/43))
- Add Homebrew formula. ([#68](https://github.com/segmentio/terraform-docs/issues/68))


<a name="v0.4.5"></a>
## [v0.4.5] - 2018-10-07

- Bump version to 0.4.5.
- Allow unquoted item names. Fixes [#64](https://github.com/segmentio/terraform-docs/issues/64) ([#70](https://github.com/segmentio/terraform-docs/issues/70))
- Change build dir structure ([#74](https://github.com/segmentio/terraform-docs/issues/74))
- Update makefile to fix Windows build filename ([#72](https://github.com/segmentio/terraform-docs/issues/72))
- Remove extra newlines between comments and inputs/outputs to fix MarkDownLint warnings ([#66](https://github.com/segmentio/terraform-docs/issues/66))
- Fix loading of comments from main.tf on Windows ([#65](https://github.com/segmentio/terraform-docs/issues/65))


<a name="v0.4.0"></a>
## [v0.4.0] - 2018-09-23

- Bump version to 0.4.0.
- Add option --with-aggregate-type-defaults to enable printing of default values for types 'list' and 'map'. ([#53](https://github.com/segmentio/terraform-docs/issues/53))
- Add option --no-sort to omit sorted rendering of inputs and outputs. ([#61](https://github.com/segmentio/terraform-docs/issues/61))
- Refactor package 'doc' for better modularity. ([#60](https://github.com/segmentio/terraform-docs/issues/60))
- Refactor package 'print' for better modularity. ([#59](https://github.com/segmentio/terraform-docs/issues/59))
- Complete CircleCI config. Add vendor directory. ([#58](https://github.com/segmentio/terraform-docs/issues/58))
- Update AUTHORS.
- Add issue and pull request templates.
- Add contributing guidelines.
- Fix indentation.
- Update documentation and license to reflect the terraform-docs authors.
- Update documentation.
- Move packages 'doc' and 'print' to internal/pkg.
- Add automated tests for package 'print'.
- Add automated tests for package 'doc'.
- Refactor code in main and prepare for tests.
- Add documentation of --version option.
- Add make target to run Go tests.
- Add make target to create and push a Git tag.
- Add make target to check Go sources for errors and warnings. Remove unused code.
- Add make target to create AUTHORS file from git logs.
- Add make target to clean the workspace.
- Add dependency management using go deps.
- Add Makefile header and build target.
- Add base CI config ([#56](https://github.com/segmentio/terraform-docs/issues/56))
- Add Maintenance section to Readme ([#55](https://github.com/segmentio/terraform-docs/issues/55))
- Update Readme.md
- Merge pull request [#44](https://github.com/segmentio/terraform-docs/issues/44) from coveo/description-before-comments
- If there is a description on an output, it should be considered before the preceding comment


<a name="v0.3.0"></a>
## [v0.3.0] - 2017-10-22

- Release v0.3.0
- auto version
- Merge pull request [#39](https://github.com/segmentio/terraform-docs/issues/39) from BWITS/[#38](https://github.com/segmentio/terraform-docs/issues/38)
- bugfix/[#38](https://github.com/segmentio/terraform-docs/issues/38)
- Merge pull request [#36](https://github.com/segmentio/terraform-docs/issues/36) from nwalke/fix_version_string
- closes [#35](https://github.com/segmentio/terraform-docs/issues/35) Updated version string


<a name="v0.2.0"></a>
## [v0.2.0] - 2017-08-15

- Release v0.2.0


<a name="v0.1.1"></a>
## [v0.1.1] - 2017-08-15

- Release v0.1.1
- Merge pull request [#34](https://github.com/segmentio/terraform-docs/issues/34) from COzero/master
- Merge pull request [#1](https://github.com/segmentio/terraform-docs/issues/1) from COzero/unquoted_names
- fixed name handling to handle unquoted hcl variable names.
- Merge pull request [#31](https://github.com/segmentio/terraform-docs/issues/31) from BWITS/typo
- fix typo
- Merge pull request [#28](https://github.com/segmentio/terraform-docs/issues/28) from s-urbaniak/no-required
- Merge pull request [#25](https://github.com/segmentio/terraform-docs/issues/25) from fatmcgav/support_output_description
- Prefer leading comments over description for outputs to maintain compatability.
- *: add --no-required option
- snakecase -> camelcase
- Merge pull request [#27](https://github.com/segmentio/terraform-docs/issues/27) from fatmcgav/support_printing_type
- Add support for printing the variable 'type' in Markdown. Currently only markdown supported, but trivial to add to other outputs.
- Add support for reading `description` tag from `output` resources. Fixes [#24](https://github.com/segmentio/terraform-docs/issues/24)
- Merge pull request [#23](https://github.com/segmentio/terraform-docs/issues/23) from jacobwgillespie/patch-1
- Add note about installing with Homebrew
- Merge pull request [#22](https://github.com/segmentio/terraform-docs/issues/22) from jacobwgillespie/patch-1
- Strip # prefix from comments
- add proper license


<a name="v0.1.0"></a>
## [v0.1.0] - 2017-03-21

- Release v0.1.0
- Merge pull request [#21](https://github.com/segmentio/terraform-docs/issues/21) from s-urbaniak/files
- add support for files
- Merge pull request [#20](https://github.com/segmentio/terraform-docs/issues/20) from nwalke/update_readme_example
- closes [#17](https://github.com/segmentio/terraform-docs/issues/17) Updated example in README
- Merge pull request [#19](https://github.com/segmentio/terraform-docs/issues/19) from nwalke/add_sorting
- Closes [#18](https://github.com/segmentio/terraform-docs/issues/18) Added a very basic sort to inputs and outputs
- Merge pull request [#16](https://github.com/segmentio/terraform-docs/issues/16) from paybyphone/master
- Account for single whitespace after comment character in header
- print/markdown: Better markdown description normalizations
- print/markdown: Added line break conversion for outputs
- placeholder for list types
- Allow top-level comments for variables when description missing
- print/markdown: Replace table cell newlines with HTML line breaks
- Merge pull request [#13](https://github.com/segmentio/terraform-docs/issues/13) from jbussdieker/jbb-fix-heredoc-description
- Trim whitespace on markdown description too
- Fix HEREDOC descriptions


<a name="v0.0.2"></a>
## [v0.0.2] - 2016-06-29

- Release v0.0.2
- Merge pull request [#11](https://github.com/segmentio/terraform-docs/issues/11) from segmentio/fix-md
- wrap default values
- Merge pull request [#10](https://github.com/segmentio/terraform-docs/issues/10) from segmentio/fix-map-type
- fix map type
- add more install notes
- add dist


<a name="v0.0.1"></a>
## v0.0.1 - 2016-06-15

- Merge pull request [#5](https://github.com/segmentio/terraform-docs/issues/5) from segmentio/fix-comment
- use /** comment for module commnet
- actually print head comment
- img
- Merge pull request [#4](https://github.com/segmentio/terraform-docs/issues/4) from segmentio/layout
- fix view
- docs
- ignore comments with /** prefix
- add head comment
- cleanup
- cleanup
- ocd
- ocd
- update doc
- add installation
- cleanup
- add usage
- better md output
- add markdown output
- use comments as description
- ocd
- clean
- working
- Initial commit


[Unreleased]: https://github.com/segmentio/terraform-docs/compare/v0.8.2...HEAD
[v0.8.2]: https://github.com/segmentio/terraform-docs/compare/v0.8.1...v0.8.2
[v0.8.1]: https://github.com/segmentio/terraform-docs/compare/v0.8.0...v0.8.1
[v0.8.0]: https://github.com/segmentio/terraform-docs/compare/v0.7.0...v0.8.0
[v0.7.0]: https://github.com/segmentio/terraform-docs/compare/v0.6.0...v0.7.0
[v0.6.0]: https://github.com/segmentio/terraform-docs/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/segmentio/terraform-docs/compare/v0.4.5...v0.5.0
[v0.4.5]: https://github.com/segmentio/terraform-docs/compare/v0.4.0...v0.4.5
[v0.4.0]: https://github.com/segmentio/terraform-docs/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/segmentio/terraform-docs/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/segmentio/terraform-docs/compare/v0.1.1...v0.2.0
[v0.1.1]: https://github.com/segmentio/terraform-docs/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/segmentio/terraform-docs/compare/v0.0.2...v0.1.0
[v0.0.2]: https://github.com/segmentio/terraform-docs/compare/v0.0.1...v0.0.2
