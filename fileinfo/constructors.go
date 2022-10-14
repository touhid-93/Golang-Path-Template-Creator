package fileinfo

import (
	"io/ioutil"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/ispath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func New(rawPath, separator string) *Wrapper {
	isEmptyPath := ispath.Empty(rawPath)

	if isEmptyPath {
		return &Wrapper{
			FileInfo:     nil,
			ErrorWrapper: errnew.Path.Empty(),
			RawPath:      rawPath,
			IsDirectory:  false,
			IsFile:       false,
			IsEmptyPath:  isEmptyPath,
		}
	}

	fileInfo, err := os.Stat(rawPath)
	isDir := err == nil && fileInfo.IsDir()
	var errWrapper *errorwrapper.Wrapper

	if err != nil {
		errWrapper = errnew.
			Path.
			Error(
				errtype.PathRelatedIssue,
				err,
				rawPath)
	}

	return &Wrapper{
		FileInfo:     fileInfo,
		ErrorWrapper: errWrapper,
		RawPath:      rawPath,
		IsDirectory:  isDir,
		IsFile:       err == nil && !isDir,
		IsEmptyPath:  isEmptyPath,
		Separator:    separator,
	}
}

func NewError(
	filePath, separator string,
	err error,
) *Wrapper {
	isFilePathEmpty := ispath.Empty(filePath)

	if err != nil {
		return &Wrapper{
			FileInfo:     nil,
			ErrorWrapper: errnew.Path.Empty(),
			RawPath:      filePath,
			IsDirectory:  false,
			IsFile:       false,
			IsEmptyPath:  isFilePathEmpty,
			Separator:    separator,
		}
	}

	fileErrWrapper := errnew.
		Path.
		Error(
			errtype.PathMissingOrInvalid,
			err,
			filePath)

	return &Wrapper{
		ErrorWrapper: fileErrWrapper,
		RawPath:      filePath,
		IsDirectory:  false,
		IsFile:       false,
		IsEmptyPath:  isFilePathEmpty,
		Separator:    separator,
	}
}

func NewUsingInfo(
	osFileInfo os.FileInfo,
	filePath, separator string,
	err error,
) *Wrapper {
	isFilePathEmpty := ispath.Empty(filePath)

	if err != nil || osFileInfo == nil {
		return NewError(
			filePath,
			separator,
			err)
	}

	fileErr := errnew.
		Path.
		Error(
			errtype.PathMissingOrInvalid,
			err,
			filePath)

	isDir := osFileInfo.IsDir()

	return &Wrapper{
		FileInfo:     osFileInfo,
		ErrorWrapper: fileErr,
		RawPath:      filePath,
		IsDirectory:  isDir,
		IsFile:       !isDir,
		IsEmptyPath:  isFilePathEmpty,
		Separator:    separator,
	}
}

func NewWrappersPtrUsingCapacity(rootPath string, capacity int) *Wrappers {
	collection := make([]*Wrapper, 0, capacity)

	return &Wrappers{
		RootPath: rootPath,
		Items:    collection,
	}
}

func EmptyWrappers() *Wrappers {
	return NewWrappersPtrUsingCapacity("", 0)
}

func NewWrappersPtr(
	filePath, separator string,
	isNormalize bool,
) *Wrappers {
	fileInfos, err := ioutil.ReadDir(filePath)

	if err != nil {
		return &Wrappers{
			RootPath:     filePath,
			ErrorWrapper: errnew.Path.Empty(),
			Separator:    separator,
		}
	}

	collection := make(
		[]*Wrapper,
		len(fileInfos))

	filePathNormalized := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		filePath,
	)

	for i, fileInfo := range fileInfos {
		newFilePath := filePathNormalized +
			separator +
			fileInfo.Name()

		wrapper := NewUsingInfo(
			fileInfo,
			newFilePath,
			separator,
			nil)

		collection[i] = wrapper
	}

	return &Wrappers{
		RootPath:  filePath,
		Items:     collection,
		Separator: separator,
	}
}
