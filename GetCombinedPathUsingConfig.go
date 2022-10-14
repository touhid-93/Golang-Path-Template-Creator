package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// if pathConfig is nil then pathConfig gets created pathhelpercore.NewDefaultPathConfig()
// By NewDefaultPathConfig:
// Separator = os.PathSeparator
// IsNormalize = true
// IsIgnoreEmptyPath = false
func GetCombinedPathUsingConfig(pathConfig *pathhelpercore.PathConfig, paths ...string) string {
	return getCombinedPathUsingConfigInternal(
		pathConfig,
		paths)
}
