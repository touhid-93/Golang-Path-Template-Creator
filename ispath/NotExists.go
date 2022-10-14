package ispath

import "os"

func NotExists(path string) bool {
	fileInfo, err := os.Stat(path)

	return err != nil &&
		os.IsNotExist(err) ||
		fileInfo == nil
}
