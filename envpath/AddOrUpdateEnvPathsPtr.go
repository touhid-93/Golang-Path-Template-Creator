package envpath

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func AddOrUpdateEnvPathsPtr(addOrUpdateEnvPaths *[]string) *errorwrapper.Wrapper {
	if addOrUpdateEnvPaths == nil || len(*addOrUpdateEnvPaths) == 0 {
		return nil
	}

	envPaths := ReadEnvPathsPtr()
	hashset := corestr.New.Hashset.StringsPtr(
		envPaths)

	hashset.AddStringsPtr(addOrUpdateEnvPaths)
	compiledPath := hashsetEnvPathToSingleString(hashset)

	return SetEnvPath(compiledPath)
}
