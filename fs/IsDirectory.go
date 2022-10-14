package fs

import (
	"os"
)

func IsDirectory(location string) bool {
	fileInfo, err := os.Stat(location)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return fileInfo != nil && fileInfo.IsDir()
}
