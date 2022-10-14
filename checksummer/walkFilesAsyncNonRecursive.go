package checksummer

import (
	"os"
	"path/filepath"
)

func walkFilesAsyncNonRecursive(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		dirEntries, err := os.ReadDir(root)
		if err != nil {
			errc <- err

			return
		}

		// before exiting send nil error
		defer func() {
			errc <- nil
		}()

		for _, dirEntry := range dirEntries {
			// If the error is not nil or if the file is not regular, skip it
			if fileMode, err := dirEntry.Info(); err != nil || !fileMode.Mode().IsRegular() {
				continue
			}

			path := filepath.Join(root, dirEntry.Name())

			select {
			case paths <- path: // send path to the paths channel
			case <-done:
				return
			}
		}
	}()

	return paths, errc
}
