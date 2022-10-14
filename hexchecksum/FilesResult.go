package hexchecksum

import (
	"strconv"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

type FilesResult struct {
	HexFilesListChecksum     string
	HexFilesContentsChecksum string
	FilesCount               int
	Method                   hashas.Variant
	ErrorWrapper             *errorwrapper.Wrapper
}

func (it *FilesResult) HexChecksumOfResult() *errstr.Result {
	if it == nil {
		return nil
	}

	var slice [5]string
	slice[0] = it.HexFilesListChecksum
	slice[1] = it.HexFilesContentsChecksum
	slice[2] = it.Method.Name()
	slice[3] = it.ErrorWrapper.String()
	slice[4] = strconv.Itoa(it.FilesCount)

	return it.Method.HexSumOfAny(slice)
}

func (it *FilesResult) IsEmpty() bool {
	return it.HasNoChecksum()
}

func (it *FilesResult) HasNoChecksum() bool {
	return it == nil ||
		it.HexFilesListChecksum == "" &&
			it.HexFilesContentsChecksum == ""
}

func (it *FilesResult) HasAnyChecksum() bool {
	return it != nil &&
		it.HexFilesListChecksum != "" ||
		it.HexFilesContentsChecksum != ""
}

func (it *FilesResult) HasFilesListChecksum() bool {
	return it != nil && it.HexFilesListChecksum != ""
}

func (it *FilesResult) Json() corejson.Result {
	return corejson.New(it)
}

func (it *FilesResult) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *FilesResult) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func (it *FilesResult) HasContentsChecksum() bool {
	return it != nil && it.HexFilesContentsChecksum != ""
}

func (it *FilesResult) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *FilesResult) IsSuccess() bool {
	return it != nil && it.ErrorWrapper.IsSuccess()
}

func (it *FilesResult) IsFailed() bool {
	return it != nil && it.ErrorWrapper.IsFailed()
}

func (it *FilesResult) IsEqual(another *FilesResult) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.FilesCount != another.FilesCount {
		return false
	}

	if it.HexFilesListChecksum != another.HexFilesListChecksum {
		return false
	}

	if it.HexFilesContentsChecksum != another.HexFilesContentsChecksum {
		return false
	}

	if it.Method != another.Method {
		return false
	}

	return it.ErrorWrapper.IsEquals(another.ErrorWrapper)
}
