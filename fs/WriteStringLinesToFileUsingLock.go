package fs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringLinesToFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	contentLines []string,
) *errorwrapper.Wrapper {
	if len(contentLines) == 0 {
		return WriteEmptyStringLock(
			isCreateParentDir,
			filePath)
	}

	content := strings.Join(
		contentLines,
		constants.NewLineUnix)

	return WriteFileLock(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
