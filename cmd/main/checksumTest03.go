package main

import (
	"fmt"

	"gitlab.com/evatix-go/asynchelper/syncparallel"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
)

func checksumTest03() {
	files := pkgRootFiles()
	var detailedResult *hexchecksum.DetailedResult
	var filesResult *hexchecksum.FilesResult
	requestSample := hexchecksum.FilesRequest{
		Method:                     hashas.DefaultFastHashMethod,
		IsSortFileNames:            false,
		IsSortFilesChecksum:        true,
		IsGenerateContentsChecksum: true,
		IsGenerateFileListChecksum: true,
		Files:                      files,
	}

	// requestSample.SortFileNames()
	// fmt.Println(strings.Join(requestSample.Files, constants.NewLineUnix))

	syncparallel.Tasks(
		func() {
			detailedResult = hexchecksum.DetailedResultOfRequestAsync(&requestSample)

			detailedResult.Hashmap = nil
			detailedResult.FilesResult.ErrorWrapper = nil
		},
		func() {
			filesResult = hexchecksum.OfFilesRequest(&requestSample)
			// filesResult.ErrorWrapper = nil
		},
	)

	fmt.Println("detailedResult")
	fmt.Println(detailedResult.JsonString())

	fmt.Println("filesResult")
	fmt.Println(filesResult.JsonString())
}
