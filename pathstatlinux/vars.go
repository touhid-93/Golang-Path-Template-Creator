package pathstatlinux

import (
	"gitlab.com/evatix-go/pathhelper/internal/deferrwrappers"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

var (
	invalidGroupInfo = pathsysinfo.InvalidGroupInfo(deferrwrappers.InvalidSystemGroup)
	invalidUserInfo  = pathsysinfo.InvalidUserInfo(deferrwrappers.InvalidSystemUser)
)
