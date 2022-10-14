package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/filestate"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func fileStateTest01() {
	info, errWrap := filestate.NewInfoDefault(pathsconst.RootDir)

	errWrap.HandleError()

	fmt.Println(info.String())
}
