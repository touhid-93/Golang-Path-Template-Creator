package fs

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func copyExistingFile(
	srcPath,
	dstPath string,
	dstFileInfo,
	sourceFileInfo os.FileInfo,
) *errorwrapper.Wrapper {
	if !(dstFileInfo.Mode().IsRegular()) {
		cannotCopyErr := fmt.Errorf(
			"CopyFile: non-regular destination file %s (%q)",
			dstFileInfo.Name(),
			dstFileInfo.Mode().String())

		return errnew.
			Path.
			Error(
				errtype.Copy,
				cannotCopyErr,
				srcPath)
	}

	chownCopyErr := copyChmodChownInternal(
		srcPath,
		dstPath,
		sourceFileInfo)

	if chownCopyErr.HasError() {
		return chownCopyErr
	}

	if os.SameFile(sourceFileInfo, dstFileInfo) {
		return nil
	}

	linkErr := os.Link(srcPath, dstPath)

	if linkErr != nil {
		return errnew.
			Path.
			Messages(
				errtype.PathCopy,
				linkErr.Error(),
				"Source:"+srcPath+", Destination:"+dstPath)
	}

	return nil
}
