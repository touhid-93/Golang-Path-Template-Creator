package pathhelper

import (
	"time"

	"gitlab.com/evatix-go/core/corecmp"
	"gitlab.com/evatix-go/core/corecomparator"
)

func LastModifiedDateCompare(
	left, right *time.Time,
) corecomparator.Compare {
	return corecmp.TimePtr(left, right)
}
