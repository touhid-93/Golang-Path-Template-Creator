package ispath

import (
	"os"
)

func Directory(path string) bool {
	currentFileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}

	return currentFileInfo.IsDir()
}
