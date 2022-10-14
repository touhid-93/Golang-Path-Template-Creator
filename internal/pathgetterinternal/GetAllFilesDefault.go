package pathgetterinternal

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

// GetAllFilesDefault only gives files not nested files
func GetAllFilesDefault(isFixPaths bool, rootPath string) *errstr.Results {
	return GetAllFiles(
		isFixPaths,
		osconsts.PathSeparator,
		rootPath)
}
