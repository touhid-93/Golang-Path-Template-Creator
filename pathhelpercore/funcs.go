package pathhelpercore

import "gitlab.com/evatix-go/errorwrapper"

type (
	InvokerFunc func(fileInfo *FileInfo) *errorwrapper.Wrapper
)
