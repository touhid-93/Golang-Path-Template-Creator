package pathfixer

import (
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

func FixMany(isExpand bool, paths ...string) []string {
	if len(paths) == 0 {
		return []string{}
	}

	slice := stringslice.MakeLen(len(paths))

	for i, singlePath := range paths {
		slice[i] = Fix(isExpand, singlePath)
	}

	return slice
}
