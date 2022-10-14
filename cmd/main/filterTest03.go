package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/pathsconst"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

func filterTest03() {
	files := recursivepaths.FilterDirsByName(
		true,
		pathsconst.RootDir,
		"cmd", "main", "dirinfo", "internal")

	files.ErrorWrapper.HandleError()

	fmt.Println(files.String())
}
