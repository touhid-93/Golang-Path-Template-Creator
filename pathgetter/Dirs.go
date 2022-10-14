package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Dirs(
	separator,
	rootPath string,
	isNormalize bool,
) *errstr.Results {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return errstr.New.Results.ErrorWrapper(errnew.
			Path.
			Error(
				errtype.PathExpand,
				err,
				rootPath2))
	}

	results := make([]string, 0, len(allPaths))

	for _, fileInfo := range allPaths {
		if !fileInfo.IsDir() {
			continue
		}

		combinedPath := rootPath2 +
			separator +
			fileInfo.Name()

		results = append(results, combinedPath)
	}

	return errstr.New.Results.Strings(results)
}
