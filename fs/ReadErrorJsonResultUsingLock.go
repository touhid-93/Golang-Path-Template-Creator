package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errjson"

func ReadErrorJsonResultUsingLock(filePath string) *errjson.Result {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return ReadErrorJsonResult(filePath)
}
