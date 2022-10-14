package pathinsfmt

import (
	"strings"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathfixer"
)

type SourceDestinationPlusCompiled struct {
	BaseSourceDestination
	pathfixer.PathOptions
	sourceDestinationCompiled *BaseSourceDestination
}

func NewSourceDestinationPlusCompiledNoOptions(
	src, dest string,
) *SourceDestinationPlusCompiled {
	return NewSourceDestinationPlusCompiled(
		src,
		dest,
		nil)
}

func NewSourceDestinationPlusCompiled(
	src, dest string,
	options *pathfixer.PathOptions,
) *SourceDestinationPlusCompiled {
	if options == nil {
		return &SourceDestinationPlusCompiled{
			BaseSourceDestination: BaseSourceDestination{
				Source:      src,
				Destination: dest,
			},
			PathOptions: pathfixer.PathOptions{},
		}
	}

	return &SourceDestinationPlusCompiled{
		BaseSourceDestination: BaseSourceDestination{
			Source:      src,
			Destination: dest,
		},
		PathOptions: *options,
	}
}

func (it SourceDestinationPlusCompiled) CompiledSourceDestination() *BaseSourceDestination {
	if it.sourceDestinationCompiled != nil {
		return it.sourceDestinationCompiled
	}

	source := it.GetFixedPath(it.Source)
	destination := it.GetFixedPath(it.Destination)

	it.sourceDestinationCompiled = &BaseSourceDestination{
		Source:      source,
		Destination: destination,
	}

	return it.sourceDestinationCompiled
}

func (it SourceDestinationPlusCompiled) FixedSrcDest() (src, dest string) {
	compiled := it.
		CompiledSourceDestination()

	return compiled.Source, compiled.Destination
}

func (it SourceDestinationPlusCompiled) DestinationFixedPath() string {
	return it.
		CompiledSourceDestination().
		Destination
}

func (it SourceDestinationPlusCompiled) SourceFixedPath() string {
	return it.
		CompiledSourceDestination().
		Source
}

func (it SourceDestinationPlusCompiled) IsSameSourceDestination() bool {
	if it.Destination == it.Source {
		return true
	}

	compiled := it.CompiledSourceDestination()
	if compiled.Destination == compiled.Source {
		return true
	}

	if osconsts.IsWindows {
		return strings.EqualFold(
			compiled.Destination,
			compiled.Source)
	}

	return false
}
