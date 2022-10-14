package checksummer

import (
	"errors"
	"os"
	"path/filepath"
)

// walkFilesAsync starts a goroutine to walk the directory tree at root and send the
// path of each regular file on the string channel.  It sends the walkResult of the
// walk on the error channel.  If done is closed, walkFilesAsync abandons its work.
func walkFilesAsync(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		// Close the paths channel after Walk returns.
		defer close(paths)
		// No select needed for this send, since errc is buffered.
		errc <- filepath.Walk(
			root,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.Mode().IsRegular() {
					return nil
				}

				select {
				case paths <- path: // send path to the paths channel
				case <-done: // cancel the walk, note: receiving from closed channel is instantaneous
					// caller decided to cancel the traversal
					return errors.New("walk canceled")
				}

				return nil
			})
	}()

	return paths, errc
}
