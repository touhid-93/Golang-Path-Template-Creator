package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

// GetOsFileWriteOrCreate defer function must be called to close the file.
func GetOsFileWriteOrCreate(
	existingErrorWrapper *errorwrapper.Wrapper, // can be nil
	filePath string,
	fileMode os.FileMode,
) *OsFile {
	return GetOsFile(
		existingErrorWrapper,
		FlagWriteOrCreate,
		fileMode,
		filePath,
	)
}
