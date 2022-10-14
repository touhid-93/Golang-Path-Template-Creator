package pathgetterinternal

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
)

// GetAllFiles only gives files not nested files
func GetAllFiles(
	isFixPaths bool,
	separator,
	rootPath string,
) *errstr.Results {
	if rootPath == constants.EmptyString {
		return errstr.Empty.Results()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return errstr.New.Results.ErrorWrapper(
			errnew.
				Path.
				Error(
					errtype.PathStatusCannotRead,
					err,
					rootPath))
	}

	slice := make(
		[]string,
		constants.Zero,
		len(fileInfos))

	for _, info := range fileInfos {
		currentPath := rootPath +
			separator +
			info.Name()

		currentPath = normalizeinternal.JoinPathsFixIf(
			isFixPaths, currentPath)

		if info == nil {
			continue
		}

		if info.IsDir() {
			continue
		}

		slice = append(
			slice,
			currentPath)
	}

	return errstr.New.Results.Strings(
		slice)
}
