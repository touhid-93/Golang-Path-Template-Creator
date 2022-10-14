package fs

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func CheckSumFileStringUsingLock(
	hashType hashas.Variant,
	location string,
) *errstr.Result {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return CheckSumFileString(hashType, location)
}
