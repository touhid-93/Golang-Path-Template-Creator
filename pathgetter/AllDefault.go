package pathgetter

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func AllDefault(
	isNormalize bool,
	exploringPaths ...string,
) *errstr.Results {
	length := len(exploringPaths)

	if length == 0 {
		return errstr.Empty.Results()
	}

	if length == 1 {
		return AllOfSinglePath(
			isNormalize,
			osconsts.PathSeparator,
			exploringPaths[0])
	}

	return All(
		osconsts.PathSeparator,
		isNormalize,
		exploringPaths)
}
