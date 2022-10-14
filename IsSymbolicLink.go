package pathhelper

import (
	"log"
	"os"
)

func IsSymbolicLink(path string) bool {
	file, err := os.Lstat(path)

	if err != nil {
		log.Fatal(err)
	}

	mode := file.Mode()

	return mode&os.ModeSymlink != 0
}
