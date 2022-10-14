package pathsysinfo

import (
	"gitlab.com/evatix-go/core/constants"
)

func GetUserInfo(userName string) *UserInfo {
	userResult, errorWrapper := LookupUser(userName)

	if errorWrapper.HasError() || userResult == nil {
		return InvalidUserInfo(errorWrapper)
	}

	id, convertErr := SystemUserId(userResult)

	if convertErr.HasError() || id == constants.InvalidValue {
		return &UserInfo{
			User:         userResult,
			Id:           id,
			IsValidUser:  true,
			HasValidId:   false,
			ErrorWrapper: convertErr,
		}
	}

	return &UserInfo{
		User:         userResult,
		Id:           id,
		IsValidUser:  true,
		HasValidId:   true,
		ErrorWrapper: nil,
	}
}
