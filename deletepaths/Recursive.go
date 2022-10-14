package deletepaths

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func Recursive(location string) *errorwrapper.Wrapper {
	err := os.RemoveAll(location)
	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.DeletePathFailed,
			err,
			location+"->recursive remove failed.")
}
