package fsinternal

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func Write(
	location string,
	filename string,
	content []byte,
) *errorwrapper.Wrapper {
	compileFilePath := pathjoin.JoinNormalized(
		location,
		filename) // todo check how to remove this reference

	return WriteFileDefault(
		compileFilePath, content)
}
