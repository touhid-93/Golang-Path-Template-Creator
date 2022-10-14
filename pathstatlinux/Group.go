package pathstatlinux

import (
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

type Group struct {
	IntIdNameValidation
	systemGroup *pathsysinfo.GroupInfo
}

func (g *Group) SystemGroup() *pathsysinfo.GroupInfo {
	if g.systemGroup != nil {
		return g.systemGroup
	}

	if !g.HasValidId {
		g.systemGroup = invalidGroupInfo

		return g.systemGroup
	}

	g.systemGroup = pathsysinfo.GetGroupInfo(g.Name)

	return g.systemGroup
}
