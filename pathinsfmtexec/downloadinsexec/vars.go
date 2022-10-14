package downloadinsexec

import "sync"

var (
	mutexLocker = sync.Mutex{}
)
