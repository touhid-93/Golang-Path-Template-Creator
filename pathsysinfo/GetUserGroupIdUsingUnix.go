package pathsysinfo

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func GetUserGroupIdUsingUnix(fileInfoWithPath *fileinfopath.Instance) *UserGroupId {
	if osconsts.IsWindows {
		return nil
	}

	return GetUserGroupIdUsing(fileInfoWithPath)
}
