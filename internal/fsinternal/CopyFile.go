package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

// CopyFile Future ref: https://stackoverflow.com/a/21067803
func CopyFile(srcPath, dstPath string, defaultChmod os.FileMode) *errorwrapper.Wrapper {
	erWp := CreateDirectoryAllUptoParent(dstPath, defaultChmod)

	if erWp.HasError() {
		return erWp
	}

	return CopyFileContents(srcPath, dstPath)
}
