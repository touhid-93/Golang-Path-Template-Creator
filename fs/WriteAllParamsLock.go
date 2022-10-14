package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteAllParamsLock(
	isCreateParentDir,
	isApplyChmodMust,
	isApplyChmodOnMismatchOnly,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	parentDirChmod os.FileMode,
	fileChmod os.FileMode,
	filePath string,
	contents []byte,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteAllParams(
		isCreateParentDir,
		isApplyChmodMust,
		isApplyChmodOnMismatchOnly,
		isSkipOnNilObject,
		isKeepExistingFileModeOnExist,
		parentDirChmod,
		fileChmod,
		filePath,
		contents)
}
