package pathsysinfo

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
)

func ChownCopyUnix(srcFullPath, dstFullPath string) *errorwrapper.Wrapper {
	if osconsts.IsWindows {
		return nil
	}

	return ChownCopy(srcFullPath, dstFullPath)
}
