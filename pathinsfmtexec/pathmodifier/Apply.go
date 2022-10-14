package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(
	isContinueOnErr bool,
	modifier *pathinsfmt.PathModifiersApply,
) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	if modifier == nil ||
		modifier.IsEmptyPathModifiers() ||
		modifier.IsEmptyGenericPathsCollection() {
		return errCollection
	}

	ApplyUsingErrorCollection(
		isContinueOnErr,
		errCollection,
		modifier)

	return errCollection
}
