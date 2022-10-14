package deletepaths

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SingleIf(
	isRemove bool,
	location string,
) *errorwrapper.Wrapper {
	if !isRemove {
		return nil
	}

	err := os.Remove(location)
	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(errtype.DeletePathFailed, err, location)
}
