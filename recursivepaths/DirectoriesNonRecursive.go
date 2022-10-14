package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func DirectoriesNonRecursive(
	rootPath string,
) *errstr.Results {
	return DirectoriesOptions(
		false,
		false,
		false,
		rootPath)
}
