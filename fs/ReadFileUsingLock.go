package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errbyte"

func ReadFileUsingLock(filePath string) *errbyte.Results {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return ReadFile(filePath)
}
