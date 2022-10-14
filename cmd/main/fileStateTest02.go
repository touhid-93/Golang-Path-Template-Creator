package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/filestate"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

func fileStateTest02() {
	files := recursivepaths.FilesOptionsExcept(
		false,
		false,
		[]string{".git"},
		nil,
		pathsconst.RootDir)

	info, errWrap := filestate.NewInfoCollectionUsingFilePathsAsync(
		hashas.DefaultFastHashMethod,
		true,
		files.Values...)

	errWrap.HandleError()

	fmt.Println(info.String())
}
