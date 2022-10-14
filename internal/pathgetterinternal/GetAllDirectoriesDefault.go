package pathgetterinternal

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func GetAllDirectoriesDefault(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllDirectories(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
