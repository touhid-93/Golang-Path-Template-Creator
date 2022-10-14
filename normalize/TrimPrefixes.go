package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func TrimPrefixes(
	givenPath string,
	trimPrefixes ...string,
) string {
	if len(trimPrefixes) == 0 || givenPath == constants.EmptyString {
		return givenPath
	}

	for _, prefix := range trimPrefixes {
		givenPath = strings.TrimPrefix(givenPath, prefix)
	}

	return givenPath
}
