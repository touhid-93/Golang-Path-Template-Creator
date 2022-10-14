package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func FileNameWithExtension(
	fileNameWithoutExt,
	nonDotExt string,
) string {
	return strings.Join(
		[]string{
			fileNameWithoutExt,
			nonDotExt,
		},
		constants.Dot)
}
