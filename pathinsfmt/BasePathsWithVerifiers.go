package pathinsfmt

type BasePathsWithVerifiers struct {
	PathsWithVerifiers *PathsWithVerifiers `json:"PathsWithVerifiers,omitempty"`
}

func (it *BasePathsWithVerifiers) IsPathWithVerifiersDefined() bool {
	return it != nil && it.PathsWithVerifiers.IsPathWithVerifiersDefined()
}
