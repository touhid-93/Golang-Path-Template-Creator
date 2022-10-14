package pathhelper

import (
	"sync"

	"gitlab.com/evatix-go/core"
)

func GetRelativePathsOfPtr(
	rootPath string,
	fullPaths *[]string,
) *[]string {
	if fullPaths == nil {
		return core.EmptyStringsPtr()
	}

	length := len(*fullPaths)
	list := make([]string, length)

	if length == 0 {
		return &list
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	processor := func(index int, fullPath string) {
		defer wg.Done()

		list[index] = GetRelativePath(fullPath, rootPath)
	}

	for i, fullPath := range *fullPaths {
		go processor(i, fullPath)
	}

	wg.Wait()

	return &list
}
