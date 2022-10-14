package envpath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SetEnvPath(compiledPath string) *errorwrapper.Wrapper {
	err := os.Setenv(constants.Path, compiledPath)

	if err != nil {
		return errnew.Messages.Many(
			errtype.EditFailed,
			"Failed to Add or Update environment paths.",
			compiledPath)
	}

	return nil
}
