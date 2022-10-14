package fs

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func Read(location string, filename string) *errbyte.Results {
	compiledFilePath := pathjoin.JoinNormalized(location, filename)

	return ReadFile(compiledFilePath)
}
