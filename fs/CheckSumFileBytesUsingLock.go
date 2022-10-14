package fs

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func CheckSumFileBytesUsingLock(
	hashType hashas.Variant,
	location string,
) *errbyte.Results {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return CheckSumFileBytes(hashType, location)
}
