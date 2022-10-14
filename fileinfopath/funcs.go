package fileinfopath

type (
	FilterFunc                       func(instance *Instance) (isTake, isBreak bool)
	TakeAllFilterFunc                func(instance *Instance) (isTake bool)
	StringerFmtFunc                  func(index int, instance *Instance) string
	HasFilterFunc                    func(index int, instance *Instance) (isSuccess bool)
	HasKeyFilterFunc                 func(key string, instance *Instance) (isSuccess bool)
	MapKeyValueStringFmtFunc         func(instance *Instance) (key, value string)
	MapKeyStringValueInstanceFmtFunc func(instance *Instance) (key string)
)
