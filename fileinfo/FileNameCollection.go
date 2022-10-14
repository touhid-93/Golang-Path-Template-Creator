package fileinfo

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
)

type FileNamesCollection struct {
	RootPath string
	// TODO fix this to file paths
	names          []string
	Error          *errorwrapper.Wrapper
	parentWrappers *Wrappers
}

func NewFileNames(rootPath string, capacity int) *FileNamesCollection {
	paths := make([]string, 0, capacity)

	return &FileNamesCollection{
		RootPath:       rootPath,
		names:          paths,
		Error:          nil,
		parentWrappers: nil,
	}
}

func NewFileNamesUsingWrappers(wrappers *Wrappers) *FileNamesCollection {
	if wrappers == nil {
		return &FileNamesCollection{
			RootPath:       "",
			names:          nil,
			Error:          nil,
			parentWrappers: wrappers,
		}
	}

	if wrappers.IsEmpty() {
		return &FileNamesCollection{
			RootPath:       wrappers.RootPath,
			names:          nil,
			Error:          wrappers.ErrorWrapper,
			parentWrappers: wrappers,
		}
	}

	names := make([]string, wrappers.Length())
	for i, wrapper := range wrappers.Items {
		if wrapper.FileInfo == nil {
			continue
		}

		names[i] = wrapper.FileInfo.Name()
	}

	return &FileNamesCollection{
		RootPath:       wrappers.RootPath,
		names:          names,
		Error:          wrappers.ErrorWrapper,
		parentWrappers: wrappers,
	}
}

func NewFileNamesUsing(
	directoryPath, separator string,
	isNormalize bool,
) *FileNamesCollection {
	wrappers := NewWrappersPtr(
		directoryPath,
		separator,
		isNormalize)

	return NewFileNamesUsingWrappers(wrappers)
}

func (it *FileNamesCollection) IsEmpty() bool {
	return it == nil ||
		it.names == nil ||
		len(it.names) == 0
}

func (it *FileNamesCollection) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.Error.HasError()
}

func (it *FileNamesCollection) IsContains(
	fileName string,
	isCaseSensitive bool,
) bool {
	return stringutil.IsContainsPtrSimple(
		&it.names,
		fileName,
		0,
		isCaseSensitive)
}

func (it *FileNamesCollection) Length() int {
	if it == nil || it.names == nil {
		return 0
	}

	return len(it.names)
}

// GetFilePaths
//
// Root level files paths, no nested paths.
func (it *FileNamesCollection) GetFilePaths(
	separator string,
) []string {
	filePaths := make([]string, 0, it.Length())

	for _, name := range it.names {
		newPath := it.RootPath + separator + name
		filePaths = append(filePaths, newPath)
	}

	return filePaths
}

func (it *FileNamesCollection) IsParentWrappersEmpty() bool {
	return it.parentWrappers == nil ||
		it.parentWrappers.IsEmpty()
}

func (it *FileNamesCollection) HasParentWrappers() bool {
	return it.parentWrappers != nil
}

func (it *FileNamesCollection) ParentWrappers() *Wrappers {
	return it.parentWrappers
}

func (it *FileNamesCollection) Add(fileName string) *FileNamesCollection {
	it.names = append(it.names, fileName)

	return it
}

func (it *FileNamesCollection) AddWrapper(wrapper Wrapper) *FileNamesCollection {
	it.names = append(it.names, wrapper.FileInfo.Name())

	return it
}

func (it *FileNamesCollection) OnlyNamesCollection() []string {
	return it.names
}
