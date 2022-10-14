package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func JsonWriteMarshalUsingLock(
	isCreateParentDir,
	isApplyChmodOnMismatchOnly bool,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	marshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return JsonWriteMarshal(
		isCreateParentDir,
		isApplyChmodOnMismatchOnly,
		isSkipOnNilObject,
		isKeepExistingFileModeOnExist,
		dirMode,
		fileMode,
		filePath,
		marshallingObjectRef)
}
