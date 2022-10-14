package pathjoin

import (
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/core/codestack"
)

func CurrentDirectory() string {
	_, b, _, _ := runtime.Caller(codestack.Skip1)

	return filepath.Dir(b)
}
