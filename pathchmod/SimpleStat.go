package pathchmod

import (
	"os"
	"time"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

type SimpleStat struct {
	Location        string
	FileInfo        os.FileInfo
	Name            string // name with extension
	HasFileInfo     bool
	InvalidFileInfo bool
	IsNotExist      bool
	IsExist         bool
	IsDir           bool
	IsFile          bool
	ErrorWrapper    *errorwrapper.Wrapper
}

func (it *SimpleStat) LastModifiedAt() *time.Time {
	if it == nil || it.InvalidFileInfo {
		return nil
	}

	mod := it.FileInfo.ModTime()

	return &mod
}

func (it *SimpleStat) Size() *int64 {
	if it == nil || it.InvalidFileInfo {
		return nil
	}

	size := it.FileInfo.Size()

	return &size
}

func (it *SimpleStat) IsEmptyError() bool {
	return it == nil || it.ErrorWrapper.IsEmptyError()
}

func (it *SimpleStat) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *SimpleStat) ReadString() *errstr.Result {
	errWrap := it.notFileError()
	if errWrap.HasError() {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: errWrap,
		}
	}

	errBytes := fsinternal.ReadFile(it.Location)

	return &errstr.Result{
		Value:        errBytes.String(),
		ErrorWrapper: errBytes.ErrorWrapper,
	}
}

func (it *SimpleStat) ReadStringMust() string {
	rs := it.ReadString()
	rs.ErrorWrapper.HandleError()

	return rs.Value
}

func (it *SimpleStat) notFileError() *errorwrapper.Wrapper {
	if !it.IsExist || it.IsDir {
		return errnew.
			Path.
			Messages(
				errtype.File,
				it.Location,
				"Cannot read invalid path or a directory.")
	}

	return nil
}

func (it *SimpleStat) ReadBytes() *errbyte.Results {
	errWrap := it.notFileError()
	if errWrap.HasError() {
		return errbyte.New.Results.ErrorWrapper(
			errWrap)
	}

	return fsinternal.ReadFile(it.Location)
}

func (it *SimpleStat) ReadBytesMust() []byte {
	rs := it.ReadBytes()
	rs.ErrorWrapper.HandleError()

	return rs.SafeValues()
}

// FileName
// Returns file name with extension
func (it *SimpleStat) FileName() string {
	if it.HasFileInfo {
		return it.Name
	}

	return constants.EmptyString
}

func (it *SimpleStat) FileNameWithoutExt() string {
	if it.HasFileInfo {
		return splitinternal.GetFileNameWithoutExt(it.Name)
	}

	return constants.EmptyString
}

func (it *SimpleStat) BothExtension() (dotExt, ext string) {
	if it.HasFileInfo {
		return splitinternal.GetBothExtension(it.FileInfo.Name())
	}

	return constants.EmptyString, constants.EmptyString
}

func (it *SimpleStat) DotExtension() (fileName, dotExt string) {
	if it.HasFileInfo {
		return splitinternal.GetFileNameDotExt(it.FileInfo.Name())
	}

	return constants.EmptyString, constants.EmptyString
}

func (it *SimpleStat) FileNameExt() string {
	return it.FileName()
}

func (it *SimpleStat) CheckSum(hashType hashas.Variant) *errbyte.Results {
	errWrap := it.notFileError()
	if errWrap.HasError() {
		return errbyte.New.Results.ErrorWrapper(
			errWrap)
	}

	allBytes := it.ReadBytes()

	if allBytes.HasError() {
		return errbyte.New.Results.ErrorWrapper(
			allBytes.ErrorWrapper)
	}

	return hashType.SumOf(allBytes.Values)
}

func (it *SimpleStat) HexCheckSumString(hashType hashas.Variant) *errstr.Result {
	errWrap := it.notFileError()
	if errWrap.HasError() {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: errWrap,
		}
	}

	allBytes := it.ReadBytes()

	if allBytes.HasError() {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: allBytes.ErrorWrapper,
		}
	}

	return hashType.HexSumOf(allBytes.Values)
}

func (it *SimpleStat) GetChmodWithError() *ChmodWithError {
	return ExistingChmodWithError(it.Location)
}

func (it *SimpleStat) GetRwxWithError() *RwxWrapperWithError {
	return ExistingRwxWrapperWithError(it.Location)
}
