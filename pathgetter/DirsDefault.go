package pathgetter

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func DirsDefault(
	isNormalize bool,
	rootPath string,
) *errstr.Results {
	return Dirs(
		osconsts.PathSeparator,
		rootPath,
		isNormalize)
}
