package envpath

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/messages"
)

// linuxCrudEnvPath
//
// linuxEnvPathCrudFunc could be -
//  - linuxEnvRemoveAction or
//  - linuxEnvAddOrUpdateAction
func linuxCrudEnvPath(
	performingFunc linuxEnvPathCrudFunc,
	envPaths []string,
	isApplyEnvironmentSource bool,
) *errorwrapper.Wrapper {
	if len(envPaths) == 0 {
		return nil
	}

	if !osconsts.IsUnixGroup || !osconsts.IsLinux {
		return errnew.Messages.Many(
			errtype.NotSupportOperatingSystem,
			"linuxCrudEnvPath",
			messages.CannotAddUpdateRemoveEnvPath)
	}

	contentsBytes := fsinternal.ReadFile(etcEnvPath)

	if contentsBytes.HasError() {
		return contentsBytes.ErrorWrapper
	}

	envPathString := contentsBytes.String()

	newEnvPathCompiled := performingFunc(envPaths, envPathString)
	prependPathEqual := compileEnvPathToLinuxRawEnvPathFormat(
		newEnvPathCompiled)

	writeErrorWrapper := fsinternal.WriteStringToFile(
		etcEnvPath,
		prependPathEqual)

	if writeErrorWrapper.HasError() {
		return writeErrorWrapper
	}

	if setEnvErr := SetEnvPath(newEnvPathCompiled); setEnvErr.HasError() {
		return setEnvErr
	}

	if isApplyEnvironmentSource {
		return LinuxApplySourceEnvironment()
	}

	return nil
}
