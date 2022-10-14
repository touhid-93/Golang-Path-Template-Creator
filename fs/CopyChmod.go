package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CopyChmod(
	srcPath,
	dstPath string,
) *errorwrapper.Wrapper {
	srcFileInfo, err := os.Stat(srcPath)

	if IsNotPathExistsUsing(srcFileInfo, err) {
		return errnew.SrcDst.Error(
			errtype.ChownUserOrGroupApplyIssue,
			err,
			srcPath,
			dstPath,
		)
	}

	chmodApplyErr := os.Chmod(dstPath, srcFileInfo.Mode())

	if chmodApplyErr != nil {
		return errnew.SrcDst.Error(
			errtype.ChmodApplyFailed,
			err,
			srcPath,
			dstPath,
		)
	}

	return nil
}
