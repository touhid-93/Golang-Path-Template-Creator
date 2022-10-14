package nginxlinuxpath

import "gitlab.com/evatix-go/core/coredata/corestr"

var (
	DefaultDirStructure = GetFullDirStructure(
		true,
		DefaultDirChmod,
		DefaultRoot)
	defaultMimeTypesPath = corestr.SimpleStringOnce{}
)
