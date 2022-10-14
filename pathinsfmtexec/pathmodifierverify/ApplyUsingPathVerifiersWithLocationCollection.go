package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingPathVerifiersWithLocationCollection(
	isContinueOnError bool,
	verifiersWithLocations *pathinsfmt.PathVerifiersWithLocationCollection,
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	if verifiersWithLocations == nil ||
		verifiersWithLocations.IsEitherEmpty() {
		return true
	}

	return ApplyUsingFlatPaths(
		isContinueOnError,
		verifiersWithLocations.PathVerifiers,
		errCollection,
		verifiersWithLocations.LocationCollection.LazyFlatPaths())
}
