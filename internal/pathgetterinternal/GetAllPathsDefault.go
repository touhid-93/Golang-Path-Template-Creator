package pathgetterinternal

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func GetAllPathsDefault(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllPaths(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
