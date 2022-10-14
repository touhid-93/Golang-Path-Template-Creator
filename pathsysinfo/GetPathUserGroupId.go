package pathsysinfo

import "gitlab.com/evatix-go/pathhelper/fileinfopath"

func GetPathUserGroupId(filePath string) *PathUserGroupId {
	pathWithInfo := fileinfopath.New(filePath)

	return GetPathUserGroupIdUsing(pathWithInfo)
}
