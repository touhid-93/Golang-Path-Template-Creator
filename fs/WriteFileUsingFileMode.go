package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func WriteFileUsingFileMode(
	isCreateParentDir,
	isApplyChmodOnMismatch bool, // or else always apply
	isKeepExistingFileModeOnExist bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return fsinternal.NullContentErrorWrap(
			filePath)
	}

	isExist := IsPathExists(filePath)

	// file already exist
	if isExist && isKeepExistingFileModeOnExist {
		return writeExistingFileContent(
			filePath,
			content)
	} else if isExist && !isKeepExistingFileModeOnExist {
		return writeExistingFileContentUsingFileMode(
			isApplyChmodOnMismatch,
			filePath,
			content,
			fileMode)
	}

	// new content
	return writeNewFileContentUsingFileMode(
		isCreateParentDir,
		true,
		isApplyChmodOnMismatch,
		dirMode,
		fileMode,
		filePath,
		content,
	)
}
