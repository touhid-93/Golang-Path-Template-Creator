package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AllUptoParent(location string, mode os.FileMode) *errorwrapper.Wrapper {
	return fsinternal.CreateDirectoryAllUptoParent(location, mode)
}
