package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

func CreateDirectoryAllDefault(location string) *errorwrapper.Wrapper {
	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			os.MkdirAll(location, consts.DefaultDirectoryFileMode),
			location)
}
