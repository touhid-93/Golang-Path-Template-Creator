package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringToFile(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	return WriteFile(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
