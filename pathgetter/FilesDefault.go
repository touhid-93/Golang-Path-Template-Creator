package pathgetter

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func FilesDefault(
	isNormalize bool,
	location string,
) *errstr.Results {
	return Files(
		isNormalize,
		osconsts.PathSeparator,
		location)
}
