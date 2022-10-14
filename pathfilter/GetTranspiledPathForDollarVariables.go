package pathfilter

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func GetTranspiledPathForDollarVariables(
	pathTranspiler *corestr.Hashmap,
	givenPath string,
) string {
	if pathTranspiler != nil && !pathTranspiler.IsEmpty() {
		mappedItems := pathTranspiler.Items()

		for k, v := range mappedItems {
			givenPath = strings.ReplaceAll(givenPath, constants.Dollar+k, v)
		}
	}

	return givenPath
}
