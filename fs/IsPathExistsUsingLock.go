package fs

import "os"

func IsPathExistsUsingLock(location string) bool {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	_, err := os.Stat(location)

	return err == nil || !os.IsNotExist(err)
}
