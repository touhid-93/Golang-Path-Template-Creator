package pathscreateinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
)

func applyUserNameGroupNameOnPathCreators(
	pathsCreator *pathinsfmt.PathsCreator,
	errorCollection *errwrappers.Collection,
) *errorwrapper.Wrapper {
	if pathsCreator == nil || pathsCreator.ApplyUserGroup == nil {
		return nil
	}

	errWrap := namegroup.Apply(
		true,
		false,
		pathsCreator.ApplyUserGroup,
		pathsCreator.RootDir)

	errorCollection.AddWrapperPtr(errWrap)

	return errWrap
}
