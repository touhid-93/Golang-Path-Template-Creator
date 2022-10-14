package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinstruction"
)

type SymbolicLinks struct {
	coreinstruction.BaseSpecPlusRequestIds
	IsContinueOnError bool           `json:"IsContinueOnError,omitempty"`
	SymbolicLinks     []SymbolicLink `json:"SymbolicLinks,omitempty"`
}

func (receiver *SymbolicLinks) Length() int {
	if receiver == nil {
		return constants.Zero
	}

	return len(receiver.SymbolicLinks)
}

func (receiver *SymbolicLinks) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *SymbolicLinks) HasAnyItem() bool {
	return receiver.Length() > 0
}
