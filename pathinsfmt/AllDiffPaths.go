package pathinsfmt

import (
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type AllDiffPaths struct {
	lazyFlatPaths    []string
	Paths            []string `json:"Locations"`
	IsNormalizeApply bool     `json:"IsNormalizeApply"`
}

func (allDiffPaths *AllDiffPaths) Length() int {
	return len(allDiffPaths.Paths)
}

func (allDiffPaths *AllDiffPaths) IsEmpty() bool {
	return allDiffPaths.Length() == 0
}

func (allDiffPaths *AllDiffPaths) HasAnyItem() bool {
	return allDiffPaths.Length() > 0
}

func (allDiffPaths *AllDiffPaths) LazyFlatPaths() []string {
	if allDiffPaths.lazyFlatPaths != nil {
		return allDiffPaths.lazyFlatPaths
	}

	allDiffPaths.lazyFlatPaths = allDiffPaths.FlatPaths()

	return allDiffPaths.lazyFlatPaths
}

func (allDiffPaths *AllDiffPaths) FlatPaths() []string {
	return normalize.PathsOnConditions(
		allDiffPaths.IsNormalizeApply,
		allDiffPaths.Paths)
}
