package pathstatlinux

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/pathhelper/internal/strremove"
)

// parentThesisWrappedSlashLeftRight : `(0755/drwxr-xr-x)  Uid` or `(left/right) whatever` data to left : 0755, right : drwxr-xr-x
func getLeftRightOfBracketsPatternData(parentThesisWrappedSlashLeftRight string) corestr.LeftRight {
	findSubStringsByRegex := bracketsMatcherWithContents.FindStringSubmatch(
		parentThesisWrappedSlashLeftRight)

	parentThesisPatternedData := stringslice.FirstOrDefault(findSubStringsByRegex)
	splitItems := strremove.SimpleManySplitsBy(
		parentThesisPatternedData,
		constants.ForwardSlash,
		constants.ParenthesisStart,
		constants.ParenthesisEnd)

	length := len(splitItems)

	if length == 2 {
		return corestr.LeftRight{
			Left:    strings.TrimSpace(splitItems[coreindexes.First]),
			Right:   strings.TrimSpace(splitItems[coreindexes.Second]),
			IsValid: true,
			Message: "",
		}
	}

	first, last := stringslice.FirstLastDefault(splitItems)

	return corestr.LeftRight{
		Left:    strings.TrimSpace(first),
		Right:   strings.TrimSpace(last),
		IsValid: false,
		Message: errcore.Expecting("Expected length", 2, length),
	}
}
