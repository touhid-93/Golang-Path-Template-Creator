package fs

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
)

// WriteAnyLock should be used with caution,
// it should only be used when one is trying
// to convert struct to string then Save it to file.
// it is not the right way to the file back.
func WriteAnyLock(
	isCreateParentDir bool,
	filePath string,
	content interface{},
) *errorwrapper.Wrapper {
	anyToString := stringutil.AnyToString(
		content)

	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteFile(
		isCreateParentDir,
		filePath,
		[]byte(anyToString))
}
