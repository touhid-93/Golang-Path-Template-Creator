package fs

import "os"

func IsNotPathExistsUsing(fileInfo os.FileInfo, err error) bool {
	return fileInfo == nil || err != nil && os.IsNotExist(err)
}
