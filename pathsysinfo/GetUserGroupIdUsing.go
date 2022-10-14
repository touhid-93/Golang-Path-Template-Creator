package pathsysinfo

import (
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func GetUserGroupIdUsing(fileInfoWithPath *fileinfopath.Instance) *UserGroupId {
	instance := GetPathUserGroupIdUsing(fileInfoWithPath)

	if instance != nil {
		return instance.UserGroupId()
	}

	return nil
}
