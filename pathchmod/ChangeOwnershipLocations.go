package pathchmod

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

// ChangeOwnershipLocations not implemented
//
// TODO :
// https://gitlab.com/evatix-go/pathhelper/-/issues/56
func ChangeOwnershipLocations(
	isContinueOnError,
	isRecursive bool,
	user, group string,
	errorCollection *errwrappers.Collection,
	locations []string,
) *errorwrapper.Wrapper {
	panic(errnew.NotImpl("https://gitlab.com/evatix-go/pathhelper/-/issues/56"))
}
