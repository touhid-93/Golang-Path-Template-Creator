package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// ProgramFiles64
//
// Returns Program Files directory on windows with OS architecture x64
//
// Program files for x64 bit is always `Program Files`,
// one cannot install x64 bit application on a x32 bit machine.
//
// Reference : https://bit.ly/3tX95N6
func ProgramFiles64() string {
	if !osconsts.IsWindows {
		return ""
	}

	return knowndir.
		ProgramFiles32Or64.
		CombineWith(WindowsRoot())
}
