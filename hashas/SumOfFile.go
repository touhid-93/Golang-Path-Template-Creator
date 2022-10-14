package hashas

import (
	"io"
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SumOfFile(method Variant, filePath string) *errbyte.Results {
	if filePath == constants.EmptyString {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Messages.Many(
				errtype.EmptyString,
				"File name is empty"))
	}

	isExist, fileInfo := chmodhelper.IsPathExistsPlusFileInfo(filePath)

	if !isExist || fileInfo == nil || fileInfo.IsDir() {
		return errbyte.New.Results.ErrorWrapper(
			errnew.
				Path.
				Messages(
					errtype.InvalidPath,
					filePath,
					"File path either invalid or has permission issue or a folder for hash-checksum."))
	}

	hashWriter, errWrap := method.NewHash()

	if errWrap.HasError() {
		return errbyte.New.Results.ErrorWrapper(errWrap)
	}

	file, errOpen := os.Open(filePath)
	if errOpen != nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.
				Path.
				Error(
					errtype.FileRead,
					errOpen,
					"Error opening file : "+filePath,
				))
	}

	defer file.Close()

	_, errCopy := io.Copy(hashWriter, file)
	if errCopy != nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.
				Path.
				Error(
					errtype.Copy,
					errOpen,
					"Error copying to  file : "+filePath,
				))
	}

	hashedBytes := hashWriter.Sum(nil)

	return errbyte.New.Results.ValuesOnly(hashedBytes)
}
