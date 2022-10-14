package pathscreateinsexec

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
)

func ApplyPathsCreatorUsingErrorCollection(
	isLock,
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	errorCollection *errwrappers.Collection,
	pathsCreator *pathinsfmt.PathsCreator,
) (
	isSuccess bool,
) {
	if pathsCreator == nil {
		return true
	}

	workingPaths := pathsCreator.
		LazyFlatPathsIf(isLazyPaths)

	errCount := errorCollection.Length()

	if isDeleteAllBeforeCreate {
		errorCollection.AddAnyFunctions(
			pathsCreator.DeleteAllPaths)
	}

	// paths create
	if pathsCreator.HasRwx() {
		fileMode, errWrap := pathchmod.ParseRwxOwnerGroupOtherToFileMode(
			pathsCreator.ApplyRwx)

		errorCollection.AddWrapperPtr(errWrap)

		if errWrap.HasError() {
			return false
		}

		// create using chmod
		_, filesCreateErr := createpath.CreateManySameDirWithFileMode(
			isLock,
			isIgnoreOnExist,
			fileMode,
			pathsCreator.RootDir,
			workingPaths)

		errorCollection.AddWrapperPtr(filesCreateErr)

		if filesCreateErr.HasError() {
			return false
		}
	} else {
		// create without chmod
		_, filesCreateErr := createpath.CreateMany(
			isLock,
			isIgnoreOnExist,
			workingPaths)

		errorCollection.AddWrapperPtr(filesCreateErr)

		if filesCreateErr.HasError() {
			return false
		}
	}

	// apply groups
	if pathsCreator.HasUserGroup() && osconsts.IsUnixGroup {
		errWrap := namegroup.Apply(
			true,
			false,
			pathsCreator.ApplyUserGroup,
			pathsCreator.RootDir)

		errorCollection.AddWrapperPtr(errWrap)
	}

	return errCount == errorCollection.Length()
}
