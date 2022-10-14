package fs

import "os"

func IsPathExistsUsing(
	fileInfo os.FileInfo,
	err error,
) bool {
	return err == nil &&
		fileInfo != nil
}
