package fs

import "gitlab.com/evatix-go/errorwrapper"

type (
	CopierOrMoverFunc = func(source, destination string) *errorwrapper.Wrapper
)
