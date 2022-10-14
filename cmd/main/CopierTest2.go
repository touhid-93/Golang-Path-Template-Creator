package main

import (
	"fmt"
	"io/ioutil"

	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func CopierTest2() {
	tmpDir, _ := ioutil.TempDir("", "ttt")

	fmt.Println(pathsconst.RootDir)
	errWrap := copyrecursive.Do(
		false,
		pathsconst.RootDir,
		tmpDir)

	errWrap.HandleError()
}
