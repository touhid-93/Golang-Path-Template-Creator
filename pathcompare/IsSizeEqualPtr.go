package pathcompare

import (
	"gitlab.com/evatix-go/core/corecmp"
)

func IsSizeEqualPtr(
	left, right *int64,
) bool {
	return corecmp.Integer64Ptr(left, right).IsEqual()
}
