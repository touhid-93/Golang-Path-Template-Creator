package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplySimpleSingle(
	modifier *pathinsfmt.PathModifier,
	location string,
) *errorwrapper.Wrapper {
	errCollection := errwrappers.Empty()

	locations := []string{
		location,
	}

	isSuccess := ApplySimple(
		false,
		errCollection,
		modifier,
		locations)

	if !isSuccess {
		return errCollection.GetAsErrorWrapperPtr()
	}

	return nil
}
