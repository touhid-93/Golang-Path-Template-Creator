package recursivepaths

import (
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func SimpleFilterMust(
	filter pathfuncs.SimpleFilter,
	rootPath string,
) []string {
	results := SimpleFilter(filter, rootPath)
	results.ErrorWrapper.HandleError()

	return results.SafeValues()
}
