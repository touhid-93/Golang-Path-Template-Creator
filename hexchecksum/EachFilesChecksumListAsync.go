package hexchecksum

import (
	"sync"

	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

// EachFilesChecksumListAsync
//
// Returns each files checksum as slice of errstr.Results
//
// each index represents file Index => checksum index same.
// It continues on error
func EachFilesChecksumListAsync(
	hashMethod hashas.Variant,
	fullFilePaths ...string,
) *errstr.Results {
	length := len(fullFilePaths)

	if length == 0 {
		return errstr.Empty.Results()
	}

	if length <= consts.NonAsyncSafeRange {
		return EachFilesChecksumList(
			hashMethod,
			fullFilePaths...)
	}

	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}
	var sliceErr []string
	checkSumSlice := make(
		[]string,
		length)

	hexChecksum := func(index int, source string) bool {
		defer wg.Done()
		hexFileChecksumResult := hashMethod.
			HexSumOfFile(source)

		if hexFileChecksumResult.IsSuccess() {
			checkSumSlice[index] = hexFileChecksumResult.Value

			return true
		}

		// failed
		locker.Lock()
		sliceErr = append(
			sliceErr,
			hexFileChecksumResult.
				ErrorWrapper.
				String())
		locker.Unlock()

		return false
	}

	wg.Add(length)

	for i, filePath := range fullFilePaths {
		go hexChecksum(i, filePath)
	}

	wg.Wait()

	err := errcore.SliceToError(sliceErr)

	if err == nil {
		return errstr.New.Results.Strings(checkSumSlice)
	}

	errWrap := errnew.Type.Error(
		errtype.CheckSumCorrupted,
		err)

	// has error
	return errstr.New.Results.Create(errWrap, checkSumSlice)
}
