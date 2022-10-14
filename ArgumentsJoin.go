package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func ArgumentsJoin(args ...string) string {
	if len(args) == 0 {
		return constants.EmptyString
	}

	return strings.Join(args, constants.Space)
}
