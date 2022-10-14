package createpath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/chmodinternal"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateManySameDirWithFileMode(
	isLock,
	isIgnoreOnExist bool,
	mode os.FileMode,
	rootDir string,
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

			file, errWrap := CreateSingle(
				false,
				filePath,
			)

			if errWrap.HasError() {
				return slice, errWrap
			}

			slice = append(slice, file)
		}

		return slice, chmodinternal.Apply(
			false,
			mode,
			rootDir,
			files)
	}

	// no checking create
	for _, filePath := range files {
		file, errWrap := CreateSingle(
			false,
			filePath,
		)

		if errWrap.HasError() {
			return slice, errWrap
		}

		slice = append(slice, file)
	}

	return slice, chmodinternal.Apply(
		false,
		mode,
		rootDir,
		files)
}
