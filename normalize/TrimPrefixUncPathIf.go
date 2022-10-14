package normalize

import "gitlab.com/evatix-go/core/osconsts"

func TrimPrefixUncPathIf(
	isSkipOnUnix bool,
	givenPath string,
) string {
	if isSkipOnUnix && osconsts.IsUnixGroup {
		return givenPath
	}

	return TrimPrefixes(
		givenPath,
		uncPrefixes...)
}
