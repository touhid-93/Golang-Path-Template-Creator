package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AppendFile(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return fsinternal.NullContentErrorWrap(
			filePath)
	}

	// file already exist, append
	if IsPathExists(filePath) {
		return appendFileContent(filePath, content)
	}

	return writeNewFileContent(
		isCreateParentDir,
		filePath,
		content)
}
