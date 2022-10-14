package pathinsfmt

import (
	"sort"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreinstruction"
)

type GenericPathsCollection struct {
	Specification                      *coreinstruction.Specification `json:"Specification,omitempty"`
	SimilarPaths                       []SimilarPaths                 `json:"SimilarPaths,omitempty"`
	AllDiffPaths                       []AllDiffPaths                 `json:"AllDiffPaths,omitempty"`
	DynamicPaths                       *DynamicPaths                  `json:"DynamicPaths,omitempty"`
	lazyFlatPaths, lazyFlatPathsSorted *[]string
}

// Length of len(receiver.SimilarPaths) +
// len(receiver.AllDiffPaths) +
// items in DynamicPaths (not all specific paths)
func (it *GenericPathsCollection) Length() int {
	length := len(it.SimilarPaths) +
		len(it.AllDiffPaths)

	if it.DynamicPaths == nil {
		return length
	}

	return length + it.DynamicPaths.Length()
}

func (it *GenericPathsCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *GenericPathsCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *GenericPathsCollection) LazyFlatPathsSorted() []string {
	if it.lazyFlatPathsSorted != nil {
		return *it.lazyFlatPathsSorted
	}

	lazyPaths := it.LazyFlatPaths()
	sort.Strings(lazyPaths)

	it.lazyFlatPathsSorted = &lazyPaths

	return *it.lazyFlatPaths
}

func (it *GenericPathsCollection) LazyFlatPaths() []string {
	if it.lazyFlatPaths != nil {
		return *it.lazyFlatPaths
	}

	flatPaths := it.FlatPaths()
	it.lazyFlatPaths = &flatPaths

	return flatPaths
}

func (it *GenericPathsCollection) IsEmptySimilarPaths() bool {
	return it.SimilarPaths == nil || len(it.SimilarPaths) == 0
}

func (it *GenericPathsCollection) SimilarPathsIndividualItemsLength() int {
	length := 0

	if it.SimilarPaths == nil {
		return 0
	}

	for _, similarPaths := range it.SimilarPaths {
		length += similarPaths.Length()
	}

	return length
}

func (it *GenericPathsCollection) SimilarPathsFlatPaths() []string {
	if it.IsEmptySimilarPaths() {
		return []string{}
	}

	slice := make(
		[]string,
		constants.Zero,
		it.SimilarPathsIndividualItemsLength()+constants.ArbitraryCapacity10)

	for _, similarPaths := range it.SimilarPaths {
		if similarPaths.IsEmpty() {
			continue
		}

		slice = append(
			slice,
			similarPaths.FlatPaths()...)
	}

	return slice
}

func (it *GenericPathsCollection) IsEmptyAllDiffPaths() bool {
	return it.AllDiffPaths == nil || len(it.AllDiffPaths) == 0
}

func (it *GenericPathsCollection) AllDiffPathsIndividualItemsLength() int {
	length := 0

	if it.AllDiffPaths == nil {
		return 0
	}

	for _, allDiff := range it.AllDiffPaths {
		length += allDiff.Length()
	}

	return length
}

func (it *GenericPathsCollection) AllDiffPathsFlatPaths() []string {
	if it.IsEmptyAllDiffPaths() {
		return []string{}
	}

	slice := make(
		[]string,
		constants.Zero,
		it.AllDiffPathsIndividualItemsLength()+constants.ArbitraryCapacity10)

	for _, allDiffPaths := range it.AllDiffPaths {
		if allDiffPaths.IsEmpty() {
			continue
		}

		slice = append(
			slice,
			allDiffPaths.FlatPaths()...)
	}

	return slice
}

func (it *GenericPathsCollection) IsEmptyDynamicPaths() bool {
	return it.DynamicPaths == nil || len(it.DynamicPaths.AllDiffPaths) == 0
}

func (it *GenericPathsCollection) DynamicPathsIndividualItemsLength() int {
	length := 0

	if it.DynamicPaths == nil {
		return 0
	}

	for _, allDiff := range it.DynamicPaths.AllDiffPaths {
		length += allDiff.Length()
	}

	return length
}

func (it *GenericPathsCollection) DynamicPathsFlatPaths() []string {
	if it.IsEmptyDynamicPaths() {
		return []string{}
	}

	length := it.DynamicPathsIndividualItemsLength()

	slice := make(
		[]string,
		constants.Zero,
		length+constants.ArbitraryCapacity10)

	for _, allDiffPaths := range it.DynamicPaths.AllDiffPaths {
		if allDiffPaths.IsEmpty() {
			continue
		}

		slice = append(
			slice,
			allDiffPaths.FlatPaths()...)
	}

	return slice
}

func (it *GenericPathsCollection) LazyPathsIf(isLazyPaths bool) []string {
	if isLazyPaths {
		return it.LazyFlatPaths()
	}

	return it.FlatPaths()
}

func (it *GenericPathsCollection) FlatPaths() []string {
	collections := corestr.Empty.LinkedCollections()
	wg := &sync.WaitGroup{}

	wg.Add(constants.Capacity3)
	collections.AddAsyncFuncItems(
		wg,
		false,
		it.AllDiffPathsFlatPaths,
		it.DynamicPathsFlatPaths,
		it.SimilarPathsFlatPaths,
	)

	return collections.
		ToCollection(constants.Zero).
		ListStrings()
}

func (it *GenericPathsCollection) FlatPathsSorted() []string {
	flatPaths := it.
		FlatPaths()

	sort.Strings(
		flatPaths)

	return flatPaths
}
