package fsinternal

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func GetDirFileName(location string) (parentDir, fileName string) {
	return splitinternal.GetWithoutSlash(location)
}
