package symlink

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/messages"
)

func RemoveSymlink(path string) *errorwrapper.Wrapper {
	result := IsSymLink(path)

	if result.HasError() {
		return result.ErrorWrapper
	}

	if !result.Value {
		invalidSymLink := fmt.Sprintf(messages.InvalidSymlinkMessageFormat, path)

		return errnew.Messages.Many(errtype.SymbolicLink, invalidSymLink)
	}

	err := os.Remove(path)

	if err != nil {
		invalidSymLink := fmt.Sprintf(messages.CannotRemoveSymLink, path)

		return errnew.Messages.Many(errtype.SymbolicLink, invalidSymLink)
	}

	return nil
}
