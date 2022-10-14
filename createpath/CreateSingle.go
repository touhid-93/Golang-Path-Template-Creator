package createpath

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateSingle(
	isLock bool,
	filePath string,
) (
	*os.File,
	*errorwrapper.Wrapper,
) {
	if isLock {
		lockerMutex.Lock()
		defer lockerMutex.Unlock()
	}

	dirCreateErr := fsinternal.CreateDirectoryAllUptoParentDefault(
		filePath)

	if dirCreateErr.HasError() {
		return nil, dirCreateErr
	}

	file, err := os.Create(filePath)

	if err != nil {
		return file, errnew.
			Path.
			Error(
				errtype.CreatePathFailed,
				err,
				filePath)
	}

	return file, nil
}
