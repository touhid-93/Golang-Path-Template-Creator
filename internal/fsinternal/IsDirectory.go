package fsinternal

import (
	"os"
)

func IsDirectory(location string) bool {
	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) {
		return false
	}

	return fileInfo != nil && fileInfo.IsDir()
}
