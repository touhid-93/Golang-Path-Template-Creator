package pathjoin

import (
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/core/codestack"
)

func DirByStackSkip(stackSkip int) string {
	_, b, _, _ := runtime.Caller(stackSkip + codestack.Skip1)

	return filepath.Dir(b)
}
