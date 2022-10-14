package recursivepaths

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func FilesExtensionFilter(
	isRecursive bool,
	rootPath string,
	filteringDotExtensions ...string,
) *errstr.Results {
	if len(filteringDotExtensions) == 0 {
		return errstr.Empty.Results()
	}

	extensionsHashset := corestr.New.Hashset.StringsPtr(
		&filteringDotExtensions)

	simpleFilterFunc := func(fullPath string) (isTake bool, err error) {
		extension := filepath.Ext(fullPath)
		isTake = extensionsHashset.Has(extension)

		return isTake, nil
	}

	return SimpleFilterFilesOptionsIf(
		isRecursive,
		false,
		false,
		simpleFilterFunc,
		rootPath)
}
