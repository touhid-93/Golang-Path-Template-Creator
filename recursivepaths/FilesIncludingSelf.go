package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

// FilesIncludingSelf includes given path
func FilesIncludingSelf(
	isNormalize,
	isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		IsIncludeFilesOnly:     true,
		IsExcludeRoot:          false,
		IsRecursive:            true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	return instruction.StringsResults()
}
