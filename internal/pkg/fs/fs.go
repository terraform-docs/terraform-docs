package fs

import (
	"os"
)

// DirectoryExists indicates if path exists and is a directory.
func DirectoryExists(path string) bool {
	stats, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return stats.IsDir()
}

// FileExists indicates if a path exists and is a file.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
