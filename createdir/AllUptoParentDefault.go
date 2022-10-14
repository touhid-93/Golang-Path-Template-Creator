package createdir

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AllUptoParentDefault(location string) *errorwrapper.Wrapper {
	return fsinternal.CreateDirectoryAllUptoParent(location, DefaultDirectoryFileMode)
}
