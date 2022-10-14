package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func Filter(
	separator, rootPath string,
	isNormalize bool,
	isIgnoreOnError bool,
	filter pathfuncs.Filter,
) []*pathfuncs.FilterResult {
	compiledRootPath := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(compiledRootPath)

	if err != nil {
		empty := make([]*pathfuncs.FilterResult, 0)

		return empty
	}

	results := make([]*pathfuncs.FilterResult, 0, len(allPaths))

	for _, fileInfo := range allPaths {
		if fileInfo.IsDir() {
			continue
		}

		isDir := fileInfo.IsDir()
		name := fileInfo.Name()
		combinedPath := compiledRootPath +
			separator +
			name

		arg := &pathfuncs.FilterArg{
			RootPath:    rootPath,
			FileName:    name,
			FullPath:    combinedPath,
			IsFile:      !isDir,
			IsDirectory: isDir,
			FileInfo:    fileInfo,
		}

		result := filter(arg)
		hasError := result.ErrorWrapper.HasError()

		if hasError && isIgnoreOnError {
			continue
		} else {
			result.ErrorWrapper.HandleError()
		}

		if result.IsTake {
			results = append(results, result)
		}

		if result.IsBreak {
			return results
		}
	}

	return results
}
