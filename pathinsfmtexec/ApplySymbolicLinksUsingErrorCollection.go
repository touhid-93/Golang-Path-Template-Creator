package pathinsfmtexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/symlink"
)

func ApplySymbolicLinksUsingErrorCollection(
	errorCollection *errwrappers.Collection,
	symLinks *pathinsfmt.SymbolicLinks,
) (isSuccess bool) {
	if symLinks == nil || symLinks.IsEmpty() {
		return true
	}

	errCount := errorCollection.Length()
	if symLinks.IsContinueOnError {
		for _, symLink := range symLinks.SymbolicLinks {
			errWrap := symlink.CreateUsingSymbolicLink(&symLink)

			errorCollection.AddWrapperPtr(errWrap)
		}

		return errorCollection.Length() == errCount
	}

	// immediate exit
	for _, symLink := range symLinks.SymbolicLinks {
		errWrap := symlink.CreateUsingSymbolicLink(&symLink)

		if errWrap.HasError() {
			errorCollection.AddWrapperPtr(errWrap)

			return false
		}
	}

	return errorCollection.Length() == errCount
}
