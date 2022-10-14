package unipaths

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/unipath"
)

func New(
	sep string,
) *Wrappers {
	list := make(
		[]*unipath.Wrapper,
		constants.Zero,
		defaultCapacity)

	return &Wrappers{
		separator: sep,
		items:     list,
	}
}

func NewUsingPath(
	path, sep string,
) *Wrappers {
	list := make(
		[]*unipath.Wrapper,
		constants.Zero,
		constants.ArbitraryCapacity1)

	wrappers := &Wrappers{
		separator: sep,
		items:     list,
	}

	return wrappers.
		AddPathAs(path)
}

func NewUsingCap(
	cap int, sep string,
) *Wrappers {
	list := make(
		[]*unipath.Wrapper,
		constants.Zero,
		cap)

	return &Wrappers{
		separator: sep,
		items:     list,
	}
}

func NewMap(
	sep string,
) *WrappersMap {
	list := make(
		map[string]*unipath.Wrapper,
		defaultCapacity)

	return &WrappersMap{
		separator: sep,
		items:     list,
	}
}

func NewMapUsingCap(
	cap int, sep string,
) *WrappersMap {
	list := make(
		map[string]*unipath.Wrapper,
		cap)

	return &WrappersMap{
		separator: sep,
		items:     list,
	}
}
