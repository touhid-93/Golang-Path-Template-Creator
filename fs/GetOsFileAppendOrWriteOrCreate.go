package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

// GetOsFileAppendOrWriteOrCreate defer function must be called to close the file.
func GetOsFileAppendOrWriteOrCreate(
	existingErrorWrapper *errorwrapper.Wrapper, // can be nil
	filePath string,
	fileMode os.FileMode,
) *OsFile {
	return GetOsFile(
		existingErrorWrapper,
		FlagWriteOrCreateOrAppend,
		fileMode,
		filePath,
	)
}
