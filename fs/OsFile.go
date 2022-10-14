package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

// OsFile defer function must be called to close the file.
type OsFile struct {
	Location         string
	OsFile           *os.File
	ErrorWrapper     *errorwrapper.Wrapper
	DeferClosingFunc func() *errorwrapper.Wrapper // must be called upon checking IsDeferCloseRequired
	chmodWithError   *pathchmod.ChmodWithError
}

func (it *OsFile) ParentDir() string {
	return splitinternal.GetBaseDir(it.Location)
}

func (it *OsFile) BothExtension() (dotExt, ext string) {
	return splitinternal.GetBothExtension(it.Location)
}

func (it *OsFile) FileName() string {
	return splitinternal.GetName(it.Location)
}

func (it *OsFile) ChmodWithError() *pathchmod.ChmodWithError {
	if it.chmodWithError != nil {
		return it.chmodWithError
	}

	it.chmodWithError = pathchmod.ExistingChmodWithError(it.Location)

	return it.chmodWithError
}

func (it *OsFile) HasFile() bool {
	return it != nil && it.OsFile != nil
}

func (it *OsFile) IsSafeFile() bool {
	return it != nil && it.OsFile != nil && it.IsEmptyError()
}

func (it *OsFile) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *OsFile) IsEmptyFile() bool {
	return it == nil || it.OsFile == nil
}

func (it *OsFile) IsEmptyError() bool {
	return it != nil && it.ErrorWrapper.IsEmpty()
}

func (it *OsFile) IsDeferCloseRequired() bool {
	return it.HasFile() || it.DeferClosingFunc != nil
}

func (it *OsFile) AttachDeferCloseOnRequire() *errorwrapper.Wrapper {
	if it.IsDeferCloseRequired() {
		return it.DeferClosingFunc()
	}

	return nil
}

func (it *OsFile) ClearFileContents() *errorwrapper.Wrapper {
	writingError := it.OsFile.Truncate(0)

	if writingError != nil {
		return errnew.
			Path.
			Error(
				errtype.WriteFailed,
				writingError,
				it.Location)
	}

	return nil
}

func (it *OsFile) WriteString(
	writingString string,
) (hasWrittenSuccessfully bool, errWrap *errorwrapper.Wrapper) {
	writtenLen, writingError := it.OsFile.WriteString(writingString)

	if writingError != nil {
		return false,
			errnew.
				Path.
				Error(
					errtype.WriteFailed,
					writingError,
					it.Location)
	}

	return writtenLen > 0, nil
}

func (it *OsFile) WriteBytes(
	writingBytes []byte,
) (hasWrittenSuccessfully bool, errWrap *errorwrapper.Wrapper) {
	writtenLen, writingError := it.OsFile.Write(writingBytes)

	if writingError != nil {
		return false,
			errnew.
				Path.
				Error(
					errtype.WriteFailed,
					writingError,
					it.Location)
	}

	return writtenLen > 0, nil
}

func (it *OsFile) Close() *errorwrapper.Wrapper {
	return it.DeferClosingFunc()
}

func (it *OsFile) RwxWrapper() *pathchmod.RwxWrapperWithError {
	return pathchmod.ExistingRwxWrapperWithError(it.Location)
}

func (it *OsFile) FileInfo() (os.FileInfo, *errorwrapper.Wrapper) {
	fileInfo, err := it.OsFile.Stat()

	return fileInfo, errnew.
		Path.
		Error(
			errtype.PathStat,
			err,
			it.Location)
}

func (it *OsFile) ApplyChmod(mode os.FileMode) *errorwrapper.Wrapper {
	err := it.OsFile.Chmod(mode)

	if err != nil {
		return errnew.
			Path.
			Error(
				errtype.ChmodApplyFailed,
				err,
				it.Location)
	}

	return nil
}
