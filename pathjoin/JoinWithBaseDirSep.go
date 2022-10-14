package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinWithBaseDirSep
//
//  isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinWithBaseDirSep(
	isSkipEmpty,
	isExpandEnvVariables,
	isNormalizePlusLongPathFix bool,
	sep string,
	baseDir string,
	relativePaths ...string,
) string {
	if len(relativePaths) == 0 {
		return constants.EmptyString
	}

	if isSkipEmpty {
		relativePaths = stringslice.
			NonEmptySlice(relativePaths)
	}

	finalPath := strings.Join(
		relativePaths,
		sep)

	if baseDir != "" {
		finalPath = baseDir + sep + finalPath
	}

	expand := expandpath.ExpandVariablesIf(
		isExpandEnvVariables,
		finalPath)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
