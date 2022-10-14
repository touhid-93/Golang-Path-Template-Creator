package downloadinsexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyMany(
	errCollection *errwrappers.Collection,
	downloads *pathinsfmt.Downloads,
) (isSuccess bool) {
	if downloads.IsEmpty() {
		return true
	}

	if downloads.IsAsyncAll {
		return applyManyAsync(errCollection, downloads)
	}

	stateTracker := errCollection.StateTracker()
	isImmediateExit := !downloads.IsContinueOnError

	for i := range downloads.Downloads {
		err := Apply(&downloads.Downloads[i])
		errCollection.AddWrapperPtr(err)

		if isImmediateExit && err.HasError() {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
