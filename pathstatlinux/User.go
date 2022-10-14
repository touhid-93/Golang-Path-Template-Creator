package pathstatlinux

import (
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

type User struct {
	IntIdNameValidation
	systemUser *pathsysinfo.UserInfo
}

func (u *User) SystemUser() *pathsysinfo.UserInfo {
	if u.systemUser != nil {
		return u.systemUser
	}

	if !u.HasValidId {
		u.systemUser = invalidUserInfo

		return u.systemUser
	}

	u.systemUser = pathsysinfo.GetUserInfo(u.Name)

	return u.systemUser
}
