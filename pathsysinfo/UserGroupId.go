package pathsysinfo

import (
	"os"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type UserGroupId struct {
	UserId, GroupId int
	Error           string
}

func (it *UserGroupId) HasError() bool {
	return it != nil && it.Error != ""
}

func (it *UserGroupId) IsEmptyError() bool {
	return it == nil || it.Error == ""
}

func (it *UserGroupId) HasValidUserId() bool {
	return it != nil && it.Error != "" && it.UserId > constants.InvalidValue
}

func (it *UserGroupId) HasValidGroupId() bool {
	return it != nil && it.Error != "" && it.GroupId > constants.InvalidValue
}

func (it *UserGroupId) HasValidUserAndGroupId() bool {
	return it != nil &&
		it.Error != "" &&
		it.UserId > constants.InvalidValue &&
		it.GroupId > constants.InvalidValue
}

func (it *UserGroupId) IsInvalidUserOrGroupId() bool {
	return it == nil ||
		it.Error != "" ||
		it.UserId == constants.InvalidValue ||
		it.GroupId == constants.InvalidValue
}

func (it *UserGroupId) InvalidError(fullPath string) *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.WithRefs(
			"",
			it,
			ref.Value{
				Variable: "fullPath",
				Value:    fullPath,
			})
	}

	if it.HasError() {
		return it.ErrorWrapper(fullPath)
	}

	if it.IsInvalidUserOrGroupId() {
		return errnew.Messages.Many(
			errtype.ChownUserOrGroupApplyIssue,
			"either user or group id is invalid for path : ",
			fullPath,
			"UserId",
			strconv.Itoa(it.UserId),
			"GroupId",
			strconv.Itoa(it.GroupId),
		)
	}

	return nil
}

func (it *UserGroupId) ErrorWrapper(fullPath string) *errorwrapper.Wrapper {
	if it.HasError() {
		return errnew.Messages.Many(
			errtype.ChownUserOrGroupApplyIssue,
			it.Error,
			"UserId",
			strconv.Itoa(it.UserId),
			"GroupId",
			strconv.Itoa(it.GroupId),
			"path:",
			fullPath,
		)
	}

	return nil
}

func (it *UserGroupId) ApplyChown(fullPath string) *errorwrapper.Wrapper {
	if it.IsInvalidUserOrGroupId() {
		return it.InvalidError(fullPath)
	}

	err := os.Chown(fullPath, it.UserId, it.GroupId)

	if err == nil {
		return nil
	}

	// has error
	return errnew.Path.Error(
		errtype.ChownUserOrGroupApplyIssue,
		err,
		fullPath,
	)
}

func (it *UserGroupId) ApplyUserId(fullPath string) *errorwrapper.Wrapper {
	if it.IsInvalidUserOrGroupId() {
		return it.InvalidError(fullPath)
	}

	applyPathUserInfo := GetPathUserGroupId(fullPath)

	if applyPathUserInfo.IsInvalidUserOrGroupId() {
		return applyPathUserInfo.InvalidError()
	}

	err := os.Chown(
		fullPath,
		it.UserId,
		applyPathUserInfo.GroupId)

	if err != nil {
		return errnew.
			Path.
			Error(
				errtype.ChownUserOrGroupApplyIssue,
				err,
				fullPath)
	}

	return nil
}

func (it *UserGroupId) ApplyGroupId(fullPath string) *errorwrapper.Wrapper {
	if it.IsInvalidUserOrGroupId() {
		return it.InvalidError(fullPath)
	}

	applyPathUserInfo := GetPathUserGroupId(fullPath)

	if applyPathUserInfo.IsInvalidUserOrGroupId() {
		return applyPathUserInfo.InvalidError()
	}

	err := os.Chown(
		fullPath,
		applyPathUserInfo.UserId,
		it.GroupId)

	if err != nil {
		return errnew.
			Path.
			Error(
				errtype.ChownUserOrGroupApplyIssue,
				err,
				fullPath)
	}

	return nil
}

func (it *UserGroupId) IsEqual(right *UserGroupId) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it == right {
		return true
	}

	if it.UserId != right.UserId {
		return false
	}

	if it.GroupId != right.GroupId {
		return false
	}

	return it.Error == right.Error
}

func (it *UserGroupId) String() string {
	if it == nil {
		return constants.NilAngelBracket
	}

	if it.HasError() {
		return errcore.VarThreeNoType(
			"UserId", it.UserId,
			"GroupId", it.GroupId,
			"Error", it.Error)
	}

	return errcore.VarTwoNoType(
		"UserId", it.UserId,
		"GroupId", it.GroupId,
	)
}
