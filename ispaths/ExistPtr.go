package ispaths

import (
	"os"
	"sync"
)

func Exist(paths ...string) []bool {
	if paths == nil {
		return []bool{}
	}

	length := len(paths)
	list := make([]bool, length)

	if length == 0 {
		return list
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	inPlaceProcessor := func(index int, fullPath string) {
		defer wg.Done()
		currentFileInfo, err := os.Stat(fullPath)

		if err == nil && currentFileInfo != nil {
			list[index] = true
		} else {
			list[index] = false
		}
	}

	for i, fullPath := range paths {
		go inPlaceProcessor(i, fullPath)
	}

	wg.Wait()

	return list
}
