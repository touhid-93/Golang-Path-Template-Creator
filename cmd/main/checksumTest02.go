package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
)

func checksumTest02() {
	rs := hexchecksum.OfFilesRequest(&hexchecksum.FilesRequest{
		Method:                     hashas.DefaultFastHashMethod,
		IsGenerateContentsChecksum: true,
		IsGenerateFileListChecksum: true,
		Files:                      pkgRootFiles(),
	})

	fmt.Println(rs.JsonString())
}
