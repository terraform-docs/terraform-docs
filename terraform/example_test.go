package terraform

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraform-docs/terraform-docs/print"
)

func TestExampleFileExists(t *testing.T) {
	pathToTempFolder := ""
	tempFolderPattern := "exmplFileExstChck"

	dir, _ := ioutil.TempDir(pathToTempFolder, tempFolderPattern)
	tempFile, _ := os.CreateTemp(dir, "test.tf")
	defer os.RemoveAll(dir)

	tests := []struct {
		name          string
		pathToTest    string
		expectedValue bool
	}{
		{
			name:          "checkForExistingFolder",
			pathToTest:    dir,
			expectedValue: true,
		},
		{
			name:          "checkForNonExistingFolder",
			pathToTest:    filepath.Join(dir, "Non-existentFolder"),
			expectedValue: false,
		},
		{
			name:          "checkForExistingFile",
			pathToTest:    tempFile.Name(),
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expectedValue, FileExists(tt.pathToTest))
		})
	}

	assert.Equal(t, true, FileExists(dir))
}

func TestExampleGetChildFolders(t *testing.T) {
	//var expectedValue []string
	pathToTempFolder := ""
	tempFolderPattern := "exampleGetChildFolderCheck"

	checkFor2ChildFoldersDir, _ := os.MkdirTemp(pathToTempFolder, tempFolderPattern)
	checkFor2ChildFoldersChildDir1, _ := os.MkdirTemp(checkFor2ChildFoldersDir, "testDir1")
	checkFor2ChildFoldersChildDir2, _ := os.MkdirTemp(checkFor2ChildFoldersDir, "testDir2")
	defer os.RemoveAll(checkFor2ChildFoldersDir)

	checkForNonExistingChildFoldersDir, _ := os.MkdirTemp(pathToTempFolder, tempFolderPattern)
	defer os.RemoveAll(checkForNonExistingChildFoldersDir)

	checkFor1ChildFolderWithFilesDir, _ := os.MkdirTemp(pathToTempFolder, tempFolderPattern)
	checkFor1ChildFolderWithFilesChildDir1, _ := os.MkdirTemp(checkFor1ChildFolderWithFilesDir, "testDir1")
	os.CreateTemp(checkFor1ChildFolderWithFilesDir, "testFile1")
	defer os.RemoveAll(checkFor1ChildFolderWithFilesDir)

	tests := []struct {
		name          string
		pathToTest    string
		expectedValue []string
	}{
		{
			name:          "checkFor2ChildFolders",
			pathToTest:    checkFor2ChildFoldersDir,
			expectedValue: []string{checkFor2ChildFoldersChildDir1, checkFor2ChildFoldersChildDir2},
		},
		{
			name:          "checkForNonExistingChildFolders",
			pathToTest:    checkForNonExistingChildFoldersDir,
			expectedValue: []string{},
		},
		{
			name:          "checkFor1ChildFolderWithFiles",
			pathToTest:    checkFor1ChildFolderWithFilesDir,
			expectedValue: []string{checkFor1ChildFolderWithFilesChildDir1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expectedValue, GetChildFolders(tt.pathToTest))
		})
	}
}

func TestExampleHasChildMainTf(t *testing.T) {
	pathToTempFolder := ""
	tempFolderPattern := "exampleHasChildMainTfCheck"

	checkForNoChildFilesDir, _ := os.MkdirTemp(pathToTempFolder, tempFolderPattern)
	defer os.RemoveAll(checkForNoChildFilesDir)

	checkForMainTfDir, _ := os.MkdirTemp(pathToTempFolder, tempFolderPattern)
	os.Create(filepath.Join(checkForMainTfDir, "main.tf"))
	defer os.RemoveAll(checkForMainTfDir)

	checkForNoMainTfDir, _ := os.MkdirTemp(pathToTempFolder, tempFolderPattern)
	os.Create(filepath.Join(checkForNoMainTfDir, "notmain.tf"))
	os.Create(filepath.Join(checkForNoMainTfDir, "alsonotmain.tf"))
	defer os.RemoveAll(checkForNoMainTfDir)

	tests := []struct {
		name          string
		pathToTest    string
		expectedValue bool
	}{
		{
			name:          "checkForNoChildFiles",
			pathToTest:    checkForNoChildFilesDir,
			expectedValue: false,
		},
		{
			name:          "checkForMainTf",
			pathToTest:    checkForMainTfDir,
			expectedValue: true,
		},
		{
			name:          "checkForNoMainTf",
			pathToTest:    checkForNoMainTfDir,
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expectedValue, HasChildMainTf(tt.pathToTest))
		})
	}

}

func TestIsOnlyWhitespace(t *testing.T) {
	//var loader ExampleLoader
	tests := []struct {
		name          string
		stringToTest  string
		expectedValue bool
	}{
		{
			name:          "checkForOnlySpaces",
			stringToTest:  "  ",
			expectedValue: true,
		},
		{
			name:          "checkForSpacesAndNewLine",
			stringToTest:  "  \n ",
			expectedValue: true,
		},
		{
			name:          "checkForActualContent",
			stringToTest:  "Real Text Here",
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expectedValue, isOnlyWhitespace(tt.stringToTest))
		})
	}
}

func TestExampleLoader_SearchFolder(t *testing.T) {
	basePath := ""
	tempFolderPattern := "ExampleLoaderSearchFolderTests"
	// Create the ExampleLoader and perform the search
	loader := ExampleLoader{
		config:   print.DefaultConfig(),
		examples: nil, // Initialize with an empty slice
	}

	CheckFolderWithExampleDir, _ := os.MkdirTemp(basePath, tempFolderPattern)
	ExampleFolder, _ := os.MkdirTemp(CheckFolderWithExampleDir, tempFolderPattern)
	os.WriteFile(filepath.Join(ExampleFolder, "main.tf"), []byte("File Content"), 0777)
	defer os.RemoveAll(CheckFolderWithExampleDir)

	CheckFolderWithBlankExampleDir, _ := os.MkdirTemp(basePath, tempFolderPattern)
	BlankExampleFolder, _ := os.MkdirTemp(CheckFolderWithBlankExampleDir, tempFolderPattern)
	os.Create(filepath.Join(BlankExampleFolder, "main.tf"))
	defer os.RemoveAll(CheckFolderWithBlankExampleDir)

	CheckFolderWithoutMaintfDir, _ := os.MkdirTemp(basePath, tempFolderPattern)
	os.MkdirTemp(CheckFolderWithoutMaintfDir, tempFolderPattern)
	defer os.RemoveAll(CheckFolderWithoutMaintfDir)

	tests := []struct {
		name             string
		exampleFolder    string
		expectedError    bool
		expectedExamples int
	}{
		{
			name:             "CheckFolderWithExample",
			exampleFolder:    CheckFolderWithExampleDir,
			expectedError:    false,
			expectedExamples: 1,
		},
		{
			name:             "CheckFolderWithBlankExample",
			exampleFolder:    CheckFolderWithBlankExampleDir,
			expectedError:    false,
			expectedExamples: 0, //If main.tf is blank, it should be ignored
		},
		{
			name:             "Non-existing folder",
			exampleFolder:    "non_existing_folder",
			expectedError:    true,
			expectedExamples: 0,
		},
		{
			name:             "Folder without main.tf",
			exampleFolder:    CheckFolderWithoutMaintfDir,
			expectedError:    false,
			expectedExamples: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			loader.config.ExamplesFrom = tt.exampleFolder
			err := loader.SearchFolder()

			if tt.expectedError && err == nil {
				t.Error("Expected an error, but got nil")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}

			assert.Equal(tt.expectedExamples, len(loader.examples))
			loader.examples = nil
		})
	}
}

func TestExampleLoader_SearchFolderFilters(t *testing.T) {
	basePath := ""
	tempFolderPattern := "ExampleLoaderSearchFolderTests"
	// Create the ExampleLoader and perform the search
	loader := ExampleLoader{
		config:   print.DefaultConfig(),
		examples: nil, // Initialize with an empty slice
	}

	CheckFolderWithExampleDir, _ := os.MkdirTemp(basePath, tempFolderPattern)
	ExampleFolder, _ := os.MkdirTemp(CheckFolderWithExampleDir, tempFolderPattern)
	exampleFilepath := filepath.Join(ExampleFolder, "main.tf")
	os.WriteFile(exampleFilepath, []byte("File Content"), 0777)
	defer os.RemoveAll(CheckFolderWithExampleDir)

	ExampleName := filepath.Base(ExampleFolder)

	tests := []struct {
		name             string
		exampleFolder    string
		included         []string
		excluded         []string
		expectedExamples int
	}{
		{
			name:             "CheckFolderWithExample",
			exampleFolder:    CheckFolderWithExampleDir,
			included:         nil,
			excluded:         nil,
			expectedExamples: 1,
		},
		{
			name:             "CheckFolderWithIncludedExample",
			exampleFolder:    CheckFolderWithExampleDir,
			included:         []string{ExampleName},
			excluded:         nil,
			expectedExamples: 1,
		},
		{
			name:             "CheckFolderWithExcludedExample",
			exampleFolder:    CheckFolderWithExampleDir,
			included:         nil,
			excluded:         []string{ExampleName},
			expectedExamples: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			loader.config.ExamplesFrom = tt.exampleFolder
			loader.config.Examples.Include = tt.included
			loader.config.Examples.Exclude = tt.excluded
			err := loader.SearchFolder()
			if err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}

			assert.Equal(tt.expectedExamples, len(loader.examples))
			loader.examples = nil
		})
	}
}

func TestExampleLoader_GetFile(t *testing.T) {
	basePath := ""
	tempFolderPattern := "ExampleLoaderSearchFolderTests"
	// Create the ExampleLoader and perform the search
	loader := ExampleLoader{
		config:   print.DefaultConfig(),
		examples: nil, // Initialize with an empty slice
	}

	ExamplesDir, _ := os.MkdirTemp(basePath, tempFolderPattern)
	defer os.RemoveAll(ExamplesDir)

	Test1Path := filepath.Join(ExamplesDir, "GetFileTest.tf")
	Test2Path := filepath.Join(ExamplesDir, "GetBlankFileTest.tf")
	Test3Path := filepath.Join(ExamplesDir, "FileDoesNotExist.tf")

	os.WriteFile(Test1Path, []byte("File Content"), 0777)
	os.Create(Test2Path)

	CheckFolderWithoutMaintfDir, _ := os.MkdirTemp(basePath, tempFolderPattern)
	os.MkdirTemp(CheckFolderWithoutMaintfDir, tempFolderPattern)
	defer os.RemoveAll(CheckFolderWithoutMaintfDir)

	tests := []struct {
		name             string
		exampleFile      string
		expectedError    bool
		expectedExamples int
	}{
		{
			name:             "GetFile",
			exampleFile:      Test1Path,
			expectedError:    false,
			expectedExamples: 1,
		},
		{
			name:             "GetBlankFile",
			exampleFile:      Test2Path,
			expectedError:    false,
			expectedExamples: 0,
		},
		{
			name:             "Non-existing file",
			exampleFile:      Test3Path,
			expectedError:    true,
			expectedExamples: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			err := loader.GetFile(tt.exampleFile)

			if tt.expectedError && err == nil {
				t.Error("Expected an error, but got nil")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			}

			assert.Equal(tt.expectedExamples, len(loader.examples))
			loader.examples = nil
		})
	}
}
