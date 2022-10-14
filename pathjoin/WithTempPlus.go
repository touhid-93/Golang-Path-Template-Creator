package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

// WithTempPlus
//
//  skip on empty path and others optional
//
//  Windows
//      - "%temp%\locations\..."
//  Unix
//      - "/tmp/locations/..."
func WithTempPlus(
	isNormalizeLongPathFix,
	isExpandEnvVar bool,
	baseDir string,
	locations ...string,
) string {
	return JoinBaseDirWithSep(
		true,
		isExpandEnvVar,
		isNormalizeLongPathFix,
		osconsts.PathSeparator,
		JoinSimple(pathsconst.TempPermanentDir, baseDir),
		locations...)
}
