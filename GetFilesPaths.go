package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

// GetFileNames
//
// returns file names on the path (non-lazy execution).
func GetFileNames(path string, isNormalize bool) *fileinfo.FileNamesCollection {
	return fileinfo.NewFileNamesUsing(
		path,
		osconsts.PathSeparator,
		isNormalize)
}
