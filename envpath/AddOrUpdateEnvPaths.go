package envpath

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func AddOrUpdateEnvPaths(addOrUpdateEnvPaths ...string) *errorwrapper.Wrapper {
	if len(addOrUpdateEnvPaths) == 0 {
		return nil
	}

	return AddOrUpdateEnvPathsPtr(&addOrUpdateEnvPaths)
}
