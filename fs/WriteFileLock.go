package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteFileLock(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteFile(
		isCreateParentDir,
		filePath,
		content)
}
