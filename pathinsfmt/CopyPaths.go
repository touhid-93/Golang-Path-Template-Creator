package pathinsfmt

type CopyPaths struct {
	IsContinueOnError bool
	CopyPaths         []CopyPath `json:"CopyPaths,omitempty"`
}

func (it *CopyPaths) Length() int {
	if it == nil {
		return 0
	}

	return len(it.CopyPaths)
}

func (it *CopyPaths) IsEmpty() bool {
	return it.Length() == 0
}
