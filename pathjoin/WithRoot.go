package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func WithRoot(paths ...string) string {
	return JoinFix(pathsconst.RootDir, paths...)
}
