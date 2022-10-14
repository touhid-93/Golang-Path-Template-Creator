package normalize

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
)

func TrimPrefixUncPaths(
	isSkipOnUnix bool,
	locations ...string,
) []string {
	if isSkipOnUnix && osconsts.IsUnixGroup {
		return locations
	}

	length := len(locations)
	if length == 0 {
		return locations
	}

	wg := &sync.WaitGroup{}
	slice := stringslice.MakeLen(length)

	prefixReplacerFunc := func(index int, source string) {
		slice[index] = TrimPrefixUncPath(source)
	}

	wg.Add(length)
	for i, location := range locations {
		go prefixReplacerFunc(i, location)
	}

	return slice
}
