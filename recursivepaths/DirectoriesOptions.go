package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func DirectoriesOptions(
	isRecursive,
	isNormalize,
	isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     nil,
		ExcludingPaths:         nil,
		IsIncludeFilesOnly:     false,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      true,
		IsIncludeAll:           false,
		IsExcludeRoot:          false,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	return instruction.StringsResults()
}
