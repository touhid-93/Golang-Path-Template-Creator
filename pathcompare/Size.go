package pathcompare

import (
	"gitlab.com/evatix-go/core/corecmp"
	"gitlab.com/evatix-go/core/corecomparator"
)

func Size(
	left, right int64,
) corecomparator.Compare {
	return corecmp.Integer64(left, right)
}
