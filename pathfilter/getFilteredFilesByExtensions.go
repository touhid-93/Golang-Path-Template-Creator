package pathfilter

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/pathhelper/pathext"
)

func getFilteredFilesByExtensions(
	collection *corestr.Collection,
	filter *Query,
) []string {
	return collection.Filter(
		func(currentFilePath string, index int) (
			finalCurrentPath string,
			isKeep bool,
			isBreak bool,
		) {
			currentFileExtWrapper := pathext.NewPtr(currentFilePath)

			if currentFileExtWrapper.IsExtensionFiltersMatch(
				filter.Extensions(),
				filter.ExtensionsLength()) {
				return currentFilePath,
					true,
					false
			}

			return constants.EmptyString,
				false,
				false
		})
}
