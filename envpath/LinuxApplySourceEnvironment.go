package envpath

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func LinuxApplySourceEnvironment() *errorwrapper.Wrapper {
	return errcmd.New.BashScript.ArgsDefault("source", etcEnvPath).
		CompiledResult().
		ErrorWrapper()
}
