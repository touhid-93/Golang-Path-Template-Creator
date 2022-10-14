package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinstruction"
)

type PathModifiersApply struct {
	coreinstruction.BaseSpecPlusRequestIds
	BaseGenericPathsCollection
	PathModifiers []PathModifier `json:"PathModifiers,omitempty"`
}

func (p *PathModifiersApply) PathModifiersLength() int {
	if p == nil {
		return constants.Zero
	}

	return len(p.PathModifiers)
}

func (p *PathModifiersApply) IsEmptyPathModifiers() bool {
	if p == nil {
		return true
	}

	return p.PathModifiersLength() == 0
}
