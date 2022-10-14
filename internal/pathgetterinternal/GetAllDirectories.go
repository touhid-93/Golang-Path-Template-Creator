package pathgetterinternal

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
)

// GetAllDirectories all directories but not nested directories
func GetAllDirectories(
	isFixPaths bool,
	separator,
	rootPath string,
) *errstr.Results {
	if rootPath == "" {
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

		if !IsDirectory(currentPath) {
			continue
		}

		slice = append(
			slice, currentPath)
	}

	return errstr.New.Results.Strings(
		slice)
}
