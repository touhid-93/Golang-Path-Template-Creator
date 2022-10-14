package pathcompare

import (
	"gitlab.com/evatix-go/core/corecmp"
	"gitlab.com/evatix-go/core/corecomparator"
)

func SizePtr(
	left, right *int64,
) corecomparator.Compare {
	return corecmp.Integer64Ptr(left, right)
}
