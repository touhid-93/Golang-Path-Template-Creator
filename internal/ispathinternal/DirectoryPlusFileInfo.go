package ispathinternal

import (
	"os"
)

func DirectoryPlusFileInfo(location string) (isSuccess bool, fileInfo os.FileInfo) {
	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) {
		return false, fileInfo
	}

	return fileInfo.IsDir(), fileInfo
}
