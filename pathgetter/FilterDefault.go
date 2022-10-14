package pathgetter

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func FilterDefault(
	isNormalize bool,
	isIgnoreOnError bool,
	rootPath string,
	filter pathfuncs.Filter,
) []*pathfuncs.FilterResult {
	return Filter(
		osconsts.PathSeparator,
		rootPath,
		isNormalize,
		isIgnoreOnError,
		filter)
}
