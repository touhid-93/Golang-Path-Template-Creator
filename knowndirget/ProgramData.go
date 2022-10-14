package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

func ProgramData() string {
	if !osconsts.IsWindows {
		return ""
	}

	return knowndir.ProgramData.CombineWith(WindowsRoot())
}
