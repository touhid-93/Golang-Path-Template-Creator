package pathfuncs

import "sync"

type (
	Filter                   func(arg *FilterArg) *FilterResult
	SimpleFilterUsingArg     func(arg *FilterArg) (isTake bool)
	SimpleFilter             func(fullPath string) (isTake bool, err error)
	FilterPathToFilterResult func(fullPath string) *FilterResult

	ProcessedResult struct {
		Result          string
		IsKeep, IsBreak bool
	}

	ProcessorIn struct {
		Index       int
		CurrentPath string
		Wg          *sync.WaitGroup
	}

	Processor func(index int, currentPath string) (result string)
)
