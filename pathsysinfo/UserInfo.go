package pathsysinfo

import (
	"fmt"
	"os/user"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
)

type UserInfo struct {
	*user.User
	Id           int
	IsValidUser  bool
	HasValidId   bool
	ErrorWrapper *errorwrapper.Wrapper
}

func InvalidUserInfo(errorWrapper *errorwrapper.Wrapper) *UserInfo {
	return &UserInfo{
		Id:           constants.InvalidValue,
		IsValidUser:  false,
		HasValidId:   false,
		ErrorWrapper: errorWrapper,
	}
}

func (it *UserInfo) String() string {
	return fmt.Sprintf(printFormat,
		it.Name,
		it.Id,
		it.IsValidUser,
		it.HasValidId)
}
