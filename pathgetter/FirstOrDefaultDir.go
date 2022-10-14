package pathgetter

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func FirstOrDefaultDir(
	isNormalize bool,
	rootPath string,
) *errstr.Result {
	results := Dirs(
		osconsts.PathSeparator,
		rootPath,
		isNormalize)

	return results.FirstOrDefaultResult()
}
