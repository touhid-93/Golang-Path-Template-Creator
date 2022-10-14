package pathcompare

import (
	"time"

	"gitlab.com/evatix-go/core/corecmp"
)

func IsNotEqualLastModifiedPtr(
	left, right *time.Time,
) bool {
	return !corecmp.TimePtr(left, right).IsEqual()
}
