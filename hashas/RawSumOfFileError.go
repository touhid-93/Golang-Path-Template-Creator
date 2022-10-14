package hashas

import (
	"io"
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func RawSumOfFileError(
	isVerifyPathExistence bool,
	method Variant,
	filePath string,
) ([]byte, error) {
	if isVerifyPathExistence {
		if filePath == constants.EmptyString {
			return nil, errtype.EmptyFilePath.ErrorNoRefs("")
		}

		isExist, fileInfo := chmodhelper.IsPathExistsPlusFileInfo(filePath)

		if !isExist || fileInfo == nil || fileInfo.IsDir() {
			return nil, errtype.InvalidPath.ReferencesCsvError(
				filePathInvalidOrPermissionIssueMessage,
				filePath)
		}
	}

	hashWriter, err := method.NewHashError()

	if err != nil {
		return nil, err
	}

	file, errOpen := os.Open(filePath)
	if errOpen != nil {
		return nil, errtype.FileRead.ReferencesCsvError(
			fileOrPathReadIssue+errOpen.Error(),
			filePath)
	}

	defer file.Close()

	_, errCopy := io.Copy(hashWriter, file)
	if errCopy != nil {
		return nil, errtype.Copy.ReferencesCsvError(
			"io.Copy issue."+errCopy.Error(),
			filePath)
	}

	hashedBytes := hashWriter.Sum(nil)

	return hashedBytes, nil
}
