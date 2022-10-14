package elitepath

type ExistFilter struct {
	IsExist bool
	IsDir   bool
}

func (it *ExistFilter) IsMatch(path *Path) bool {
	if it == nil {
		return true
	}

	if path == nil {
		return false
	}

	stat := path.ExistStat()

	return stat.IsExist == it.IsExist &&
		stat.IsDir() == it.IsDir
}

func (it *ExistFilter) IsMatchLazy(lazyPath *LazyPath) bool {
	if it == nil {
		return true
	}

	if lazyPath == nil {
		return false
	}

	stat := lazyPath.PathExistStat()

	return stat.IsExist == it.IsExist &&
		stat.IsDir() == it.IsDir
}
