package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyGenericPathCollection(
	isContinueOnErr bool,
	errorCollection *errwrappers.Collection,
	modifier *pathinsfmt.PathModifier,
	genericPathsCollection *pathinsfmt.GenericPathsCollection,
) (isSuccess bool) {
	if modifier == nil ||
		genericPathsCollection == nil ||
		genericPathsCollection.IsEmpty() {
		return true
	}

	return ApplySimple(
		isContinueOnErr,
		errorCollection,
		modifier,
		genericPathsCollection.LazyFlatPaths())
}
