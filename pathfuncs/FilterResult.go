package pathfuncs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

type FilterResult struct {
	FullPath        string
	IsTake, IsBreak bool
	ErrorWrapper    *errorwrapper.Wrapper
}
