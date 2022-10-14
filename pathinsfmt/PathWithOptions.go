package pathinsfmt

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type PathWithOptions struct {
	IsNormalize       bool   `json:"IsNormalize,omitempty"`
	IsRecursive       bool   `json:"IsRecursive,omitempty"`
	IsExpandEnvVar    bool   `json:"IsExpandEnvVar,omitempty"`
	IsSkipInvalid     bool   `json:"IsSkipInvalid,omitempty"`
	Path              string `json:"Path"` // warning : use CompiledPath
	compiledPath      *string
	compiledPathSlice []string
}

func (it *PathWithOptions) CompiledPath() string {
	if it.compiledPath != nil {
		return *it.compiledPath
	}

	compiledPath := pathjoin.FixPath(
		it.IsNormalize,
		it.IsExpandEnvVar,
		it.Path)

	it.compiledPath = &compiledPath

	return *it.compiledPath
}

func (it *PathWithOptions) CompiledPathAsSlice() []string {
	if it.compiledPathSlice != nil {
		return it.compiledPathSlice
	}

	it.compiledPathSlice = []string{it.CompiledPath()}

	return it.compiledPathSlice
}

func (it *PathWithOptions) HasPath() bool {
	return it != nil && it.Path != constants.EmptyString
}

func (it *PathWithOptions) IsInvalid() bool {
	existStat := it.ExistStat()

	return !existStat.HasFileInfo()
}

// IsPathExist returns true if exist on the file system
func (it *PathWithOptions) IsPathExist() bool {
	return it.HasPath() && chmodhelper.IsPathExists(it.CompiledPath())
}

func (it *PathWithOptions) ClonePath() *PathWithOptions {
	if it == nil {
		return nil
	}

	return &PathWithOptions{
		IsNormalize:    it.IsNormalize,
		IsRecursive:    it.IsRecursive,
		IsExpandEnvVar: it.IsExpandEnvVar,
		IsSkipInvalid:  it.IsSkipInvalid,
		Path:           it.Path,
	}
}

func (it *PathWithOptions) IsEqual(another *PathWithOptions) bool {
	if !it.IsEqualWithoutOptions(another) {
		return false
	}

	return it.IsNormalize == another.IsNormalize &&
		it.IsRecursive == another.IsRecursive &&
		it.IsExpandEnvVar == another.IsExpandEnvVar
}

func (it *PathWithOptions) IsEqualWithoutOptions(another *PathWithOptions) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.Path == another.Path
}

func (it *PathWithOptions) ExistStat() *chmodhelper.PathExistStat {
	return chmodhelper.GetPathExistStat(it.CompiledPath())
}

func (it *PathWithOptions) FileInfo() os.FileInfo {
	fileInfo, _ := os.Stat(it.CompiledPath())

	return fileInfo
}

func (it *PathWithOptions) FileMode() os.FileMode {
	fileInfo, _ := os.Stat(it.CompiledPath())

	return fileInfo.Mode()
}

func (it *PathWithOptions) SafeFileMode() os.FileMode {
	fileInfo, _ := os.Stat(it.CompiledPath())

	if fileInfo != nil {
		return fileInfo.Mode()
	}

	return constants.Zero
}

func (it *PathWithOptions) IsDir() bool {
	existStat := it.ExistStat()

	return existStat.IsDir()
}

func (it *PathWithOptions) IsInvalidPath() bool {
	existStat := it.ExistStat()

	return !existStat.IsExist
}

func (it *PathWithOptions) IsExistButDir() bool {
	existStat := it.ExistStat()

	return existStat.IsExist && existStat.IsDir()
}

func (it *PathWithOptions) IsExistButFile() bool {
	existStat := it.ExistStat()

	return existStat.IsExist && existStat.IsFile()
}

func (it *PathWithOptions) IsFile() bool {
	existStat := it.ExistStat()

	return existStat.IsFile()
}
