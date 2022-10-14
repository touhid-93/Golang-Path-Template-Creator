package pathrecurseinfo

import (
	"gitlab.com/evatix-go/core/corecmp"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type PathsResult struct {
	ExpandingPaths *corestr.SimpleSlice
	IsExist,
	IsFile,
	IsDir bool
}

func (it *PathsResult) IsEqual(another *PathsResult) bool {
	if !it.IsEqualOnlyItems(another) {
		return false
	}

	if it.IsExist != another.IsExist {
		return false
	}

	if it.IsDir != another.IsDir {
		return false
	}

	if it.IsFile != another.IsFile {
		return false
	}

	return true
}

func (it *PathsResult) IsEqualOnlyItems(another *PathsResult) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.ExpandingPaths.Length() != another.ExpandingPaths.Length() {
		return false
	}

	return corecmp.IsStringsEqual(
		it.ExpandingPaths.Items,
		another.ExpandingPaths.Items)
}

func (it *PathsResult) IsDistinctEqualItems(another *PathsResult) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	current := it.ExpandingPaths.Hashset()
	anotherHashset := another.ExpandingPaths.Hashset()

	return current.IsEqualsPtr(anotherHashset)
}

func (it *PathsResult) JoinWithRoot(
	isFixAll,
	isNormalize bool,
	root string,
) *corestr.SimpleSlice {
	if it.ExpandingPaths.IsEmpty() {
		return corestr.New.SimpleSlice.Cap(0)
	}

	rootFix := normalize.PathUsingSingleIf(
		isNormalize,
		root)
	newSlice := make(
		[]string,
		it.ExpandingPaths.Length())

	isLibFunc := !isNormalize

	for i, item := range it.ExpandingPaths.Items {
		newPath := pathjoin.JoinSimpleIf(
			isLibFunc,
			rootFix,
			item)

		if isFixAll {
			newPath = normalize.Path(
				newPath)
		}

		newSlice[i] = newPath
	}

	return corestr.New.SimpleSlice.Strings(
		newSlice)
}

func (it *PathsResult) Clone(isDeepClone bool) *PathsResult {
	if it == nil {
		return nil
	}

	return &PathsResult{
		ExpandingPaths: corestr.New.SimpleSlice.Direct(
			isDeepClone,
			it.ExpandingPaths.Items),
		IsExist: it.IsExist,
		IsFile:  it.IsFile,
		IsDir:   it.IsDir,
	}
}

func (it *PathsResult) ConcatNew(
	isDeepClone,
	isClone bool,
	other *PathsResult,
) *PathsResult {
	if other == nil && !isClone {
		return it
	}

	if other == nil && isClone {
		return it.Clone(isDeepClone)
	}

	slice := it.ExpandingPaths.ConcatNewSimpleSlices(
		other.ExpandingPaths)

	return &PathsResult{
		ExpandingPaths: slice,
		IsExist:        it.IsExist && other.IsExist,
		IsFile:         it.IsFile && other.IsFile,
		IsDir:          it.IsDir && other.IsDir,
	}
}

func (it *PathsResult) Dispose() {
	if it == nil {
		return
	}

	it.ExpandingPaths.Dispose()
	it.ExpandingPaths = nil
}
