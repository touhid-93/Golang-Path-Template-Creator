package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func SimpleFilter(
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	return SimpleFilterOptions(
		false,
		false,
		filter,
		rootPath)
}
