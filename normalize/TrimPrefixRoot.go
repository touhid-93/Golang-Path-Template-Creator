package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/osconsts"
)

func TrimPrefixRoot(
	source,
	rootTrimmingPrefix string,
) string {
	if source == "" || rootTrimmingPrefix == "" {
		return source
	}

	if osconsts.IsWindows {
		source = TrimPrefixUncPath(source)
		rootTrimmingPrefix = TrimPrefixUncPath(rootTrimmingPrefix)
	}

	return strings.TrimPrefix(source, rootTrimmingPrefix)
}
