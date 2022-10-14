package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
)

// ChmodLinuxCmdApply
//
// Format : chmod -R 777 /var/lib/powerdns
func ChmodLinuxCmdApply(
	isRecursive bool,
	isSkipOnError bool,
	rwxWrapper *chmodhelper.RwxWrapper,
	paths ...string,
) *errwrappers.Collection {
	length := len(paths)
	errCollection := errwrappers.Empty()

	if length == 0 {
		return errCollection
	}

	chmodPrefix := cmdprefix.Chmod(
		isRecursive,
		rwxWrapper)

	if isSkipOnError {
		for _, currentPath := range paths {
			command := errcmd.ArgsJoin(
				chmodPrefix,
				currentPath)

			errWrap := errcmd.
				New.
				BashScript.
				ArgsErr(command)

			if errWrap.HasError() {
				return errCollection.AddWrapperPtr(errWrap)
			}
		}

		return errCollection
	}

	for _, currentPath := range paths {
		command := errcmd.ArgsJoin(chmodPrefix, currentPath)
		errorWrapper := errcmd.
			New.
			BashScript.
			LinesResult(command).
			ErrorWrapper()

		errCollection.AddWrapperPtr(errorWrapper)
	}

	return errCollection
}
