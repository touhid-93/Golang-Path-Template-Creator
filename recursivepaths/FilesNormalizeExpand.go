package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func FilesNormalizeExpand(rootPath string) *errstr.Results {
	return FilesOptions(
		true,
		true,
		true,
		rootPath)
}
