package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func TempFileCreate(
	existingFileErrorWrapper *errorwrapper.Wrapper, // can be nil, current function return error
	relativeFilePath string,
	chmod os.FileMode,
) *OsFile {
	return DirFileCreate(
		existingFileErrorWrapper,
		os.TempDir(),
		relativeFilePath,
		chmod)
}
