package normalize

import (
	"gitlab.com/evatix-go/core/osconsts"
)

func PathUsingSingleIf(
	isNormalizeLongPathForce bool,
	givenPath string,
) string {
	return PathUsingSeparatorIf(
		isNormalizeLongPathForce,
		isNormalizeLongPathForce,
		isNormalizeLongPathForce,
		osconsts.PathSeparator,
		givenPath)
}
