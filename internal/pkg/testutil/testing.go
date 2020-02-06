package testutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
)

// GetExpected returns 'example' Module and expected Golden file content
func GetExpected(goldenFile string) (*tfconf.Module, string, error) {
	path, err := getExampleFolder()
	if err != nil {
		return nil, "", err
	}

	module, err := tfconf.CreateModule(path, "")
	if err != nil {
		return nil, "", err
	}

	expected, err := readGoldenFile(goldenFile)

	return module, expected, err
}

func getExampleFolder() (string, error) {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "..", "..", "..", "examples")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	return path, nil
}

func readGoldenFile(name string) (string, error) {
	path := filepath.Join(testDataPath(), name+".golden")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func testDataPath() string {
	return filepath.Join("testdata")
}
