package ispath

import (
	"os"
)

func File(path string) bool {
	currentFileInfo, err := os.Stat(path)

	if err != nil || currentFileInfo == nil {
		return false
	}

	return !currentFileInfo.IsDir()
}
