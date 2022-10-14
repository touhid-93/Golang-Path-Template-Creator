package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyBaseUsingErrorCollection(
	isContinueOnErr bool,
	errorCollection *errwrappers.Collection,
	modifier *pathinsfmt.BasePathModifiersApply,
) (isSuccess bool) {
	if modifier == nil ||
		modifier.PathModifiersApply == nil ||
		modifier.PathModifiersApply.IsEmptyPathModifiers() ||
		modifier.PathModifiersApply.IsEmptyGenericPathsCollection() {
		return true
	}

	return ApplyUsingErrorCollection(
		isContinueOnErr,
		errorCollection,
		modifier.PathModifiersApply)
}
