package pathinsfmt

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
)

type RemoveInstruction struct {
	BaseIsRecursive
	GenericPathsCollection
	IsRemoveOnlyOnExist bool `json:"IsRemoveOnlyOnExist,omitempty"`
	IsContinueOnError   bool `json:"IsContinueOnError,omitempty"`
}

func (it *RemoveInstruction) Apply(isLazyPaths bool) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	it.ApplyUsingErrorCollection(
		isLazyPaths,
		errCollection)

	return errCollection
}

func (it *RemoveInstruction) ApplyUsingErrorCollection(
	isLazyPaths bool,
	errorCollection *errwrappers.Collection,
) (
	isSuccess bool,
) {
	if it.IsRecursive {
		return it.applyUsingRecursive(
			isLazyPaths,
			errorCollection)
	}

	return it.applyUsingNonRecursive(
		isLazyPaths,
		errorCollection)
}

func (it *RemoveInstruction) applyUsingRecursive(
	isLazyPaths bool,
	errorCollection *errwrappers.Collection,
) (
	isSuccess bool,
) {
	if it.IsContinueOnError {
		return deletepaths.AllRecursiveContinueOnErrorIfExist(
			it.IsRemoveOnlyOnExist,
			errorCollection,
			it.LazyPathsIf(isLazyPaths))
	}

	errWrap := deletepaths.AllRecursiveOnExistIf(
		it.IsRemoveOnlyOnExist,
		it.LazyPathsIf(isLazyPaths))

	errorCollection.AddWrapperPtr(errWrap)

	return errWrap.IsEmpty()
}

func (it *RemoveInstruction) applyUsingNonRecursive(
	isLazyPaths bool,
	errorCollection *errwrappers.Collection,
) (
	isSuccess bool,
) {
	if it.IsContinueOnError {
		return deletepaths.AllContinueOnErrorIfExists(
			it.IsRemoveOnlyOnExist,
			errorCollection,
			it.LazyPathsIf(isLazyPaths))
	}

	errWrap := deletepaths.AllOnExistIf(
		it.IsRemoveOnlyOnExist,
		it.LazyPathsIf(isLazyPaths)...)

	errorCollection.AddWrapperPtr(errWrap)

	return errWrap.IsEmpty()
}
