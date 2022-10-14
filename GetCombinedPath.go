package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// GetCombinedPath
//
// @isIgnoreEmptyPath if true then ignore empty string (nil, "", or any empty spaces "  ")
func GetCombinedPath(
	separator string,
	isIgnoreEmptyPath bool,
	isLongPathFix bool,
	isNormalize bool,
	paths ...string,
) string {
	pathConfig := &pathhelpercore.PathConfig{
		Separator:         separator,
		IsNormalize:       isNormalize,
		IsLongPathFix:     isLongPathFix,
		IsIgnoreEmptyPath: isIgnoreEmptyPath,
	}

	return getCombinedPathUsingConfigInternal(
		pathConfig,
		paths)
}
