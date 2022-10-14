package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinBaseDirWithSep isNormalizePlusLongPathFix if true then for windows add UNC Location fix
//
// Omits baseDir if not given
func JoinBaseDirWithSep(
	isSkipEmpty,
	isExpandEnvVariables,
	isNormalizePlusLongPathFix bool,
	sep string,
	baseDir string,
	paths ...string,
) string {
	if len(paths) == 0 && baseDir == "" {
		return constants.EmptyString
	}

	if len(paths) == 0 && baseDir != "" {
		return JoinConditionalNormalizedExpandIf(
			isNormalizePlusLongPathFix,
			isExpandEnvVariables,
			baseDir,
			constants.EmptyString)
	}

	if isSkipEmpty {
		paths = stringslice.
			NonEmptySlice(paths)
	}

	finalPath := strings.Join(
		paths,
		sep)

	if baseDir != "" {
		finalPath = baseDir +
			sep +
			finalPath
	}

	expand := expandpath.ExpandVariablesIf(
		isExpandEnvVariables,
		finalPath)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
