package pathhelper

import (
	"sync"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

// For each path converted to file info wrapper and finally returns as an array.
func GetEachPathConvertedToFileInfoWrapper(separator string, fullPaths *[]string) *[]*fileinfo.Wrapper {
	if fullPaths == nil {
		return &[]*fileinfo.Wrapper{}
	}

	length := len(*fullPaths)
	list := make([]*fileinfo.Wrapper, length)

	if length == 0 {
		return &list
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	processor := func(index int, fullPath string) {
		defer wg.Done()

		list[index] = fileinfo.New(
			fullPath,
			separator)
	}

	for i, fullPath := range *fullPaths {
		go processor(i, fullPath)
	}

	wg.Wait()

	return &list
}
