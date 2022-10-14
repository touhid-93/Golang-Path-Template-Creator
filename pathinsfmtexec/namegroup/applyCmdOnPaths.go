package namegroup

import "gitlab.com/evatix-go/errorwrapper"

func applyCmdOnPaths(
	cmdPrefix string,
	paths []string,
	isContinueOnError bool,
) *errorwrapper.Wrapper {
	if !isContinueOnError {
		return applyCmdOnPathsReturnOnErr(cmdPrefix, paths)
	}

	return applyCmdOnPathContinueOnError(cmdPrefix, paths)
}
