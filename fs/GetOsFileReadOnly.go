package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

// GetOsFileReadOnly defer function must be called to close the file.
func GetOsFileReadOnly(
	existingErrorWrapper *errorwrapper.Wrapper, // can be nil
	filePath string,
	fileMode os.FileMode,
) *OsFile {
	return GetOsFile(
		existingErrorWrapper,
		FlagReadOnly,
		fileMode,
		filePath,
	)
}
