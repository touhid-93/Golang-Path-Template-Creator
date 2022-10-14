package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func FilesPlusDirsByName(
	isRecursive bool,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     nil,
		ExcludingPaths:         nil,
		IsIncludeFilesOnly:     true,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      true,
		IsIncludeAll:           false,
		IsExcludeRoot:          true,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: false,
		IsNormalize:            false,
	}

	return instruction.StringsResults()
}
