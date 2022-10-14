package fileinfogetter

import (
	"os"
	"sync"

	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

// For each path converted to file info wrapper and finally returns as an array.
func GetWithErrors(
	fullPaths *[]string,
) (
	*[]os.FileInfo,
	*errwrappers.Collection,
) {
	errsCollection := errwrappers.Empty()
	errMutex := sync.Mutex{}
	if fullPaths == nil {
		return &[]os.FileInfo{}, errsCollection
	}

	length := len(*fullPaths)
	list := make([]os.FileInfo, length)

	if length == 0 {
		return &list, errsCollection
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
			errMutex.Lock()
			errsCollection.
				AddTypeError(
					errtype.FileInfo,
					err)
			errMutex.Unlock()
		}
	}

	for i, fullPath := range *fullPaths {
		go processor(i, fullPath)
	}

	wg.Wait()

	return &list, errsCollection
}
