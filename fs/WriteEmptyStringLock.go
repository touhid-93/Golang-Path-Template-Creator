package fs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteEmptyStringLock(
	isCreateParentDir bool,
	filePath string,
) *errorwrapper.Wrapper {
	return WriteFileLock(
		isCreateParentDir,
		filePath,
		[]byte(constants.EmptyString))
}
