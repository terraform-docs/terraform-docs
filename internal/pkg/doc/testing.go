package doc

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-testing-interface"
)

// TestDoc creates a *Doc from a Terraform module.
func TestDoc(t testing.T, path string) *Doc {
	path, err := absPath(testDataPath(path))
	if err != nil {
		t.Fatal(err)
	}

	doc, err := CreateFromPaths([]string{path})
	if err != nil {
		t.Fatal(err)
	}

	return doc
}

// TestDocFromFile creates a *Doc from a Terraform file.
func TestDocFromFile(t testing.T, path string, name string) *Doc {
	path, err := absPath(filepath.Join(testDataPath(path), name))

	if err != nil {
		t.Fatal(err)
	}

	doc, err := CreateFromPaths([]string{path})
	if err != nil {
		t.Fatal(err)
	}

	return doc
}

func absPath(relative string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, relative), nil
}

func testDataPath(path string) string {
	return filepath.Join(path, "..", "doc", "testdata")
}
