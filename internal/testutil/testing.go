package testutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// GetModule returns 'example' Module
func GetModule(options *module.Options) (*tfconf.Module, error) {
	path, err := getExampleFolder()
	if err != nil {
		return nil, err
	}
	options.Path = path
	if options.OutputValues {
		options.OutputValuesPath = filepath.Join(path, options.OutputValuesPath)
	}
	tfmodule, err := module.LoadWithOptions(options)
	if err != nil {
		return nil, err
	}
	return tfmodule, nil
}

// GetExpected returns 'example' Module and expected Golden file content
func GetExpected(format, name string) (string, error) {
	path := filepath.Join(testDataPath(), format, name+".golden")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getExampleFolder() (string, error) {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "..", "..", "examples")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	return path, nil
}

func testDataPath() string {
	return filepath.Join("testdata")
}
