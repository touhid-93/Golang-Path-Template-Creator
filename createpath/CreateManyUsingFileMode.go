package createpath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateManyUsingFileMode(
	isLock,
	isIgnoreOnExist bool,
	mode os.FileMode,
	files []string,
) ([]*os.File, *errorwrapper.Wrapper) {
	if len(files) == 0 {
		return []*os.File{}, nil
	}

	if isLock {
		lockerMutex.Lock()
		defer lockerMutex.Unlock()
	}

	slice := make(
		[]*os.File,
		constants.Zero,
		len(files))

	if isIgnoreOnExist {
		for _, filePath := range files {
			if fsinternal.IsPathExists(filePath) {
				continue
			}

			file, errWrap := CreateSingleUsingFileMode(
				false,
				mode,
				filePath,
			)

			if errWrap.HasError() {
				return slice, errWrap
			}

			slice = append(slice, file)
		}

		return slice, nil
	}

	// no checking create
	for _, filePath := range files {
		file, errWrap := CreateSingleUsingFileMode(
			false,
			mode,
			filePath,
		)

		if errWrap.HasError() {
			return slice, errWrap
		}

		slice = append(slice, file)
	}

	return slice, nil
}
