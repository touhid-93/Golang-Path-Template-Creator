package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func WithTempTest(paths ...string) string {
	return normalize.JoinNormalizedPaths(
		pathsconst.DefaultTempTestDir,
		paths...)
}
