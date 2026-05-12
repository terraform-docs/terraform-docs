/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"strings"

	"github.com/opentofu/opentofu-schema/module"
	tfaddr "github.com/opentofu/registry-address"
)

// moduleSourceAndVersion flattens the typed `module.SourceAddr` back into a
// (source, version) pair compatible with the historical
// `terraform-config-inspect` output. Registry addresses surface their
// `version = "..."` argument as the version; remote/unknown sources fall back
// to the `?ref=...` query-string form via formatSource.
func moduleSourceAndVersion(moduleCall *module.DeclaredModuleCall) (string, string) {
	declaredVersion := ""
	if len(moduleCall.Version) > 0 {
		declaredVersion = moduleCall.Version.String()
	}

	switch source := moduleCall.SourceAddr.(type) {
	case tfaddr.Module:
		// registry address version comes from the `version = "..."` arg.
		return source.ForDisplay(), declaredVersion
	case module.LocalSourceAddr:
		return string(source), declaredVersion
	case module.RemoteSourceAddr, module.UnknownSourceAddr:
		// Remote/unknown sources may carry `?ref=...` which should surface as
		// version. Prefer RawSourceAddr to preserve the user's original syntax
		// (e.g. SCP-style `git@github.com:org/repo` rather than go-getter's
		// canonicalized `git::ssh://git@github.com/org/repo`).
		_ = source
		return formatSource(moduleCall.RawSourceAddr, declaredVersion)
	default:
		// nil SourceAddr falls back to raw string.
		return formatSource(moduleCall.RawSourceAddr, declaredVersion)
	}
}

// formatSource splits a `?ref=...` suffix off a remote module source URL
// and returns it as the version when no explicit `version = "..."` was
// declared. If the source has no `?ref=...` suffix (or version is already
// set), the inputs are returned unchanged.
func formatSource(s, v string) (source, version string) {
	substr := "?ref="

	if v != "" {
		return s, v
	}

	pos := strings.LastIndex(s, substr)
	if pos == -1 {
		return s, version
	}

	adjustedPos := pos + len(substr)
	if adjustedPos >= len(s) {
		return s, version
	}

	source = s[0:pos]
	version = s[adjustedPos:]

	return source, version
}
