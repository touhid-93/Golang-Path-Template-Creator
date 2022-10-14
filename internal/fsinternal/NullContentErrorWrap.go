package fsinternal

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func NullContentErrorWrap(filePath string) *errorwrapper.Wrapper {
	return errnew.Ref.Message(
		errtype.NullOrEmpty,
		"file-path",
		filePath,
		"cannot write or append empty or bytes into file.",
	)
}
