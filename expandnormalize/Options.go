package expandnormalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Options struct {
	normalize.Options
	IsExpandEnvVar bool
}

func (it Options) Fix(location string) string {
	expandedPath := expandpath.ExpandVariablesIf(
		it.IsExpandEnvVar,
		location)

	return it.Options.FixPath(expandedPath)
}

func (it Options) JoinPath(path1, path2 string) string {
	if path1 == constants.EmptyString && path2 == constants.EmptyString {
		return constants.EmptyString
	}

	if path1 == constants.EmptyString {
		return it.FixPath(path2)
	}

	if path2 == constants.EmptyString {
		return it.FixPath(path1)
	}

	join := path1 +
		osconsts.PathSeparator +
		path2

	return it.FixPath(join)
}

func (it Options) JoinPath3(path1, path2, path3 string) string {
	joined := normalize.SimpleJoinPath3(path1, path2, path3)

	return it.FixPath(joined)
}

func (it Options) JoinPaths(locations ...string) string {
	joined := normalize.SimpleJoinPaths(locations...)

	return it.FixPath(joined)
}

func (it Options) JoinWithBaseDirPaths(
	baseDir string, locations ...string,
) string {
	joined := normalize.SimpleBaseDirJoinPaths(
		baseDir,
		locations...)

	return it.FixPath(joined)
}
