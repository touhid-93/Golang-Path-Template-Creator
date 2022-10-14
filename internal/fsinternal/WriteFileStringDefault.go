package fsinternal

import (
	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper"
)

// WriteFileStringDefault
//
//  Default chmod dir - 0755, file - 0644
func WriteFileStringDefault(
	filePath string,
	contentString string,
) *errorwrapper.Wrapper {
	return WriteFile(
		filemode.DirDefault,
		filemode.FileDefault,
		filePath,
		[]byte(contentString))
}
