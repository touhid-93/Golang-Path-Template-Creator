package fstestwrapper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

var (
	SetupFiles = pathinsfmt.PathsCreator{
		BasePathsCreator: pathinsfmt.BasePathsCreator{
			RootDir: pathsconst.DefaultTempTestDir,
			Files: []string{
				setupFilePath,
			},
			IsNormalize: true,
		},
		ApplyRwx: &chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "rwx",
			Other: "rwx",
		},
		ApplyUserGroup: nil,
	}
)
