package pathsconst

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func getExecutableDirectory() string {
	exe, _ := os.Executable()
	exeDir, _ := splitinternal.GetWithoutSlash(exe)

	return exeDir
}
