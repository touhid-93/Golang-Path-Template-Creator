package pathhelper

import "os"

func IsDir(location string) bool {
	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) {
		return false
	}

	return fileInfo != nil && fileInfo.IsDir()
}
