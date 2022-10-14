package pathgetter

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func FilesUsingPaths(
	isNormalize bool,
	separator string,
	exploringPaths ...string,
) *errstr.Results {
	if exploringPaths == nil {
		return errstr.Empty.Results()
	}

	return FilesUsingPathsPtr(
		isNormalize,
		separator,
		exploringPaths)
}
