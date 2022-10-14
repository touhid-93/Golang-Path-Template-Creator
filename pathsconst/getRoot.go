package pathsconst

import (
	"path/filepath"
	"runtime"
)

func getRoot() string {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	return filepath.Join(
		filepath.Dir(b),
		RootRelativeDir)
}
