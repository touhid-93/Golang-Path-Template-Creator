package pathcompareinternal

import (
	"os"

	"gitlab.com/evatix-go/core/corecomparator"
)

func FileInfoLastModified(
	left, right os.FileInfo,
) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	leftMod := left.ModTime()
	rightMod := right.ModTime()

	return LastModifiedPtr(&leftMod, &rightMod)
}
