package fileinfogetter

import (
	"os"
	"sync"
)

// Get For each path converted to file info wrapper and finally returns as an array.
func Get(fullPaths *[]string) *[]os.FileInfo {
	if fullPaths == nil {
		return &[]os.FileInfo{}
	}

	length := len(*fullPaths)
	list := make([]os.FileInfo, length)

	if length == 0 {
		return &list
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	processor := func(index int, fullPath string) {
		defer wg.Done()
		currentFileInfo, err := os.Stat(fullPath)

		if err == nil && currentFileInfo != nil {
			list[index] = currentFileInfo
		} else {
			list[index] = nil
		}
	}

	for i, fullPath := range *fullPaths {
		go processor(i, fullPath)
	}

	wg.Wait()

	return &list
}
