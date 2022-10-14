package fs

import (
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/createdir"
)

func CreateParentDirWithChmodChown(srcPath string, dstPath string) *errorwrapper.Wrapper {
	srcDir := filepath.Dir(srcPath)
	srcBaseDirInfo, err := os.Stat(srcDir)

	if IsNotPathExistsUsing(srcBaseDirInfo, err) {
		return errnew.
			Path.
			Error(
				errtype.Copy,
				err,
				srcDir)
	}

	dstParentDir := filepath.Dir(dstPath)
	createDirErr := createdir.AllOnNonExist(
		dstParentDir,
		srcBaseDirInfo.Mode())

	if createDirErr.HasError() {
		return createDirErr
	}

	return CopyChmodChown(srcDir, dstParentDir)
}
