package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/pathsconst"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

func filterTest02() {
	files := recursivepaths.FilesExtensionFilter(
		true,
		pathsconst.RootDir+" x",
		".md", ".txt")

	files.ErrorWrapper.HandleError()

	fmt.Println(files.String())
}
