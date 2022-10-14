package pathinsfmt

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/coreindexes"
)

type DynamicPaths struct {
	lazyFlatPaths []string
	Vars          []PathVar      `json:"Vars,omitempty"`
	AllDiffPaths  []AllDiffPaths `json:"AllDiffPaths,omitempty"`
}

func (dynamicPaths *DynamicPaths) LazyFlatPaths() []string {
	if dynamicPaths.lazyFlatPaths != nil {
		return dynamicPaths.lazyFlatPaths
	}

	dynamicPaths.lazyFlatPaths = dynamicPaths.FlatPaths()

	return dynamicPaths.lazyFlatPaths
}

func (dynamicPaths *DynamicPaths) EachItemsLength() int {
	length := 0

	if dynamicPaths.AllDiffPaths == nil {
		return length
	}

	for _, allDiffPaths := range dynamicPaths.AllDiffPaths {
		length += allDiffPaths.Length()
	}

	return length
}

func (dynamicPaths *DynamicPaths) Length() int {
	if dynamicPaths == nil {
		return constants.Zero
	}

	return len(dynamicPaths.AllDiffPaths)
}

func (dynamicPaths *DynamicPaths) IsEmpty() bool {
	if dynamicPaths == nil {
		return false
	}

	return dynamicPaths.Length() == 0
}

// FlatPaths don't apply PathVar at this moment, feature is complicated and not implemented yet.
//
// Current flat paths only returns all the paths combined.
func (dynamicPaths *DynamicPaths) FlatPaths() []string {
	if dynamicPaths.IsEmpty() {
		return []string{}
	}

	length := dynamicPaths.Length()

	if length == 1 {
		return dynamicPaths.
			AllDiffPaths[coreindexes.First].
			FlatPaths()
	}

	if length == 2 {
		items1 := dynamicPaths.
			AllDiffPaths[coreindexes.First].
			FlatPaths()

		items2 := dynamicPaths.
			AllDiffPaths[coreindexes.Second].
			FlatPaths()

		return *stringslice.MergeNewSlicesPtrOfSlices(
			&items1,
			&items2)
	}

	collectionOfCollection :=
		corestr.Empty.LinkedCollections()

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(length)

	var asyncAdd = func(diffPath AllDiffPaths) {
		filePaths := diffPath.FlatPaths()

		mutex.Lock()
		defer mutex.Unlock()

		collectionOfCollection.AddStringsPtr(
			false,
			&filePaths,
		)

		wg.Done()
	}

	for _, diffPath := range dynamicPaths.AllDiffPaths {
		go asyncAdd(diffPath)
	}

	wg.Wait()

	collection := collectionOfCollection.
		ToCollection(constants.Zero)

	return collection.ListStrings()
}
