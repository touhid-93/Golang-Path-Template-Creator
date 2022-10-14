package namegroup

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

func ApplySimpleSingle(
	isRecursive bool,
	userName string,
	groupName string,
	location string,
) *errorwrapper.Wrapper {
	return pathchmod.ChangeOwnershipOptions(
		isRecursive,
		location,
		userName,
		groupName)
}
