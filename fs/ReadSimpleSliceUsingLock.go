package fs

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func ReadSimpleSliceUsingLock(
	filePath string,
) (*corestr.SimpleSlice, *errorwrapper.Wrapper) {
	results := ReadFileLinesUsingLock(filePath)

	if results.HasIssuesOrEmpty() {
		return corestr.Empty.SimpleSlice(), results.ErrorWrapper
	}

	return &corestr.SimpleSlice{Items: results.Values}, nil
}
