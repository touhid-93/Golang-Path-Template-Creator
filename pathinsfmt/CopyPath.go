package pathinsfmt

import (
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
)

type CopyPath struct {
	SourceDestinationPlusCompiled
	Options *CopyPathOptions `json:"Options,omitempty"`
	copier  *copyrecursive.Copier
}

func (it *CopyPath) HasOptions() bool {
	return it != nil && it.Options != nil
}

func (it *CopyPath) HasApplyPathModifier() bool {
	return it != nil && it.Options != nil && it.Options.ApplyPathModifier != nil
}

func (it *CopyPath) Copier() *copyrecursive.Copier {
	compiled := it.CompiledSourceDestination()
	options := it.Options.CopyRecursiveOptions()

	return copyrecursive.NewCopier(
		compiled.Source,
		compiled.Destination,
		*options)
}

func (it *CopyPath) LazyCopier() *copyrecursive.Copier {
	if it.copier != nil {
		return it.copier
	}

	it.copier = it.Copier()

	return it.copier
}
