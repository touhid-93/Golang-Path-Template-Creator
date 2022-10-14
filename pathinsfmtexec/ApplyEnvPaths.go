package pathinsfmtexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/envpath"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyEnvPaths(baseEnvPaths *pathinsfmt.BaseEnvPaths) *errorwrapper.Wrapper {
	if baseEnvPaths == nil || len(baseEnvPaths.EnvPaths) == 0 {
		return nil
	}

	return envpath.AddOrUpdateEnvPaths(baseEnvPaths.EnvPaths...)
}
