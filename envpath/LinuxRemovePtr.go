package envpath

import "gitlab.com/evatix-go/errorwrapper"

func LinuxRemovePtr(envPaths []string, isApplyEnvironmentSource bool) *errorwrapper.Wrapper {
	return linuxCrudEnvPath(
		linuxEnvRemoveAction,
		envPaths,
		isApplyEnvironmentSource)
}
