package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

type BaseGroupName struct {
	GroupName string `json:"GroupName,omitempty"` // Not define or empty string or * means keeping the existing one
}

func (it *BaseGroupName) Clone() *BaseGroupName {
	if it == nil {
		return nil
	}

	return &BaseGroupName{
		GroupName: it.GroupName,
	}
}

func (it *BaseGroupName) IsGroupNameEmpty() bool {
	return it == nil || stringutil.IsEmptyOrWhitespace(it.GroupName)
}

func (it *BaseGroupName) IsGroupName(checkingGroupName string) bool {
	return checkingGroupName == it.GroupName
}

func (it *BaseGroupName) HasGroupName() bool {
	return it != nil && it.GroupName != constants.EmptyString
}
