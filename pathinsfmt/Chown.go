package pathinsfmt

type Chown struct {
	BaseIsRecursive
	UserGroupName
}

func (it *Chown) Clone() *Chown {
	if it == nil {
		return nil
	}

	return &Chown{
		BaseIsRecursive: BaseIsRecursive{
			IsRecursive: it.IsRecursive,
		},
		UserGroupName: *it.UserGroupName.Clone(),
	}
}
