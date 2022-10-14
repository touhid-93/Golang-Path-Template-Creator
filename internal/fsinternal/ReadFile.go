package fsinternal

import (
	"io/ioutil"

	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ReadFile(filePath string) *errbyte.Results {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.
				Path.
				Error(
					errtype.ReadRequestFailed,
					err,
					filePath,
				))
	}

	if fileBytes == nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.
				Path.
				Messages(
					errtype.EmptyContent,
					filePath,
					"fsinternal.ReadFile",
					"Location doesn't contain any valid fileBytes but nil."))
	}

	return errbyte.New.Results.ValuesOnly(fileBytes)
}
