package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

// SimpleFilterNonRecursiveFiles normalize, expand false
func SimpleFilterNonRecursiveFiles(
	simpleFilter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	return SimpleFilterFilesOptionsAsync(
		false,
		false,
		false,
		simpleFilter,
		rootPath)
}
