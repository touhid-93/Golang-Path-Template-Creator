package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func All(location string, mode os.FileMode) *errorwrapper.Wrapper {
	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			os.MkdirAll(location, mode),
			location)
}
