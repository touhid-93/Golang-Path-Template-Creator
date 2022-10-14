package envpath

import "gitlab.com/evatix-go/core/coredata/corestr"

func hashsetEnvPathToSingleString(hashset *corestr.Hashset) string {
	compiledJoinedPath := hashset.
		Collection().
		SortedAsc().
		Join(GetEnvSeparator())

	return compiledJoinedPath
}
