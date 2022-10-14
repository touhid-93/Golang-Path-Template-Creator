package normalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
)

func SimpleJoinPaths(locations ...string) string {
	if len(locations) == 0 {
		return constants.EmptyString
	}

	return stringslice.NonEmptyJoin(
		locations,
		osconsts.PathSeparator)
}
