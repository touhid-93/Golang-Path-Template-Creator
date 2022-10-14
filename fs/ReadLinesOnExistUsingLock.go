package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func ReadLinesOnExistUsingLock(filePath string) *errstr.Results {
	if IsPathExistsUsingLock(filePath) {
		return ReadFileLinesUsingLock(filePath)
	}

	return errstr.Empty.Results()
}
