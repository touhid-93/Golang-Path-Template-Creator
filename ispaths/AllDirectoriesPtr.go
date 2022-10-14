package ispaths

import (
	"gitlab.com/evatix-go/pathhelper/internal/fileinfogetter"
)

func AllDirectoriesPtr(fullPaths *[]string) bool {
	if fullPaths == nil {
		return false
	}

	convertedFileInfoItems := fileinfogetter.Get(
		fullPaths)

	for _, wrapper := range *convertedFileInfoItems {
		isAnyNotDir := wrapper == nil ||
			!wrapper.IsDir()

		if isAnyNotDir {
			return false
		}
	}

	return true
}
