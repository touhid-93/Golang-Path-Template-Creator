package hexchecksum

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func OfChecksums(
	isGenerate bool,
	isSortChecksum bool,
	hashMethod hashas.Variant,
	hexChecksums ...string,
) *errstr.Result {
	if len(hexChecksums) == 0 || !isGenerate {
		return errstr.Empty.Result()
	}

	sortIf(isSortChecksum, hexChecksums)

	return hashMethod.HexSumOfAny(
		hexChecksums)
}
