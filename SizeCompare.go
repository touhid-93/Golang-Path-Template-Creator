package pathhelper

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompareinternal"
)

func SizeCompare(
	left, right *int64,
) corecomparator.Compare {
	return pathcompareinternal.SizePtr(left, right)
}
