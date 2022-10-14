package normalize

import "gitlab.com/evatix-go/core/osconsts"

func PathIf(
	isNormalize bool,
	givenPath string,
) string {
	return PathUsingSeparatorIf(
		false,
		isNormalize,
		isNormalize,
		osconsts.PathSeparator,
		givenPath,
	)
}
