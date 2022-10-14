package hexchecksum

import (
	"sort"

	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

type FilesRequest struct {
	Method            hashas.Variant
	IsContinueOnError bool
	IsSortFileNames   bool
	// if true, then sorts checksums
	// before generate final single
	// checksum ( not needed )
	IsSortFilesChecksum        bool
	IsGenerateContentsChecksum bool
	IsGenerateFileListChecksum bool
	Files                      []string
	isSorted                   issetter.Value
}

func (it *FilesRequest) SortFileNames() {
	if it == nil || it.isSorted.IsInitBoolean() || len(it.Files) == 0 {
		return
	}

	sort.Strings(it.Files)
	it.isSorted = issetter.True
}

func (it *FilesRequest) IsExitOnError() bool {
	return it == nil || !it.IsContinueOnError
}

func (it *FilesRequest) SortFileNamesIf(isSort bool) {
	if it == nil || !isSort {
		return
	}

	it.SortFileNames()
}

func (it *FilesRequest) SortFileNamesBasedOnCondition() {
	if it == nil || !it.IsSortFileNames {
		return
	}

	it.SortFileNames()
}
