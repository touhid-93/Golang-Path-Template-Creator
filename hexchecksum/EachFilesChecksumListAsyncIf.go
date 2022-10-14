package hexchecksum

import (
	"sort"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

// EachFilesChecksumListAsyncIf
//
// Returns each files checksum as slice of errstr.Results
//
// each index represents file Index => checksum index same.
//
// Generates only if condition meets
func EachFilesChecksumListAsyncIf(
	isGenerate bool,
	isSortFileName bool,
	hashMethod hashas.Variant,
	fullFilePaths ...string,
) *errstr.Results {
	if isGenerate && isSortFileName && len(fullFilePaths) > 0 {
		sort.Strings(fullFilePaths)
	}

	if isGenerate {
		return EachFilesChecksumListAsync(
			hashMethod,
			fullFilePaths...)
	}

	return errstr.Empty.Results()
}
