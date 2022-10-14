package ispath

import "os"

func Exists(location string) bool {
	fileInfo, err := os.Stat(location)

	return err == nil && fileInfo != nil || fileInfo != nil && !os.IsNotExist(err)
}
