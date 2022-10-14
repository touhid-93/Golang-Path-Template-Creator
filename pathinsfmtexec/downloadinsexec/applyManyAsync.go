package downloadinsexec

import (
	"sync"

	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

// applyManyAsync cannot exit on immediate exit
func applyManyAsync(
	errCollection *errwrappers.Collection,
	downloads *pathinsfmt.Downloads,
) (isSuccess bool) {
	if downloads.IsEmpty() {
		return true
	}

	stateTracker := errCollection.StateTracker()
	wg := sync.WaitGroup{}
	wg.Add(downloads.Length())

	applyAsyncFunc := func(index int) {
		err := Apply(&downloads.Downloads[index])

		mutexLocker.Lock()
		defer mutexLocker.Unlock()
		errCollection.AddWrapperPtr(err)
		wg.Done()
	}

	for i := range downloads.Downloads {
		go applyAsyncFunc(i)
	}

	wg.Wait()

	return stateTracker.IsSuccess()
}
