package pathchmod

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func changeOwnershipLinux(path, user, group string) *errorwrapper.Wrapper {
	chownUserGroupArg := user + constants.Colon + group

	return errcmd.New.ShellScript.ArgsErr(
		constants.ChmodCommand,
		chownUserGroupArg,
		path,
	)
}
