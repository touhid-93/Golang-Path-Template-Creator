package checksummer

import (
	"runtime"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaultcapacity"

	"gitlab.com/evatix-go/pathhelper/hashas"
)

// hashAllAsync reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, hashAllAsync returns an error.  In that case,
// hashAllAsync does not wait for inflight read operations to complete.
func hashAllAsync(
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) (map[string][]byte, error) {
	// hashAllAsync closes the done channel when it returns; it may do so before
	// receiving all the values from results and errc.
	done := make(chan struct{})
	defer close(done)

	var paths <-chan string
	var errc <-chan error

	if isRecursive {
		paths, errc = walkFilesAsync(done, root)
	} else {
		paths, errc = walkFilesAsyncNonRecursive(done, root)
	}

	// Start a fixed number of goroutines to read and digest files.
	results := make(chan walkResult)
	var wg sync.WaitGroup
	var numDigesters = runtime.NumCPU() * 2
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			// creating new for each is a must
			// otherwise race condition
			hasher, _ := hashType.NewHash()
			digester(hasher, done, paths, results)
			wg.Done()
		}()
	}

	// the `range results` for loop will block
	// we are asynchronously waiting for all the worker to finish
	// after they are finished, we are going to close the results
	// channel, which will allow to finish the `range results` for
	// loop bellow
	go func() {
		wg.Wait()
		close(results)
	}()

	fileHashes := make(
		map[string][]byte,
		defaultcapacity.PredictiveDefault(constants.ArbitraryCapacity50),
	)

	for result := range results {
		if result.err != nil {
			return nil, result.err
		}

		fileHashes[result.path] = result.sum
	}

	// Check whether the Walk failed.
	if err := <-errc; err != nil {
		return nil, err
	}

	return fileHashes, nil
}
