package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteFileString(
	dirMode, fileMode os.FileMode,
	filePath string,
	contentString string,
) *errorwrapper.Wrapper {
	return WriteFile(
		dirMode,
		fileMode,
		filePath,
		[]byte(contentString))
}
