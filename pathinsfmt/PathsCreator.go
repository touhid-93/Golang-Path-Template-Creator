package pathinsfmt

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
)

type PathsCreator struct {
	BasePathsCreator
	ApplyRwx       *chmodins.RwxOwnerGroupOther
	ApplyUserGroup *UserGroupName
}

func (it *PathsCreator) HasRwx() bool {
	return it.ApplyRwx != nil
}

func (it *PathsCreator) HasUserGroup() bool {
	return it.ApplyUserGroup != nil &&
		it.ApplyUserGroup.HasUserNameOrGroup()
}
