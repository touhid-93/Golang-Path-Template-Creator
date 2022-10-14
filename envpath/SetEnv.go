package envpath

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SetEnv(variable, value string) *errorwrapper.Wrapper {
	err := os.Setenv(variable, value)

	if err != nil {
		return errnew.Messages.Many(
			errtype.EditFailed,
			"Failed add or update env variable.",
			"Variable, Value : ",
			variable,
			value)
	}

	return nil
}
