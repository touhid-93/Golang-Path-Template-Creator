package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

// SimpleFilterNonRecursiveDirs normalize, expand false
func SimpleFilterNonRecursiveDirs(
	simpleFilter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	return SimpleFilterDirsOptions(
		false,
		false,
		false,
		simpleFilter,
		rootPath)
}
