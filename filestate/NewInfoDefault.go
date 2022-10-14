package filestate

import "gitlab.com/evatix-go/errorwrapper"

func NewInfoDefault(
	filePath string,
) (*Info, *errorwrapper.Wrapper) {
	return NewInfo(
		DefaultHashMethod,
		true,
		filePath)
}
