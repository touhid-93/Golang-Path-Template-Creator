package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func AllIf(
	isRecursive bool,
	rootPath string,
) *errstr.Results {
	if !isRecursive {
		return ReturnAsIs(
			false,
			false, rootPath)
	}

	return AllOptions(
		false,
		false,
		rootPath)
}
