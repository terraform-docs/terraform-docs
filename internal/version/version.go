/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package version

import (
	"fmt"
	"runtime"
)

// current version
const (
	coreVersion = "0.20.0"
	prerelease  = "alpha"
)

// Provisioned by ldflags
var commit string

// Core return the core version.
func Core() string {
	return coreVersion
}

// Short return the version with pre-release, if available.
func Short() string {
	v := coreVersion

	if prerelease != "" {
		v += "-" + prerelease
	}

	return v
}

// Full return the full version including pre-release, commit hash, runtime os and arch.
func Full() string {
	if commit != "" && commit[:1] != " " {
		commit = " " + commit
	}

	return fmt.Sprintf("v%s%s %s/%s", Short(), commit, runtime.GOOS, runtime.GOARCH)
}
