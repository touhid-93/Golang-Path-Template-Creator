package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func FilesOptions(
	isRecursive,
	isNormalize,
	isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     nil,
		ExcludingPaths:         nil,
		IsIncludeFilesOnly:     true,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      false,
		IsIncludeAll:           false,
		IsExcludeRoot:          true,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	return instruction.StringsResults()
}
