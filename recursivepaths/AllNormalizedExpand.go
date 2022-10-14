package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func AllNormalizedExpand(rootPath string) *errstr.Results {
	return AllOptions(
		true,
		true, rootPath)
}
