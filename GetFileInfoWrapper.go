package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

func GetFileInfoWrapper(path string) *fileinfo.Wrapper {
	return fileinfo.New(path, osconsts.PathSeparator)
}
