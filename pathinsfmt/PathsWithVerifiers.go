package pathinsfmt

type PathsWithVerifiers struct {
	IsContinueOnError bool
	Verifiers         []PathWithVerifier `json:"PathsWithVerifiers,omitempty"`
}

func (it *PathsWithVerifiers) IsPathWithVerifiersUndefined() bool {
	return !it.IsPathWithVerifiersDefined()
}

func (it *PathsWithVerifiers) IsPathWithVerifiersDefined() bool {
	return it != nil && len(it.Verifiers) > 0
}
