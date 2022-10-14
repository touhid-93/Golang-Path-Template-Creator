package pathinsfmt

type ChangeGroup struct {
	BaseIsRecursive
	BaseGroupName
}

func (it *ChangeGroup) Clone() *ChangeGroup {
	if it == nil {
		return nil
	}

	return &ChangeGroup{
		BaseIsRecursive: BaseIsRecursive{
			IsRecursive: it.IsRecursive,
		},
		BaseGroupName: *it.BaseGroupName.Clone(),
	}
}
