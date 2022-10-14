package symlink

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func Create(path, linkName string) *errorwrapper.Wrapper {
	err := os.Symlink(path, linkName)

	if err != nil {
		return errnew.Ref.TwoWithError(
			errtype.SymbolicLink,
			err,
			"Source Symbolic Link",
			path,
			"Symbolic Link Place",
			linkName)
	}

	return nil
}
