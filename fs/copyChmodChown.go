package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CopyChmodChown(
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

	return copyChmodChownInternal(srcPath, dstPath, srcFileInfo)
}
