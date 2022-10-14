package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

// JoinPrefixSuffixIf
//
// only adds prefix if not exist in main (given empty prefix will ignore)
// only adds suffix if not exist in main (given empty suffix will ignore)
func JoinPrefixSuffixIf(
	isFix, isExpand bool,
	prefix, main, suffix string,
) string {
	if prefix == "" && main == "" && suffix == "" {
		return constants.EmptyString
	}

	pathParts := make(
		[]string,
		0,
		constants.Capacity3)

	if prefix != "" && !strings.HasPrefix(main, prefix) {
		pathParts = append(
			pathParts,
			prefix)
	}

	pathParts = append(
		pathParts,
		main)

	if suffix != "" && !strings.HasSuffix(main, suffix) {
		pathParts = append(
			pathParts,
			suffix)
	}

	joined := strings.Join(
		pathParts,
		osconsts.PathSeparator)

	return FixPath(
		isFix,
		isExpand,
		joined)
}
