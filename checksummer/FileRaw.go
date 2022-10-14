package checksummer

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"

	"gitlab.com/evatix-go/pathhelper/hashas"
)

func FileRaw(filePath string, v hashas.Variant) *errbyte.Results {
	return v.SumOfFile(filePath)
}
