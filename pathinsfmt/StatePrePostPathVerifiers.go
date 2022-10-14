package pathinsfmt

type StatePrePostPathVerifiers struct {
	PreStates  *PathsWithVerifiers `json:"PrePathModifiers,omitempty"`
	PostStates *PathsWithVerifiers `json:"PostPathModifiers,omitempty"`
}

func (it *StatePrePostPathVerifiers) IsPreStatesDefined() bool {
	return it != nil && it.PreStates.IsPathWithVerifiersDefined()
}

func (it *StatePrePostPathVerifiers) IsPostStatesDefined() bool {
	return it != nil && it.PostStates.IsPathWithVerifiersDefined()
}
