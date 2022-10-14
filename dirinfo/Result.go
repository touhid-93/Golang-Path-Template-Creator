package dirinfo

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
	"gitlab.com/evatix-go/pathhelper/performingas"
)

type Result struct {
	IsValidDir        bool
	HasIssues         bool
	IsIgnoredAction   bool
	Action            performingas.Action
	RawPath           string
	FileModeRequested *os.FileMode
	FileInfoWrapper   *fileinfo.Wrapper
	Error             *errorwrapper.Wrapper
}

func Empty() *Result {
	return EmptyUsingInfo(nil)
}

func EmptyUsingInfo(fileWrapperInfo *fileinfo.Wrapper) *Result {
	return &Result{
		FileInfoWrapper:   fileWrapperInfo,
		Error:             nil,
		IsValidDir:        false,
		RawPath:           "",
		FileModeRequested: nil,
		HasIssues:         false,
		IsIgnoredAction:   true,
		Action:            performingas.EmptyDirectoryResult,
	}
}

func New(fileOrDirPath string) *Result {
	isFilePathEmpty := fileOrDirPath == ""
	fileInfo, err := os.Stat(fileOrDirPath)
	errWrapper := errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			err,
			fileOrDirPath)
	isErrorEmpty := errWrapper.IsEmpty()

	fileInfoWrapper := &fileinfo.Wrapper{
		FileInfo:     fileInfo,
		ErrorWrapper: errWrapper,
		RawPath:      fileOrDirPath,
		IsDirectory:  isErrorEmpty && fileInfo.IsDir(),
		IsFile:       isErrorEmpty && !fileInfo.IsDir(),
		IsEmptyPath:  isFilePathEmpty,
	}

	if !fileInfoWrapper.IsDirectory {
		return &Result{
			FileInfoWrapper:   fileInfoWrapper,
			Error:             errWrapper,
			RawPath:           fileOrDirPath,
			FileModeRequested: nil,
			IsValidDir:        false,
			HasIssues:         !isErrorEmpty,
			IsIgnoredAction:   true,
			Action:            performingas.EmptyDirectoryResult,
		}
	}

	fileMode := fileInfo.Mode()

	return &Result{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             errWrapper,
		RawPath:           fileOrDirPath,
		IsValidDir:        true,
		FileModeRequested: &fileMode,
		HasIssues:         !isErrorEmpty,
		IsIgnoredAction:   true,
		Action:            performingas.NoAction,
	}
}

func (receiver *Result) IsEmpty() bool {
	return receiver == nil ||
		receiver.Action.IsEmptyDirectoryResult()
}

func (receiver *Result) HasValidDir() bool {
	return receiver != nil &&
		receiver.IsValidDir &&
		!receiver.HasIssues
}
