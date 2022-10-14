package envvars

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyEnvVar(
	environmentVariable *pathinsfmt.EnvironmentVariable,
) *errorwrapper.Wrapper {
	if environmentVariable == nil {
		return nil
	}

	variableKeyValueAttach :=
		environmentVariable.Name +
			constants.EqualSymbol +
			environmentVariable.Value

	cmdResult := errcmd.New.BashScript.ArgsDefault(
		consts.Export,
		variableKeyValueAttach,
	)

	return cmdResult.CompiledErrorWrapper()
}
