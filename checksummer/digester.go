package checksummer

import (
	"hash"
	"io"
	"os"
)

// digester reads path names from paths and sends digests of the corresponding
// files on c until either paths or done is closed.
func digester(
	h hash.Hash,
	done <-chan struct{},
	paths <-chan string,
	c chan<- walkResult,
) {
	for path := range paths {
		file, err := os.Open(path)
		if err != nil {
			select {
			case c <- walkResult{path: path, err: err}:
			case <-done:
				return
			}

			continue
		}

		h.Reset()
		_, err = io.Copy(h, file)
		file.Close()

		select {
		case c <- walkResult{path, h.Sum(nil), err}:
		case <-done:
			return
		}
	}
}
