package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/ispath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// Given removingList array items will be replaced with "" empty string.
// If pathTemplate is given as empty string or nil or whitespace then returns as is.
func RemoveFromPath(
	pathTemplate string,
	removingList *[]string,
	isNormalizePath bool,
) string {
	if ispath.Empty(pathTemplate) {
		return pathTemplate
	}

	for _, value := range *removingList {
		pathTemplate = strings.Replace(
			pathTemplate,
			value,
			constants.EmptyString,
			constants.MinusOne)
	}

	if isNormalizePath {
		pathTemplate = normalize.Path(pathTemplate)
	}

	return pathTemplate
}
