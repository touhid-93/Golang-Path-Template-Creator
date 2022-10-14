package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdefer"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// GetOsFile defer function must be called to close the file.
func GetOsFile(
	existingErrorWrapper *errorwrapper.Wrapper, // can be nil
	osFlag int,
	fileMode os.FileMode,
	filePath string,
) *OsFile {
	osFile, fileOpenErr := os.OpenFile(
		filePath,
		osFlag,
		fileMode)

	fileOpeningError := errnew.
		Path.
		Error(
			errtype.FileRead,
			fileOpenErr,
			filePath)

	deferClosingFunc := func() *errorwrapper.Wrapper {
		return errdefer.CloseFile(
			filePath,
			existingErrorWrapper,
			osFile)
	}

	return &OsFile{
		Location:         filePath,
		OsFile:           osFile,
		ErrorWrapper:     fileOpeningError,
		DeferClosingFunc: deferClosingFunc,
	}
}
