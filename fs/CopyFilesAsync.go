package fs

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CopyFilesAsync(
	isMove bool,
	sourceToDestination map[string]string,
) *errorwrapper.Wrapper {
	if len(sourceToDestination) == 0 {
		return nil
	}

	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}

	copierOrMoveFunc := CopyFile
	if isMove {
		copierOrMoveFunc = MoveFile
	}

	var sliceErr []string
	copierOrMoveExecFunc := func(source, destination string) bool {
		defer wg.Done()

		copyErr := copierOrMoveFunc(
			source,
			destination)

		if copyErr.IsSuccess() {
			return true
		}

		// failed
		locker.Lock()
		sliceErr = append(
			sliceErr,
			copyErr.String())
		locker.Unlock()

		return false
	}

	wg.Add(len(sourceToDestination))
	for sourcePath, destinationPath := range sourceToDestination {
		go copierOrMoveExecFunc(sourcePath, destinationPath)
	}

	wg.Wait()

	err := errcore.SliceToError(sliceErr)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.Copy,
			err,
			constants.EmptyString)
}
