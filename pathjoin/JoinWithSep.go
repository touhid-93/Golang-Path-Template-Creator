package pathjoin

import (
	"gitlab.com/evatix-go/core/constants"
)

// JoinWithSep isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinWithSep(
	isSkipEmpty,
	isExpandEnvVariables,
	isNormalizePlusLongPathFix bool,
	sep string,
	relativePaths ...string,
) string {
	if len(relativePaths) == 0 {
		return constants.EmptyString
	}

	return JoinWithBaseDirSep(
		isSkipEmpty,
		isExpandEnvVariables,
		isNormalizePlusLongPathFix,
		sep,
		constants.EmptyString,
		relativePaths...)
}
