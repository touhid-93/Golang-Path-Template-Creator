package envpath

import "gitlab.com/evatix-go/core/osconsts"

func GetEnvSeparator() string {
	if osconsts.IsWindows {
		return windowsEnvPathSplitter
	}

	return unixEnvPathSplitter
}
