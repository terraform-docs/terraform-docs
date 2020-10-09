package version

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// current version
const dev = "v0.10.2-alpha"

// Provisioned by ldflags
var (
	version    string
	commitHash string
	buildDate  string
)

// Load defaults for info variables
func init() {
	if version == "" {
		version = dev
	}
	if version == "v0.0.0-" { // building in a directory which is not a git repository
		version = dev
	}
	if commitHash == "" {
		commitHash = dev
	}
	if buildDate == "" {
		buildDate = time.Now().Format(time.RFC3339)
	}
}

// Full return the full version of the binary including commit hash and build date
func Full() string {
	if !strings.HasSuffix(version, commitHash) {
		version += " " + commitHash
	}
	osArch := runtime.GOOS + "/" + runtime.GOARCH
	return fmt.Sprintf("%s %s BuildDate: %s", version, osArch, buildDate)
}
