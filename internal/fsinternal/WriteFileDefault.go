package fsinternal

import (
	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper"
)

// WriteFileDefault
//
//  Default chmod dir - 0755, file - 0644
func WriteFileDefault(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	return WriteFile(
		filemode.DirDefault,
		filemode.FileDefault,
		filePath,
		content)
}
