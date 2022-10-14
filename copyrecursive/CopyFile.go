package copyrecursive

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string, fileMode os.FileMode) error {
	sourceFileStat, errStat := os.Stat(src)
	if errStat != nil {
		return errStat
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, errOpen := os.Open(src)
	if errOpen != nil {
		return errOpen
	}

	defer source.Close()

	// Create all the parent folder if needed
	if err := os.MkdirAll(filepath.Dir(dst), fileMode); err != nil {
		return err
	}

	destination, errCreate := os.Create(dst)
	if errCreate != nil {
		return errCreate
	}

	defer destination.Close()
	_, errOpen = io.Copy(destination, source)

	return errOpen
}
