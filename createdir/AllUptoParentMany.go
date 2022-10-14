package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AllUptoParentMany(mode os.FileMode, locations ...string) *errorwrapper.Wrapper {
	for _, path := range locations {
		if errWrap := fsinternal.CreateDirectoryAllUptoParent(path, mode); errWrap.HasError() {
			return errWrap
		}
	}

	return nil
}
