package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func All(rootPath string) *errstr.Results {
	return AllOptions(
		false,
		false,
		rootPath)
}
