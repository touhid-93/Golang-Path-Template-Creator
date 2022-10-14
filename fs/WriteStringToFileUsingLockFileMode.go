package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringToFileUsingLockFileMode(
	isCreateParentDir,
	isApplyChmodOnMismatch bool, // or else always apply
	isKeepExistingFileModeOnExist bool,
	dirChmod, fileChmod os.FileMode,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	return WriteFileUsingFileMode(
		isCreateParentDir,
		isApplyChmodOnMismatch,
		isKeepExistingFileModeOnExist,
		dirChmod,
		fileChmod,
		filePath,
		[]byte(content),
	)
}
