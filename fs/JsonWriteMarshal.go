package fs

import (
	"os"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func JsonWriteMarshal(
	isCreateParentDir,
	isApplyChmodOnMismatchOnly bool,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	marshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	if marshallingObjectRef == nil && isSkipOnNilObject {
		return nil
	}

	allBytes, err := corejson.Serialize.Raw(marshallingObjectRef)

	if err != nil {
		return errnew.Path.Error(
			errtype.Marshalling,
			err,
			filePath,
		)
	}

	// new content, apply chmod any way as new
	return WriteFileUsingFileMode(
		isCreateParentDir,
		isApplyChmodOnMismatchOnly,
		isKeepExistingFileModeOnExist,
		dirMode,
		fileMode,
		filePath,
		allBytes,
	)
}
