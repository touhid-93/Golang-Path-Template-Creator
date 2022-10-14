package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/core/coreindexes"
)

// Get the final path exePath + ....
func GetExecutableCombinePath(paths ...string) string {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	exeBase, _ := Split(exe)

	if paths == nil {
		return exeBase
	}

	length := len(paths)

	if length == 0 {
		return exeBase
	}

	list := make([]string, length+1)
	list[coreindexes.First] = exeBase

	for i, cPath := range paths {
		list[i+1] = cPath
	}

	return GetCombinePathsWithPtr(&list)
}
