package symlink

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func IsSymLink(path string) *errbool.Result {
	info, err := os.Lstat(path)

	if err != nil {
		return &errbool.Result{
			Value: false,
			ErrorWrapper: errnew.Messages.Many(
				errtype.SymbolicLink,
				path,
				err.Error()),
		}
	}

	isSuccess := info.Mode()&os.ModeSymlink != 0

	return &errbool.Result{
		Value:        isSuccess,
		ErrorWrapper: nil,
	}
}
