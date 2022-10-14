package fsinternal

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func CreateDirectoryAllUptoParentDefault(location string) *errorwrapper.Wrapper {
	parentDir := ParentDir(location)

	if IsDirectory(parentDir) {
		return nil
	}

	return CreateDirectoryAllDefault(parentDir)
}
