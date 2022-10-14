package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CreateDirectoryAll(
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	err := os.MkdirAll(location, mode)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			err,
			location)
}
