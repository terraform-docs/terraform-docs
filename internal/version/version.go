package version

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// current version
const dev = "v0.9.0"

// Provisioned by ldflags
var (
	version    string
	commitHash string
	buildDate  string
)

func init() {
	// Load defaults for info variables
	if version == "" {
		version = dev
	}
	if commitHash == "" {
		commitHash = dev
	}
	if buildDate == "" {
		buildDate = time.Now().Format(time.RFC3339)
	}
}

// Version return the full version of the binary including commit hash and build date
func Version() string {
	if !strings.HasSuffix(version, commitHash) {
		version += " " + commitHash
	}
	osArch := runtime.GOOS + "/" + runtime.GOARCH
	return fmt.Sprintf("%s %s BuildDate: %s", version, osArch, buildDate)
}
