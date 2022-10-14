package testwrappers

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

var (
	TempDir          = os.TempDir()
	TempPkgDir       = TempDir + "/" + PkgName
	RootPath1        = TempPkgDir + "/path-root1"
	RootPath2        = TempPkgDir + "/path-root2"
	RootPath3        = TempPkgDir + "/path-root3"
	RootPath4        = TempPkgDir + "/path-root4"
	RootPath5        = TempPkgDir + "/path-root5"
	FilesCollection1 = []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
		"file4.txt",
		"file5.txt",
	}
	Locations = []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
		"file4.txt",
		"file5.txt",
	}

	DefaultRwxOwnerGroupOther = &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}

	DefaultBaseUserNameGroupName = &pathinsfmt.BaseUserNameGroupName{
		UserGroupName: pathinsfmt.UserGroupName{
			BaseGroupName: pathinsfmt.BaseGroupName{GroupName: "root"},
			UserName:      "a",
		},
	}
	DefaultUserNameGroupName = &DefaultBaseUserNameGroupName.UserGroupName
	DefaultWorkingPaths      = GetSetupPaths()
	AllPathsString           = strings.Join(DefaultWorkingPaths, constants.CommaSpace)
)
