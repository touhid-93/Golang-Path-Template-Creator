package fsinternal

import (
	"io/ioutil"
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteFile(
	dirMode, fileMode os.FileMode,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return NullContentErrorWrap(
			filePath)
	}

	if IsPathExists(filePath) {
		chmod, err := chmodhelper.GetExistingChmod(filePath)

		if err != nil {
			return errnew.Path.ErrorUsingStackSkip(
				codestack.Skip1,
				errtype.ExistingChmodReadFailed,
				err,
				filePath)
		}

		writeErr := ioutil.WriteFile(
			filePath,
			content,
			chmod)

		if writeErr != nil {
			return errnew.Path.ErrorUsingStackSkip(
				codestack.Skip1,
				errtype.FileWrite,
				writeErr,
				filePath)
		}

		return nil
	}

	dirCreateErr := CreateDirectoryAllUptoParent(
		filePath, dirMode)

	if dirCreateErr.HasError() {
		return dirCreateErr
	}

	writeErr := ioutil.WriteFile(
		filePath,
		content,
		fileMode)

	if writeErr != nil {
		return errnew.Path.ErrorUsingStackSkip(
			codestack.Skip1,
			errtype.FileWrite,
			writeErr,
			filePath)
	}

	return nil
}
