package pathchmod

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func GetUserGroupId(userName string, groupName string) (
	userId int,
	groupId int,
	errWrapper *errorwrapper.Wrapper,
) {
	userInfo := pathsysinfo.GetUserInfo(userName)

	if userInfo.ErrorWrapper.HasError() {
		return 0, 0, userInfo.ErrorWrapper
	}

	groupInfo := pathsysinfo.GetGroupInfo(groupName)

	if groupInfo.ErrorWrapper.HasError() {
		return userInfo.Id, 0, groupInfo.ErrorWrapper
	}

	return userInfo.Id,
		groupInfo.Id,
		nil
}
