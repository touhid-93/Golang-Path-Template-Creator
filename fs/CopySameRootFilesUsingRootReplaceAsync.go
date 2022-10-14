package fs

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

// CopySameRootFilesUsingRootReplaceAsync
//
// Replaces all sourcePaths using sourceRoot and
// usage destinationPath to replace the root
// and copies file on the destination path
//
// Restrictions:
//  - Here, all the sourcePaths needs to be in the same dir.
func CopySameRootFilesUsingRootReplaceAsync(
	isContinueOnError,
	isClearDestination,
	isMove bool,
	sourceRoot string,
	sourcePaths []string,
	destinationRootPath string,
) *errorwrapper.Wrapper {
	if len(sourcePaths) == 0 {
		return nil
	}

	deleteErr := deletepaths.RecursiveOnExistIf(
		isClearDestination,
		destinationRootPath)

	if deleteErr.HasError() {
		return deleteErr
	}

	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}
	isFailed := false
	var sliceErr []string

	copierOrMoveFunc := CopyFile
	if isMove {
		copierOrMoveFunc = MoveFile
	}

	copyOrMoveExecFunc := func(source string) bool {
		defer wg.Done()

		destination := pathjoin.TrimReplacePrefixRoot(
			true,
			false,
			source,
			sourceRoot,
			destinationRootPath)

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
		isFailed = true
		locker.Unlock()

		return false
	}

	if isContinueOnError {
		wg.Add(len(sourcePaths))
		for _, sourceFullPath := range sourcePaths {
			go copyOrMoveExecFunc(sourceFullPath)
		}
	} else {
		// no continue
		for _, sourceFullPath := range sourcePaths {
			wg.Add(constants.One)
			go copyOrMoveExecFunc(sourceFullPath)

			if isFailed {
				break
			}
		}
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
