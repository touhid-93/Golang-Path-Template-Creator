package fs

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/createdir"
)

func writeNewFileContentUsingFileMode(
	isCreateParentDir,
	isApplyChmodMust,
	isApplyChmodOnMismatchOnly bool,
	dirMode,
	fileMode os.FileMode,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	var createDirErr *errorwrapper.Wrapper
	if isCreateParentDir {
		createDirErr = createdir.AllUptoParent(
			filePath,
			dirMode)
	}

	if createDirErr.HasError() {
		return createDirErr
	}

	fileWriter := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               dirMode,
		ChmodFile:              fileMode,
		FilePath:               filePath,
		IsMustChmodApplyOnFile: isApplyChmodMust,
		IsApplyChmodOnMismatch: isApplyChmodOnMismatchOnly,
	}.InitializeDefault(isApplyChmodMust)

	writeErr := fileWriter.Write(content)

	if writeErr != nil {
		return errnew.Path.Error(
			errtype.FileWrite,
			writeErr,
			filePath)
	}

	return nil
}
