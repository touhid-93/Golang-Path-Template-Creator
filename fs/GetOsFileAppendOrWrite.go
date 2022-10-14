package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

// GetOsFileAppendOrWrite defer function must be called to close the file.
func GetOsFileAppendOrWrite(
	existingErrorWrapper *errorwrapper.Wrapper, // can be nil
	filePath string,
	fileMode os.FileMode,
) *OsFile {
	return GetOsFile(
		existingErrorWrapper,
		FlagAppendOrWrite,
		fileMode,
		filePath,
	)
}
