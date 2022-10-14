package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/pathhelper/copyrecursive"
)

func CopierTest() {
	tmpDir, _ := ioutil.TempDir("", "ttt")
	_, b, _, _ := runtime.Caller(0)
	srcDir := filepath.Join(filepath.Dir(b), "..", "..")
	fmt.Println("to", tmpDir)
	fmt.Println("from", srcDir)
	errWrap := copyrecursive.NewCopier(
		srcDir, tmpDir, copyrecursive.Options{
			IsSkipOnExist:      false,
			IsRecursive:        false,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        true,
		},
	).Copy()

	errWrap.HandleError()
}
