package pathfilter

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/pathgetter"
)

//goland:noinspection GoNilness
func getFilterForEachFilterPath(
	arg *recursiveFilterGetterParam,
) []string {
	if arg == nil {
		errtype.Null.
			PanicNoRefs("args")
	}

	allDirs := pathgetter.Dirs(
		arg.separator,
		arg.eachFilterPath,
		false)

	allDirs.ErrorWrapper.HandleError()
	length := allDirs.Length() * arg.extensionsLength
	newFilters := make([]string, length)
	i := 0

	for _, s := range allDirs.Values {
		// removing starting root, only the filters needed
		for i2, ext := range arg.extensions {
			newFilters[i+i2] = strings.Replace(
				s,
				arg.rootPathPlusSeparator,
				constants.EmptyString,
				constants.One) +
				arg.separator +
				ext
		}

		i++
	}

	return newFilters
}
