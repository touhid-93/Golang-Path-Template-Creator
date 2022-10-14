package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func ReturnAsIs(
	isNormalize bool,
	isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	fixedPath := pathjoin.FixPath(
		isNormalize,
		isExpandEnv,
		rootPath)

	return errstr.New.Results.SpreadValuesOnly(
		fixedPath)
}
