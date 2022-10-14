package fs

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
				Messages(
					errtype.ReadRequestFailed,
					filePath,
					"fs.ReadFile",
					err.Error()))
	}

	if fileBytes == nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.
				Path.
				Messages(
					errtype.EmptyContent,
					filePath,
					"fs.ReadFile",
					"Location doesn't contain any valid fileBytes but nil."))
	}

	return errbyte.New.Results.ValuesOnly(fileBytes)
}
