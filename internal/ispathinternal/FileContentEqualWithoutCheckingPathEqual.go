package ispathinternal

import (
	"bytes"
	"io"
	"os"
)

// FileContentEqualWithoutCheckingPathEqual
//
// FileContentEqual
//
// Returns false if contents are not equal at any point
//
// Reference : https://stackoverflow.com/a/30038571
func FileContentEqualWithoutCheckingPathEqual(
	leftFullPath string,
	rightFullPath string,
) bool {
	f1, f1Err := os.Open(leftFullPath)
	if f1 == nil || f1Err != nil {
		return false
	}

	defer f1.Close()

	f2, f2Err := os.Open(rightFullPath)
	if f2 == nil || f2Err != nil {
		return false
	}

	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, readErr1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, readErr2 := f2.Read(b2)

		if readErr1 != nil || readErr2 != nil {
			if readErr1 == io.EOF && readErr2 == io.EOF {
				return true
			} else if readErr1 == io.EOF || readErr2 == io.EOF {
				return false
			} else {
				return false
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}
