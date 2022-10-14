package fs

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
)

// CopyFile Future ref: https://stackoverflow.com/a/21067803
func CopyFile(srcPath, dstPath string) *errorwrapper.Wrapper {
	if srcPath == dstPath {
		return nil
	}

	sourceFileInfo, err := os.Stat(srcPath)
	if IsNotPathExistsUsing(sourceFileInfo, err) {
		return errnew.
			Path.
			Error(
				errtype.PathStatFailed,
				err,
				srcPath)
	}

	if !sourceFileInfo.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		cannotCopySymLinkErr := fmt.Errorf(
			"CopyFile: non-regular source file %s (%q)",
			sourceFileInfo.Name(),
			sourceFileInfo.Mode().String())

		return errnew.SrcDst.Error(
			errtype.Copy,
			cannotCopySymLinkErr,
			srcPath,
			dstPath,
		)
	}

	if sourceFileInfo.IsDir() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		cannotCopyDirErr := fmt.Errorf(
			"CopyFile: don't support dir copy %s (%q)",
			sourceFileInfo.Name(),
			srcPath)

		return errnew.SrcDst.Error(
			errtype.Copy,
			cannotCopyDirErr,
			srcPath,
			dstPath,
		)
	}

	dstFileInfo, dstErr := os.Stat(dstPath)
	isExist := IsPathExistsUsing(dstFileInfo, dstErr)
	if isExist && !dstFileInfo.IsDir() {
		return deletepaths.Recursive(dstPath)
	} else if isExist && dstFileInfo.IsDir() {
		return errnew.SrcDst.Messages(
			errtype.PathCopy,
			srcPath,
			dstPath,
			"don't support copy dir on file copier. destination contains same file name dir.")
	}

	// copy new file
	return CopyFileContents(srcPath, dstPath)
}
