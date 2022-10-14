package fileinfo

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

type SimplePathWrapper struct {
	Path        string
	IsDirectory bool
}

func (it *SimplePathWrapper) IsEquals(anotherWrapper SimplePathWrapper) bool {
	return it.IsDirectory == anotherWrapper.IsDirectory &&
		it.Path == anotherWrapper.Path
}

func (it *SimplePathWrapper) IsEqualsPtr(anotherWrapper *SimplePathWrapper) bool {
	if anotherWrapper == nil {
		return false
	}

	if it == anotherWrapper {
		return true
	}

	return it.IsDirectory == anotherWrapper.IsDirectory &&
		it.Path == anotherWrapper.Path
}

func (it SimplePathWrapper) String() string {
	return it.Path
}

func (it *SimplePathWrapper) BaseDir() string {
	return splitinternal.GetBaseDir(it.Path)
}

func (it *SimplePathWrapper) GetFileNamePlusExt() (filename, ext string) {
	return splitinternal.GetFilenamePlusExt(it.Path)
}

func (it *SimplePathWrapper) GetBothExtension() (dotExt, ext string) {
	return splitinternal.GetBothExtension(it.Path)
}

func (it *SimplePathWrapper) GetFileNameOnly() (fileNameWithoutExt string) {
	return splitinternal.GetFileNameWithoutExt(it.Path)
}
