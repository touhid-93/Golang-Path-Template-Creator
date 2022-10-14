package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
)

func WriteErrorJsonResultUsingLock(
	isCreateParentDir,
	isSkipErrorOnNilOrEmpty bool,
	errJsonResult *errjson.Result,
	location string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteErrorJsonResult(
		isCreateParentDir,
		isSkipErrorOnNilOrEmpty,
		errJsonResult,
		location,
	)
}
