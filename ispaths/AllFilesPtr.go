package ispaths

import (
	"gitlab.com/evatix-go/pathhelper/internal/fileinfogetter"
)

func AllFilesPtr(fullPaths *[]string) bool {
	if fullPaths == nil {
		return false
	}

	convertedFileInfoItems := fileinfogetter.Get(
		fullPaths)

	for _, wrapper := range *convertedFileInfoItems {
		isAnyInvalidOrDir := wrapper == nil ||
			wrapper.IsDir()

		if isAnyInvalidOrDir {
			return false
		}
	}

	return true
}
