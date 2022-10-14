package pathstatlinux

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

type Info struct {
	User
	Group
	RawLocationTimestamp
	RwxSimple
	Location       string
	IsDirectory    bool
	IsPathExist    bool
	IsValidParsing bool
	ErrorWrapper   *errorwrapper.Wrapper
	FileInfo       os.FileInfo // could be nil
}

func InvalidInfo(location string) *Info {
	return &Info{
		User:                 *InvalidUser(),
		Group:                *InvalidGroup(),
		RawLocationTimestamp: RawLocationTimestamp{},
		RwxSimple:            *InvalidRwxSimple(),
		Location:             location,
		IsDirectory:          false,
		IsPathExist:          false,
		IsValidParsing:       false,
		ErrorWrapper:         nil,
		FileInfo:             nil,
	}
}

func InvalidInfoUsingErr(location string, errWrapper *errorwrapper.Wrapper) *Info {
	return &Info{
		User:                 *InvalidUser(),
		Group:                *InvalidGroup(),
		RawLocationTimestamp: RawLocationTimestamp{},
		RwxSimple:            *InvalidRwxSimple(),
		Location:             location,
		IsDirectory:          false,
		IsPathExist:          false,
		IsValidParsing:       false,
		ErrorWrapper:         errWrapper,
		FileInfo:             nil,
	}
}

func (receiver *Info) IsSafe() bool {
	return receiver.IsValidParsing && !receiver.HasError()
}

func (receiver *Info) IsEmptyError() bool {
	return receiver.ErrorWrapper.IsEmptyError()
}

func (receiver *Info) HasError() bool {
	return receiver.ErrorWrapper.HasError()
}

func (receiver *Info) HandleError() {
	receiver.ErrorWrapper.HandleError()
}

func (receiver *Info) HandleErrorWithMessage(message string) {
	receiver.ErrorWrapper.HandleErrorWithMsg(message)
}
