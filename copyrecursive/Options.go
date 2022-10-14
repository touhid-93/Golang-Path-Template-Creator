package copyrecursive

import "gitlab.com/evatix-go/pathhelper/expandnormalize"

type Options struct {
	IsSkipOnExist      bool
	IsRecursive        bool
	IsMove             bool
	IsClearDestination bool
	IsUseShellOrCmd    bool
	IsNormalize        bool
	IsExpandVar        bool
}

func NewDefaultOptions() Options {
	return Options{
		IsSkipOnExist:      false,
		IsRecursive:        true,
		IsMove:             false,
		IsClearDestination: false,
		IsUseShellOrCmd:    false,
		IsNormalize:        false,
		IsExpandVar:        false,
	}
}

func (it Options) FixedPath(location string) string {
	return expandnormalize.FixIf(
		it.IsNormalize,
		it.IsExpandVar,
		location)
}
