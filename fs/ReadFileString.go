package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func ReadFileString(filePath string) *errstr.Result {
	errBytes := ReadFile(filePath)

	return &errstr.Result{
		Value:        errBytes.String(),
		ErrorWrapper: errBytes.ErrorWrapper,
	}
}
