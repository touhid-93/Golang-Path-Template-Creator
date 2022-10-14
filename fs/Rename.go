package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func Rename(srcPath, dstPath string) *errorwrapper.Wrapper {
	err := os.Rename(srcPath, dstPath)

	return errnew.
		Path.
		Error(
			errtype.RenamePathFailed,
			err,
			srcPath)
}
