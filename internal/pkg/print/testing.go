package print

import (
	"io/ioutil"
	"path/filepath"
)

// ReadGoldenFile reads a .golden file from test data by name.
func ReadGoldenFile(name string) (string, error) {
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
