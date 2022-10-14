package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Program Files directory on windows with OS architecture x32
//
// In 32 bit machine program files is just program file, in 64 bit program files is x86
//
// Reference : https://bit.ly/3tX95N6
func GetProgramFiles32() string {
	if !osconsts.IsWindows {
		return ""
	}

	if osconsts.IsX32Architecture {
		return knowndir.ProgramFiles32Or64.CombineWith(WindowsRoot())
	}

	// x64
	return knowndir.ProgramFilesX32In64.CombineWith(WindowsRoot())
}
