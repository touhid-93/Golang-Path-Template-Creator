package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func Directories(
	rootPath string,
) *errstr.Results {
	return DirectoriesOptions(
		true,
		false,
		false,
		rootPath)
}
