package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func AllOptions(
	isNormalize, isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     nil,
		ExcludingPaths:         nil,
		IsIncludeFilesOnly:     false,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      false,
		IsIncludeAll:           true,
		IsExcludeRoot:          false,
		IsRecursive:            true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	return instruction.StringsResults()
}
