package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func SimpleFilterFullPaths(
	filter pathfuncs.SimpleFilter,
	fullPaths ...string,
) *errstr.Results {
	return pathfuncs.SimpleFilterFullPathsAsync(
		false,
		filter,
		fullPaths...)
}
