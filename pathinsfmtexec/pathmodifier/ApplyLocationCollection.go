package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyLocationCollection(
	isContinueOnErr bool,
	errorCollection *errwrappers.Collection,
	modifier *pathinsfmt.PathModifier,
	locationCollection *pathinsfmt.LocationCollection,
) (isSuccess bool) {
	if modifier == nil ||
		locationCollection == nil ||
		locationCollection.IsEmpty() {
		return true
	}

	return ApplySimple(
		isContinueOnErr,
		errorCollection,
		modifier,
		locationCollection.LazyFlatPaths())
}
