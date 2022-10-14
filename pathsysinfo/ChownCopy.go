package pathsysinfo

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func ChownCopy(srcFullPath, dstFullPath string) *errorwrapper.Wrapper {
	srcInstance := fileinfopath.New(srcFullPath)

	return ChownCopyUsing(srcInstance, dstFullPath)
}
