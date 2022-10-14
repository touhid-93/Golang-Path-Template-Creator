package fsinternal

import "os"

func IsPathExists(location string) bool {
	_, err := os.Stat(location)

	return !os.IsNotExist(err)
}
