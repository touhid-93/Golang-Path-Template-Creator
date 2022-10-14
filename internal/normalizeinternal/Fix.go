package normalizeinternal

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

func Fix(givenPath string) string {
	if len(givenPath) == 0 {
		return constants.EmptyString
	}

	givenPath2 := strings.ReplaceAll(
		givenPath,
		constants.ForwardSlash,
		osconsts.PathSeparator)

	return path.Clean(givenPath2)
}
