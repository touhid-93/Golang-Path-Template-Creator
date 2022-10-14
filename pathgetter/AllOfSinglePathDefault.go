package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func AllOfSinglePathDefault(
	isNormalize bool,
	exploringPath string,
) *errstr.Results {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		osconsts.PathSeparator,
		exploringPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return errstr.New.Results.ErrorWrapper(
			errnew.Path.
				Error(
					errtype.PathExpand,
					err,
					rootPath2))
	}

	results := make([]string, len(allPaths))

	for i, fileInfo := range allPaths {
		results[i] = rootPath2 +
			osconsts.PathSeparator +
			fileInfo.Name()
	}

	return errstr.New.Results.Strings(results)
}
