package pathinsfmt

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

type UserGroupName struct {
	BaseGroupName
	UserName string `json:"UserName,omitempty"` // Not define or empty string or * means keeping the existing one
}

func NewUserGroupName(
	username, groupName string,
) *UserGroupName {
	return &UserGroupName{
		BaseGroupName: BaseGroupName{
			GroupName: groupName,
		},
		UserName: username,
	}
}

func (it *UserGroupName) HasUserNameOrGroup() bool {
	return it.HasUserName() || it.HasGroupName()
}

func (it *UserGroupName) HasUserName() bool {
	return it != nil && !stringutil.IsEmpty(it.UserName)
}

func (it *UserGroupName) IsUserNameEmpty() bool {
	return it == nil || stringutil.IsEmpty(it.UserName)
}

func (it *UserGroupName) UserNameSimple() string {
	return it.UserName
}

func (it *UserGroupName) IsUsername(checkingUserName string) bool {
	isUsernameExist := it.HasUserName()

	if !isUsernameExist && checkingUserName == "" {
		return true
	}

	if !isUsernameExist {
		return false
	}

	return checkingUserName == it.UserName
}

func (it *UserGroupName) IsGroupNameUserNameBothEmpty() bool {
	return it == nil || it.IsUserNameEmpty() && it.IsGroupNameEmpty()
}

func (it *UserGroupName) Clone() *UserGroupName {
	if it == nil {
		return nil
	}

	return &UserGroupName{
		BaseGroupName: BaseGroupName{
			GroupName: it.GroupName,
		},
		UserName: it.UserName,
	}
}
