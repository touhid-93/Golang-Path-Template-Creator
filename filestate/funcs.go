package filestate

type (
	FilterFunc           func(info *Info) (isTake, isBreak bool)
	TakeAllFilterFunc    func(info *Info) (isTake bool)
	StringerFmtFunc      func(index int, info *Info) string
	HasFilterFunc        func(index int, info *Info) (isSuccess bool)
	HasKeyFilterFunc     func(key string, info *Info) (isSuccess bool)
	MapKeyValFmtFunc     func(info *Info) (key, value string)
	MapKeyValInfoFmtFunc func(info *Info) (key string)
)
