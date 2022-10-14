package pathgetter

import (
	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

func AllInfoWrapperPtr(
	separator string,
	isNormalize bool,
	exploringPaths []string,
) []*fileinfo.Wrappers {
	length := len(exploringPaths)
	wrappersOfWrappers := make([]*fileinfo.Wrappers, length)

	if length == 0 {
		return wrappersOfWrappers
	}

	for i, expPath := range exploringPaths {
		wrappers := AllOfSinglePathInfoWrappers(
			separator,
			isNormalize,
			expPath)

		wrappersOfWrappers[i] = wrappers
	}

	return wrappersOfWrappers
}
