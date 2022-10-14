package pathinsfmtexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/symlink"
)

func ApplySymbolicLinks(symLinks *pathinsfmt.SymbolicLinks) *errorwrapper.Wrapper {
	if symLinks == nil || symLinks.IsEmpty() {
		return nil
	}

	if symLinks.IsContinueOnError {
		errCollection := errwrappers.Empty()

		for _, symLink := range symLinks.SymbolicLinks {
			errWrap := symlink.CreateUsingSymbolicLink(&symLink)

			errCollection.AddWrapperPtr(errWrap)
		}

		return errCollection.GetAsErrorWrapperPtr()
	}

	for _, symLink := range symLinks.SymbolicLinks {
		errWrap := symlink.CreateUsingSymbolicLink(&symLink)
		if errWrap.HasError() {
			return errWrap
		}
	}

	return nil
}
