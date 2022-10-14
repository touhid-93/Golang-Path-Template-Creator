package pathinsfmt

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
)

type PathModifier struct {
	chmodins.BaseRwxInstructions
	ChmodCommands *ChmodCommands `json:"ChmodCommands,omitempty"`
	Chown         *Chown         `json:"Chown,omitempty"`
	ChangeGroup   *ChangeGroup   `json:"ChangeGroup,omitempty"`
}

func (it *PathModifier) HasChmodCommands() bool {
	return it != nil &&
		it.ChmodCommands != nil &&
		it.ChmodCommands.HasAnyItem()
}

func (it *PathModifier) HasChangeGroup() bool {
	return it != nil &&
		it.ChangeGroup != nil
}

func (it *PathModifier) HasChown() bool {
	return it != nil &&
		it.Chown != nil
}

func (it *PathModifier) HasRwxInstructions() bool {
	return it != nil &&
		it.RwxInstructions != nil
}

func (it *PathModifier) Clone() *PathModifier {
	if it == nil {
		return nil
	}

	return &PathModifier{
		BaseRwxInstructions: *it.BaseRwxInstructions.Clone(),
		ChmodCommands:       it.ChmodCommands.Clone(),
		Chown:               it.Chown.Clone(),
		ChangeGroup:         it.ChangeGroup.Clone(),
	}
}
