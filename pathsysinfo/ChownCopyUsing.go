package pathsysinfo

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func ChownCopyUsing(
	srcFullPath *fileinfopath.Instance,
	dstFullPath string,
) *errorwrapper.Wrapper {
	if srcFullPath.IsInvalidPath() {
		return srcFullPath.ErrorWrapper(errtype.ChownUserOrGroupApplyIssue)
	}

	userGroupId := GetPathUserGroupIdUsing(srcFullPath)

	return userGroupId.ApplyChown(dstFullPath)
}
