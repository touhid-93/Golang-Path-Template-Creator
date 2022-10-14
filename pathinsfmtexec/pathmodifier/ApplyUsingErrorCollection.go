package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingErrorCollection(
	isContinueOnErr bool,
	errorCollection *errwrappers.Collection,
	modifier *pathinsfmt.PathModifiersApply,
) (isSuccess bool) {
	if modifier == nil ||
		modifier.IsEmptyPathModifiers() ||
		modifier.IsEmptyGenericPathsCollection() {
		return true
	}

	flatPaths := modifier.
		GenericPathsCollection.
		LazyFlatPaths()

	stateTracker := errorCollection.StateTracker()

	for _, pathModifier := range modifier.PathModifiers {
		ApplySimple(
			isContinueOnErr,
			errorCollection,
			&pathModifier,
			flatPaths)
	}

	return stateTracker.IsSuccess()
}
