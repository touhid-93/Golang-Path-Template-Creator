package checksummer

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"

	"gitlab.com/evatix-go/pathhelper/hashas"
)

func Md5(filePath string) *errbyte.Results {
	return FileRaw(filePath, hashas.Md5)
}
