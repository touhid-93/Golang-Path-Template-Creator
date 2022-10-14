package pathfilter

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/pathhelper/pathext"
)

func getFilteredFilesByNameMatching(
	collection *corestr.Collection,
	eachPathExtWrapper *pathext.Wrapper,
) []string {
	resultsBySpecificNameFilter := collection.Filter(
		func(str string, index int) (
			result string,
			isKeep bool,
			isBreak bool,
		) {
			isEndsWith := eachPathExtWrapper.
				IsNameWithExtensionMatches(
					str)

			return str,
				isEndsWith,
				isEndsWith
		})

	return resultsBySpecificNameFilter
}
