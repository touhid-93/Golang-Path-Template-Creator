package pathjoin

import (
	"path/filepath"
	"runtime"
)

func FilePathOfCaller(callerIndex int) string {
	_, b, _, _ := runtime.Caller(callerIndex)

	return filepath.Dir(b)
}
