
v0.3.0 / 2017-10-22
===================

  * auto version

v0.2.0 / 2017-08-15
===================

  * fixed name handling to handle unquoted hcl variable names.
  * fix typo
  * Prefer leading comments over description for outputs to maintain compatability.
  * *: add --no-required option
  * doc: snakecase -> camelcase
  * Add support for printing the variable 'type' in Markdown. Currently only markdown supported, but trivial to add to other outputs.
  * Add support for reading `description` tag from `output` resources. Fixes #24
  * Add note about installing with Homebrew
  * Strip # prefix from comments
  * add proper license

v0.1.0 / 2017-03-21
==================

  * main: add support for files (@s-urbaniak)
  * closes #17 Updated example in README (@nwalke)
  * Closes #18 Added a very basic sort to inputs and outputs (@nwalke)
  * doc: Account for single whitespace after comment character in header (@paybyphone)
  * print/markdown: Better markdown description normalizations (@paybyphone)
  * print/markdown: Added line break conversion for outputs (@paybyphone)
  * doc: placeholder for list types (@paybyphone)
  * doc: Allow top-level comments for variables when description missing (@paybyphone)
  * print/markdown: Replace table cell newlines with HTML line breaks (@paybyphone)


v0.0.2 / 2016-06-29
==================

  * doc: fix map type
  * add more install notes
