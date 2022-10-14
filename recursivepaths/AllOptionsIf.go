package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func AllOptionsIf(
	isRecursive bool,
	isNormalize,
	isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	if !isRecursive {
		return ReturnAsIs(
			isNormalize,
			isExpandEnv,
			rootPath)
	}

	return AllOptions(
		isNormalize,
		isExpandEnv,
		rootPath)
}
