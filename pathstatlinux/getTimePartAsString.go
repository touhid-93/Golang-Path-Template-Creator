package pathstatlinux

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

func getTimePartAsString(s string) string {
	leftRight := stringutil.SplitLeftRightType(s, constants.Colon)

	return strings.TrimSpace(leftRight.Right)
}
