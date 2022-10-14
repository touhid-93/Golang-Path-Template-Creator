package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func checkSumCheck() {
	result := hexchecksum.OfFilesContentsAsync(
		false,
		false,
		hashas.Sha256,
		"cmd/main/main.go")

	// result.ErrorWrapper.HandleError()
	result.ErrorWrapper.Log()
	fmt.Println(result.String())

	rs := pathsysinfo.GetPathUserGroupId("cmd/main")

	fmt.Println(converters.Any.ToFullNameValueString(rs))
	fmt.Println(rs.Error)
}
