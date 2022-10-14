package hexchecksum

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func OfFilesRequest(request *FilesRequest) *FilesResult {
	request.SortFileNamesIf(request.IsSortFileNames)

	hexOfListing := OfFilesListIf(
		request.IsGenerateFileListChecksum,
		request.Method,
		request.Files...)

	filesCount := len(request.Files)

	isQuickExit := !request.IsGenerateContentsChecksum &&
		filesCount == 0

	hasIssuesAndExit := !isQuickExit &&
		hexOfListing.HasError() &&
		request.IsExitOnError()

	if isQuickExit || hasIssuesAndExit {
		return &FilesResult{
			HexFilesListChecksum:     hexOfListing.Value,
			HexFilesContentsChecksum: constants.EmptyString,
			Method:                   request.Method,
			ErrorWrapper:             hexOfListing.ErrorWrapper,
			FilesCount:               filesCount,
		}
	}

	hexContentsChecksum := OfFilesContentsAsync(
		request.IsSortFilesChecksum,
		false,
		request.Method,
		request.Files...)

	mergedErr := errnew.Merge.UsingStackSkip(
		codestack.Skip1,
		hexOfListing.ErrorWrapper,
		hexContentsChecksum.ErrorWrapper)

	return &FilesResult{
		HexFilesListChecksum:     hexOfListing.Value,
		HexFilesContentsChecksum: hexContentsChecksum.Value,
		Method:                   request.Method,
		ErrorWrapper:             mergedErr,
		FilesCount:               filesCount,
	}
}
