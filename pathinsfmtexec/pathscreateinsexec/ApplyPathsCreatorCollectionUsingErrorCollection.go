package pathscreateinsexec

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyPathsCreatorCollectionUsingErrorCollection(
	isLock,
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	errorCollection *errwrappers.Collection,
	pathsCreatorCollection *pathinsfmt.PathsCreatorCollection,
) (
	isSuccess bool,
) {
	if pathsCreatorCollection == nil {
		return true
	}

	errCount := errorCollection.Length()

	if isDeleteAllBeforeCreate {
		for _, instruction := range pathsCreatorCollection.PathsCreatorItems {
			errorCollection.AddAnyFunctions(instruction.DeleteAllPaths)
		}
	}

	workingPaths := pathsCreatorCollection.
		LazyFlatPathsIf(isLazyPaths)

	// paths create
	_, filesCreateErr := createpath.CreateMany(
		isLock,
		isIgnoreOnExist,
		workingPaths)

	errorCollection.AddWrapperPtr(filesCreateErr)

	if filesCreateErr.HasError() {
		return false
	}

	// apply chmod
	for _, pathsCreator := range pathsCreatorCollection.PathsCreatorItems {
		applyRwxOnPathCreators(&pathsCreator, errorCollection)
	}

	// apply groups
	if osconsts.IsUnixGroup {
		for _, pathsCreator := range pathsCreatorCollection.PathsCreatorItems {
			applyUserNameGroupNameOnPathCreators(&pathsCreator, errorCollection)
		}
	}

	return errCount == errorCollection.Length()
}
