package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinstruction"
)

type PathVerifiers struct {
	coreinstruction.BaseSpecPlusRequestIds
	PathVerifiers           []PathVerifier `json:"PathVerifiers,omitempty"`
	IsSkipCheckingOnInvalid bool
	IsNormalize             bool
	IsRecursiveCheck        bool
}

func (v *PathVerifiers) Length() int {
	if v == nil {
		return constants.Zero
	}

	return len(v.PathVerifiers)
}

func (v *PathVerifiers) IsEmpty() bool {
	return v.Length() == 0
}

func (v *PathVerifiers) HasAnyItem() bool {
	return v.Length() > 0
}
