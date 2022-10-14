package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompareinternal"
)

func FileInfoCompare(
	left, right os.FileInfo,
) corecomparator.Compare {
	return pathcompareinternal.FileInfoLastModified(left, right)
}
