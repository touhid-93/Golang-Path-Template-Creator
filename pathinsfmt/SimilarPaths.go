package pathinsfmt

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type SimilarPaths struct {
	RootPath         string   `json:"RootPath"`
	RelativePaths    []string `json:"RelativePaths"`
	IsNormalizeApply bool     `json:"IsNormalizeApply"`
}

func (it *SimilarPaths) BasePathsCreator() *BasePathsCreator {
	return &BasePathsCreator{
		RootDir:     it.RootPath,
		Files:       it.RelativePaths,
		IsNormalize: it.IsNormalizeApply,
	}
}

func (it *SimilarPaths) Length() int {
	return len(it.RelativePaths)
}

func (it *SimilarPaths) IsEmpty() bool {
	return it.Length() == 0
}

func (it *SimilarPaths) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *SimilarPaths) ApplyLinuxRecursiveFileModeOnRoot(
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	return pathchmod.ApplyLinuxRecursiveChmodOnPathUsingFileMode(
		fileMode,
		it.RootPath)
}

func (it *SimilarPaths) FlatPaths() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make(
		[]string,
		it.Length())

	root := it.RootPath

	for i, relativePath := range it.RelativePaths {
		joinedPath := pathjoin.JoinNormalizedIf(
			it.IsNormalizeApply,
			root,
			relativePath)

		slice[i] = joinedPath
	}

	return slice
}
