package pathhelper

import "gitlab.com/evatix-go/pathhelper/fileinfo"

func GetFileInfoWrappersFromUsingSeparator(path, separator string, isNormalize bool) *fileinfo.Wrappers {
	return fileinfo.NewWrappersPtr(
		path,
		separator,
		isNormalize)
}
