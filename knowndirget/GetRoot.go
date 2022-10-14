package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"
)

func GetRoot() string {
	if osconsts.IsWindows {
		return WindowsRoot()
	}

	return UnixRoot()
}
