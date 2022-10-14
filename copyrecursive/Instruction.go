package copyrecursive

import (
	"gitlab.com/evatix-go/errorwrapper"
)

type Instruction struct {
	SourceDestination
	Options
	copier *Copier
}

func (it *Instruction) lazyCopier() *Copier {
	if it.copier != nil {
		return it.copier
	}

	it.copier = it.Copier()

	return it.copier
}

func (it *Instruction) Copier() *Copier {
	return NewCopier(
		it.Source,
		it.Destination,
		it.Options)
}

func (it *Instruction) Clone() *Instruction {
	if it == nil {
		return nil
	}

	return &Instruction{
		SourceDestination: it.SourceDestination,
		Options:           it.Options,
	}
}

func (it *Instruction) Run() *errorwrapper.Wrapper {
	copier := it.lazyCopier()

	return copier.Copy()
}
