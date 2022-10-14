package normalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

// getLongPathFixedUsingSeparator
// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix ) or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
func getLongPathFixedUsingSeparator(
	isForce bool,
	separator string,
	givenAbsolutePath string,
) string {
	isIgnoreCase :=
		givenAbsolutePath == "" ||
			len(givenAbsolutePath) < 255

	isWindowsSeparator := separator ==
		constants.WindowsPathSeparator

	if !isWindowsSeparator {
		return givenAbsolutePath
	}

	isForce = isForce && isWindowsSeparator

	if (isIgnoreCase && !isForce) || osconsts.IsUnixGroup {
		return givenAbsolutePath
	}

	return constants.LongPathQuestionMarkPrefix + givenAbsolutePath
}
