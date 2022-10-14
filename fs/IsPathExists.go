package fs

import "os"

func IsPathExists(location string) bool {
	fileInfo, err := os.Stat(location)

	return err == nil && fileInfo != nil ||
		err != nil && !os.IsNotExist(err)
}
