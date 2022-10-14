package fs

import "gitlab.com/evatix-go/errorwrapper"

func WriteStringToFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteStringToFile(
		isCreateParentDir,
		filePath,
		content)
}
