package pathrecurseinfo

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Instruction struct {
	Root               string
	ExcludingRootNames []string //  root file or dir names contains in this will be ignored
	ExcludingPaths     []string //  path contains in this will be ignored
	IsRecursive,       // Recursively get paths if dir
	IsRelativePath, // remove root path from paths
	IsIncludeAll, // includes dir, files all
	IsIncludeDirsOnly, // includes only dir if IsIncludeAll false
	IsIncludeFilesOnly, //  includes only files if IsIncludeAll false
	IsNormalize,
	IsExpandEnvironmentVar, // Expand environment variable
	IsExcludeRoot bool // Don't include root path
	excludingNamesHashset, excludingPathsHashset *corestr.Hashset
}

func (it *Instruction) Result() *Result {
	return GetInstructionResult(it)
}

func (it *Instruction) SliceResult() *corestr.SimpleSlice {
	if it.Root == "" {
		return corestr.Empty.SimpleSlice()
	}

	rs := it.Result()

	if rs.IsInvalidResult || rs.IsEmpty() {
		rs.Dispose()

		return corestr.Empty.SimpleSlice()
	}

	rs.ErrorWrapper.Dispose()

	return rs.
		PathsResult.
		ExpandingPaths
}

func (it *Instruction) StringsResults() *errstr.Results {
	rs := it.Result()

	return rs.StringsResults()
}

func (it *Instruction) StringsResultsWithoutUnc() *errstr.Results {
	rs := it.Result()

	results := rs.StringsResults()

	if results.HasSafeItems() {
		results.Values = normalize.TrimPrefixUncPaths(
			true,
			results.Values...)
	}

	return results
}

func (it *Instruction) HasExcludingRootNames() bool {
	return len(it.ExcludingRootNames) > 0
}

func (it *Instruction) HasExcludingPaths() bool {
	return len(it.ExcludingPaths) > 0
}

func (it *Instruction) ExcludingNamesHashset() *corestr.Hashset {
	if it.excludingNamesHashset != nil {
		return it.excludingNamesHashset
	}

	slicePtr := stringslice.SlicePtr(it.ExcludingRootNames)
	it.excludingNamesHashset = corestr.New.Hashset.StringsPtr(
		slicePtr)

	return it.excludingNamesHashset
}

func (it *Instruction) ExcludingPathsHashset() *corestr.Hashset {
	if it.excludingPathsHashset != nil {
		return it.excludingPathsHashset
	}

	slicePtr := stringslice.SlicePtr(it.ExcludingPaths)
	it.excludingPathsHashset = corestr.New.Hashset.StringsPtr(
		slicePtr)

	return it.excludingPathsHashset
}
