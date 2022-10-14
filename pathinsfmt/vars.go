package pathinsfmt

import (
	"sync"

	"gitlab.com/evatix-go/pathhelper/copyrecursive"
)

var (
	lockerMutex                 = sync.Mutex{}
	defaultRecursiveCopyOptions = &copyrecursive.Options{
		IsSkipOnExist:      false,
		IsRecursive:        false,
		IsMove:             false,
		IsClearDestination: false,
		IsUseShellOrCmd:    false,
		IsNormalize:        false,
		IsExpandVar:        false,
	}
)
