package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func SimpleFilterFilesPlusDirsByName(
	isRecursive bool,
	excludingRootName []string,
	excludingPaths []string,
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     excludingRootName,
		ExcludingPaths:         excludingPaths,
		IsIncludeFilesOnly:     true,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      true,
		IsIncludeAll:           false,
		IsExcludeRoot:          true,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: false,
		IsNormalize:            false,
	}

	return instruction.Result().SimpleFilterFullPathsAsync(filter)
}
