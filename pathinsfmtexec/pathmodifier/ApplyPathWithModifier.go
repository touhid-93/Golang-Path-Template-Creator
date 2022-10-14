package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

func ApplyPathWithModifier(
	errCollection *errwrappers.Collection,
	pathWithModifier *pathinsfmt.PathWithModifier,
) (isSuccess bool) {
	if pathWithModifier.IsModifierUndefined() {
		return true
	}

	if pathWithModifier.IsSkipInvalid &&
		pathWithModifier.IsInvalidPath() {
		return true
	}

	recursiveFiles := recursivepaths.AllIf(
		pathWithModifier.IsRecursive,
		pathWithModifier.CompiledPath())

	errCollection.AddWrapperPtr(
		recursiveFiles.ErrorWrapper)

	if recursiveFiles.HasError() {
		return false
	}

	return ApplySimple(
		false,
		errCollection,
		pathWithModifier.Modifier,
		recursiveFiles.SafeValues())
}
