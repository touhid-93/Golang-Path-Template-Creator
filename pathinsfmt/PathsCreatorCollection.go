package pathinsfmt

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type PathsCreatorCollection struct {
	PathsCreatorItems       []PathsCreator `json:"PathsCreatorItems,omitempty"`
	IsIgnoreOnExist         bool
	IsDeleteAllBeforeCreate bool
	lazyFlatPaths           []string
}

func (it *PathsCreatorCollection) LazyFlatPathsIf(isLazy bool) []string {
	if isLazy {
		return it.LazyFlatPaths()
	}

	return it.FlatPaths()
}

func (it *PathsCreatorCollection) LazyFlatPaths() []string {
	if it.lazyFlatPaths != nil {
		return it.lazyFlatPaths
	}

	it.lazyFlatPaths = it.FlatPaths()

	return it.lazyFlatPaths
}

// Length yields count of PathsCreatorItems, not all paths count
func (it *PathsCreatorCollection) Length() int {
	return len(it.PathsCreatorItems)
}

func (it *PathsCreatorCollection) IsEmpty() bool {
	return len(it.PathsCreatorItems) == 0
}

func (it *PathsCreatorCollection) HasAnyItem() bool {
	return len(it.PathsCreatorItems) > 0
}

func (it *PathsCreatorCollection) FlatPaths() []string {
	length := it.Length()

	if length == 0 {
		return []string{}
	}

	collection := corestr.Empty.LinkedCollections()

	wg3 := &sync.WaitGroup{}

	for _, createInstruction := range it.PathsCreatorItems {
		wg3.Add(1)
		collection.AddAsyncFuncItemsPointer(
			wg3,
			false,
			createInstruction.FlatPathsPtr)
	}

	return collection.
		ToCollection(constants.Zero).
		ListStrings()
}
