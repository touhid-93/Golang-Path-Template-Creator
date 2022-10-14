package normalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

type Options struct {
	IsNormalize,
	IsLongPathFix,
	IsForceLongPathFix bool
}

func NewOptionsAllTrue() *Options {
	return &Options{
		IsNormalize:        true,
		IsLongPathFix:      true,
		IsForceLongPathFix: true,
	}
}

func (it Options) IsAllOptionsDisabled() bool {
	return !(it.IsLongPathFix ||
		it.IsForceLongPathFix ||
		it.IsNormalize)
}

func (it Options) FixPath(path string) string {
	if path == constants.EmptyString {
		return path
	}

	if it.IsAllOptionsDisabled() {
		return path
	}

	return PathUsingSeparatorIf(
		it.IsForceLongPathFix,
		it.IsLongPathFix,
		it.IsNormalize,
		osconsts.PathSeparator,
		path)
}

func (it Options) JoinPath(path1, path2 string) string {
	joined := SimpleJoinPath(path1, path2)

	return it.FixPath(joined)
}

func (it Options) JoinPath3(path1, path2, path3 string) string {
	joined := SimpleJoinPath3(
		path1,
		path2,
		path3)

	return it.FixPath(joined)
}

func (it Options) JoinPaths(locations ...string) string {
	joined := SimpleJoinPaths(locations...)

	return it.FixPath(joined)
}

func (it Options) JoinWithBaseDirPaths(
	baseDir string, locations ...string,
) string {
	joined := SimpleBaseDirJoinPaths(
		baseDir,
		locations...)

	return it.FixPath(joined)
}
