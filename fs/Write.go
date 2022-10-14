package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func Write(
	isCreateParentDir bool,
	location string,
	filename string,
	content []byte,
) *errorwrapper.Wrapper {
	compileFilePath := pathjoin.JoinNormalized(location, filename)

	return WriteFile(
		isCreateParentDir,
		compileFilePath,
		content)
}
