package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func SimpleFilterFilesOptionsAsync(
	isRecursive,
	isNormalize, isExpandEnv bool,
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		IsRelativePath:         false,
		IsIncludeFilesOnly:     true,
		IsIncludeDirsOnly:      false,
		IsIncludeAll:           false,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	result := instruction.Result()

	return result.SimpleFilterFullPathsAsync(filter)
}
