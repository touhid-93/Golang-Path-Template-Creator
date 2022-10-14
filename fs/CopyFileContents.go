package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CopyFileContents(
	srcPath,
	dstPath string,
) (errWrap *errorwrapper.Wrapper) {
	if srcPath == dstPath {
		return nil
	}

	createDirErr := CreateParentDirWithChmodChown(
		srcPath,
		dstPath)

	if createDirErr.HasError() {
		return createDirErr
	}

	copyErr := fsinternal.CopyFileContents(
		srcPath,
		dstPath)

	if copyErr.HasError() {
		return copyErr
	}

	return CopyChmodChown(
		srcPath,
		dstPath)
}
