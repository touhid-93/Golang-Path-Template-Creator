package checksummer

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"

	"gitlab.com/evatix-go/pathhelper/hashas"
)

func Sha256(filePath string) *errbyte.Results {
	return FileRaw(filePath, hashas.Sha256)
}
