package pathfilter

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// GetFiltersArray
//
// rootPath should not ends with separator
func GetFiltersArray(
	rootPath string,
	additionalFilters []string,
	extensions []string,
	isRecursive bool,
	isNormalizePath bool,
) []string {
	separator := osconsts.PathSeparator
	rootPath2 := normalize.PathUsingSeparatorIf(
		isNormalizePath,
		isNormalizePath,
		isNormalizePath,
		separator,
		rootPath,
	)

	additionalFilterLength :=
		len(additionalFilters)

	extensionsLength :=
		len(extensions)

	arg := &recursiveFilterGetterParam{
		separator:               separator,
		rootPath:                rootPath2,
		eachFilterPath:          "", // will be generated in each case
		rootPathPlusSeparator:   rootPath2 + separator,
		extensionsLength:        extensionsLength,
		additionalFiltersLength: additionalFilterLength,
		additionalFilters:       additionalFilters,
		extensions:              extensions,
	}

	if additionalFilterLength == 0 && isRecursive {
		arg.eachFilterPath = rootPath2

		return getRecursiveFilterForEachFilterPath(arg)
	}

	if additionalFilterLength == 0 && !isRecursive {
		arg.eachFilterPath = rootPath2

		return getFilterForEachFilterPath(arg)
	}

	if isRecursive {
		return getRecursiveFilterDefinitions(arg)
	}

	return getFilterDefinitions(arg)
}
