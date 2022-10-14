package namegroup

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func applyCmdOnPathsReturnOnErr(
	cmdPrefix string,
	paths []string,
) *errorwrapper.Wrapper {
	for _, currentPath := range paths {
		if currentPath == "" {
			return errorwrapper.NewPtr(errtype.EmptyFilePath)
		}

		// chgrp groupName path or chown -R $user:$group /dir
		pathCmd := cmdPrefix +
			constants.Space +
			currentPath

		errWrapper := errcmd.
			New.BashScript.ArgsErr(pathCmd)

		if errWrapper.HasError() {
			return errWrapper
		}
	}

	return nil
}
