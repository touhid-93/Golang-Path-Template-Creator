package normalizeinternal

import (
	"path"

	"gitlab.com/evatix-go/core/constants"
)

func JoinPathsFixIf(isFix bool, givenPaths ...string) string {
	if len(givenPaths) == 0 {
		return constants.EmptyString
	}

	joinedPath := path.Join(givenPaths...)

	return FixIfSingle(isFix, joinedPath)
}
