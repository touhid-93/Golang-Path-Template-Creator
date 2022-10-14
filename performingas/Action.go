package performingas

type Action byte

//goland:noinspection ALL
const (
	CreateAction Action = iota
	ReadAction
	DeleteAction
	EmptyDirectoryResult
	NoAction
)

func (a Action) Is(action Action) bool {
	return a == action
}

func (a Action) IsCreateAction() bool {
	return a == CreateAction
}

func (a Action) IsDeleteAction() bool {
	return a == DeleteAction
}

func (a Action) IsReadAction() bool {
	return a == ReadAction
}

func (a Action) IsNoAction() bool {
	return a == NoAction
}

func (a Action) IsEmptyDirectoryResult() bool {
	return a == EmptyDirectoryResult
}
