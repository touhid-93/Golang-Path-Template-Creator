package pathhelper

import (
	"sync"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

// Each path can give number wrappers from GetFileInfoWrappersFrom
// This method call GetFileInfoWrappersFrom for each path given and returns as an array of it.
func GetFileInfoWrappersArrayOf(
	separator string,
	fullPaths *[]string,
	isNormalize bool,
) *[]*fileinfo.Wrappers {
	if fullPaths == nil {
		return &[]*fileinfo.Wrappers{}
	}

	length := len(*fullPaths)
	list := make([]*fileinfo.Wrappers, length)

	if length == 0 {
		return &list
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	processor := func(index int, fullPath string) {
		defer wg.Done()

		list[index] = fileinfo.NewWrappersPtr(
			fullPath,
			separator,
			isNormalize)
	}

	for i, fullPath := range *fullPaths {
		go processor(i, fullPath)
	}

	wg.Wait()

	return &list
}
