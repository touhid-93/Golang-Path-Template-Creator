package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/createdir"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func WriteAllParams(
	isCreateParentDir,
	isApplyChmodMust,
	isApplyChmodOnMismatchOnly,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirCreateChmod os.FileMode,
	fileChmod os.FileMode,
	filePath string,
	contents []byte,
) *errorwrapper.Wrapper {
	if isSkipOnNilObject && contents == nil {
		return nil
	}

	if contents == nil {
		return fsinternal.NullContentErrorWrap(
			filePath)
	}

	isExist := IsPathExists(filePath)
	// file already exist
	if isExist && isKeepExistingFileModeOnExist {
		return writeExistingFileContent(
			filePath,
			contents)
	} else if isExist && !isKeepExistingFileModeOnExist {
		return writeExistingFileContentUsingFileMode(
			isApplyChmodOnMismatchOnly,
			filePath,
			contents,
			fileChmod)
	}

	// new
	var createDirErr *errorwrapper.Wrapper
	if isCreateParentDir {
		createDirErr = createdir.AllUptoParent(
			filePath,
			dirCreateChmod)
	}

	if createDirErr.HasError() {
		return createDirErr
	}

	return writeNewFileContentUsingFileMode(
		isCreateParentDir,
		isApplyChmodMust,
		isApplyChmodOnMismatchOnly,
		dirCreateChmod,
		fileChmod,
		filePath,
		contents,
	)
}
