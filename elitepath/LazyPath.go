package elitepath

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type LazyPath struct {
	fileMode               os.FileMode
	isDir, isFile, isExist issetter.Value
	Path                   *Path
	locationInfo           *pathhelper.LocationInfo
	fileInfo               os.FileInfo
	pathExistStat          *chmodhelper.PathExistStat
	simpleStat             *pathchmod.SimpleStat
	internalError          error
	internalErrorWrapper   *errorwrapper.Wrapper
}

func NewLazyPath(path *Path) *LazyPath {
	return &LazyPath{
		Path: path,
	}
}

func (it *LazyPath) FileMode() os.FileMode {
	if it.fileMode != 0 {
		return it.fileMode
	}

	stat := it.PathExistStat()
	if stat.HasFileInfo() {
		it.fileMode = stat.FileInfo.Mode()
	}

	return it.fileMode
}

func (it *LazyPath) HasFileInfo() bool {
	return it.FileInfo() != nil
}

func (it *LazyPath) Error() error {
	if it.internalError != nil {
		return it.internalError
	}

	if it.internalError != nil &&
		it.
			PathExistStat().
			IsEmptyError() {
		return nil
	}

	it.internalError = it.
		PathExistStat().
		MeaningFullError()

	return it.internalError
}

func (it *LazyPath) ErrorWrapper() *errorwrapper.Wrapper {
	if it.internalErrorWrapper != nil {
		return it.internalErrorWrapper
	}

	it.internalErrorWrapper = errnew.
		Path.
		Error(
			errtype.PathMissingOrInvalid,
			it.Error(),
			it.Path.CompiledPath())

	return it.internalErrorWrapper
}

func (it *LazyPath) FileInfo() os.FileInfo {
	if it.fileInfo != nil {
		return it.fileInfo
	}

	it.fileInfo = it.
		PathExistStat().
		FileInfo

	return it.fileInfo
}

func (it *LazyPath) SimpleStat() *pathchmod.SimpleStat {
	if it.simpleStat != nil {
		return it.simpleStat
	}

	it.simpleStat = it.Path.SimpleStat()

	return it.simpleStat
}

func (it *LazyPath) PathExistStat() *chmodhelper.PathExistStat {
	if it.pathExistStat != nil {
		return it.pathExistStat
	}

	it.pathExistStat = it.Path.ExistStat()

	return it.pathExistStat
}

func (it *LazyPath) LocationInfo() *pathhelper.LocationInfo {
	if it.locationInfo != nil {
		return it.locationInfo
	}

	it.locationInfo = it.Path.LocationInfo()

	return it.locationInfo
}

func (it *LazyPath) IsDir() bool {
	if it.isDir.IsUninitialized() {
		it.isDir = issetter.GetBool(it.PathExistStat().IsDir())
	}

	return it.isDir.IsTrue()
}

func (it *LazyPath) IsFile() bool {
	if it.isFile.IsUninitialized() {
		it.isFile = issetter.GetBool(it.PathExistStat().IsFile())
	}

	return it.isFile.IsTrue()
}

func (it *LazyPath) IsExist() bool {
	if it.isExist.IsUninitialized() {
		it.isExist = issetter.GetBool(it.PathExistStat().IsExist)
	}

	return it.isExist.IsTrue()
}

func (it *LazyPath) IsExistButFile() bool {
	return it.IsExist() && it.IsFile()
}

func (it *LazyPath) IsExistButDir() bool {
	return it.IsExist() && it.IsDir()
}

func (it *LazyPath) Clone() *LazyPath {
	if it == nil {
		return nil
	}

	return &LazyPath{
		Path: it.Path.ClonePath(),
	}
}
