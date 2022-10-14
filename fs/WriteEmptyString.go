package fs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteEmptyString(
	isCreateParentDir bool,
	filePath string,
) *errorwrapper.Wrapper {
	return WriteFile(
		isCreateParentDir,
		filePath,
		[]byte(constants.EmptyString))
}
