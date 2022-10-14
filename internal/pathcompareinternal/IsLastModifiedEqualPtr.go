package pathcompareinternal

import (
	"time"

	"gitlab.com/evatix-go/core/corecmp"
)

func IsLastModifiedEqualPtr(
	left, right *time.Time,
) bool {
	return corecmp.TimePtr(left, right).IsEqual()
}
