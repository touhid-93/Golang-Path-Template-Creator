package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"
)

// Returns Program Files directory on windows with OS architecture x32
//
// In 32 bit machine program files is just program file, in 64 bit program files is x86
//
// Reference : https://bit.ly/3tX95N
func ProgramFiles() string {
	if osconsts.IsX32Architecture {
		return GetProgramFiles32()
	}

	return ProgramFiles64()
}
