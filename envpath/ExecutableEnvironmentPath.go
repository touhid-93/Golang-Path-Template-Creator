package envpath

import (
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

type ExecutableEnvironmentPath struct {
	Variable         string
	Expanded         string
	pathsCollections *fileinfo.FileNamesCollection
	fileWrappers     *fileinfo.Wrappers
	sync.Mutex
}

// GetCachedFileNamesCollection returns all pathsCollection paths on that env directory once,
// caches it and returns that in later function calls
func (it *ExecutableEnvironmentPath) GetCachedFileNamesCollection() *fileinfo.FileNamesCollection {
	// checking if fileInfos already generated
	if it.pathsCollections != nil {
		return it.pathsCollections
	}

	it.pathsCollections = it.GetFileNamesCollection()

	return it.pathsCollections
}

func (it *ExecutableEnvironmentPath) Length() int {
	return it.GetCachedFileNamesCollection().Length()
}

func (it *ExecutableEnvironmentPath) GetFileNamesCollection() *fileinfo.FileNamesCollection {
	return fileinfo.NewFileNamesUsing(
		it.Expanded,
		osconsts.PathSeparator,
		true)
}

// GetDirectories
//
// returns all directories paths on that env directory,
func (it *ExecutableEnvironmentPath) GetDirectories() []*string {
	var directories []*string

	arrayFromPath := strings.Split(it.Expanded, constants.PathSeparator)

	for _, arrayItem := range arrayFromPath {
		directories = append(directories, &arrayItem)
	}

	return directories
}

// GetFilePathsContains
//
// returns all pathsCollection paths on which contains the given string.
// If no path is found, returns empty array.
func (it *ExecutableEnvironmentPath) GetFilePathsContains(
	separator,
	contains string,
) *[]string {
	var filePathThatContains = make([]string, 0, it.Length())
	filesPaths := it.GetCachedFileNamesCollection().GetFilePaths(separator)

	for _, eachPath := range filesPaths {
		if strings.Contains(eachPath, contains) {
			filePathThatContains = append(filePathThatContains, eachPath)
		}
	}

	return &filePathThatContains
}
