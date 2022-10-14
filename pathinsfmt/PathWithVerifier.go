package pathinsfmt

type PathWithVerifier struct {
	PathWithOptions
	Verifier *PathVerifier `json:"Verifier,omitempty"`
}

func (it *PathWithVerifier) IsVerifierUndefined() bool {
	return it == nil || it.Verifier == nil
}

func (it *PathWithVerifier) IsVerifierDefined() bool {
	return it != nil && it.Verifier != nil
}

func (it *PathWithVerifier) IsPathWithVerifierDefined() bool {
	return it != nil && it.Path != "" && it.Verifier != nil
}
