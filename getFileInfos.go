package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

func GetFileInfoWrappersFrom(path string, isNormalize bool) *fileinfo.Wrappers {
	return fileinfo.NewWrappersPtr(
		path,
		osconsts.PathSeparator,
		isNormalize)
}
