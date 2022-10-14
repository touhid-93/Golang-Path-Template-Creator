package pathfuncs

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/pathext"
)

type FilterArg struct {
	RootPath, FileName, FullPath string
	IsFile, IsDirectory          bool
	InputError                   error
	os.FileInfo
}

func (receiver *FilterArg) ExtensionWrapper() *pathext.Wrapper {
	return pathext.NewPtr(receiver.RootPath)
}
