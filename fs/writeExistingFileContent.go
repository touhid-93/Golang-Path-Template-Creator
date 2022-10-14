package fs

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func writeExistingFileContent(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	chmod, err := chmodhelper.GetExistingChmod(
		filePath)
	if err != nil {
		return errnew.Path.Error(
			errtype.ExistingChmodReadFailed,
			err,
			filePath)
	}

	writeErr := ioutil.WriteFile(
		filePath,
		content,
		chmod)

	if writeErr != nil {
		return errnew.Path.
			Messages(
				errtype.FileWrite,
				filePath,
				"fs.writeExistingFileContent",
				"Failed write file contents.",
				writeErr.Error())
	}

	return nil
}
