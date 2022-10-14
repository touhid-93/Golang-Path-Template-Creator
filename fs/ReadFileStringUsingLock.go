package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func ReadFileStringUsingLock(filePath string) *errstr.Result {
	errBytes := ReadFileUsingLock(filePath)

	return errBytes.ErrStr()
}
