package hexchecksum

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

type DetailedResult struct {
	FilesResult
	Hashmap *corestr.Hashmap // FilePath -> Hex checksum
}

func (it *DetailedResult) IsValidResult() bool {
	return it != nil &&
		it.FilesResult.IsSuccess() &&
		it.FilesCount == it.Hashmap.Length()
}

func (it *DetailedResult) IsInvalidResult() bool {
	return it == nil ||
		it.FilesResult.IsFailed() ||
		it.FilesCount != it.Hashmap.Length()
}

func (it DetailedResult) Json() corejson.Result {
	return corejson.New(it)
}

func (it DetailedResult) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it DetailedResult) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func EmptyDetailedResult() *DetailedResult {
	return &DetailedResult{
		FilesResult: FilesResult{
			HexFilesListChecksum:     "",
			HexFilesContentsChecksum: "",
			FilesCount:               0,
			Method:                   0,
			ErrorWrapper:             nil,
		},
		Hashmap: nil,
	}
}

func EmptyDetailedResultWithErr(
	hashMethod hashas.Variant,
	errWrap *errorwrapper.Wrapper,
) *DetailedResult {
	return &DetailedResult{
		FilesResult: FilesResult{
			HexFilesListChecksum:     "",
			HexFilesContentsChecksum: "",
			FilesCount:               0,
			Method:                   hashMethod,
			ErrorWrapper:             errWrap,
		},
		Hashmap: corestr.Empty.Hashmap(),
	}
}
