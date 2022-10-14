package pathsysinfo

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func GetPathUserGroupIdUsing(instance *fileinfopath.Instance) *PathUserGroupId {
	return &PathUserGroupId{
		FileInfoWithPath: instance,
		UserId:           constants.InvalidValue,
		GroupId:          constants.InvalidValue,
		Error: errtype.NotSupportInWindows.Error(
			constants.EmptyString,
			"path",
			instance.FullPath),
	}
}
