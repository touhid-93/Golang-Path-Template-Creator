package envpath

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func RemoveEnvPathsPtr(removeEnvPaths *[]string) *errorwrapper.Wrapper {
	if removeEnvPaths == nil || len(*removeEnvPaths) == 0 {
		return nil
	}

	envPaths := ReadEnvPathsPtr()
	hashset := corestr.New.Hashset.StringsPtr(
		envPaths)

	for _, removeEnvPath := range *removeEnvPaths {
		hashset.Remove(removeEnvPath)
	}

	compiledPath := hashsetEnvPathToSingleString(hashset)

	return SetEnvPath(compiledPath)
}
