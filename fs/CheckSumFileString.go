package fs

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func CheckSumFileString(
	hashType hashas.Variant,
	location string,
) *errstr.Result {
	errBytes := CheckSumFileBytes(hashType, location)

	return &errstr.Result{
		Value:        errBytes.String(),
		ErrorWrapper: errBytes.ErrorWrapper,
	}
}
