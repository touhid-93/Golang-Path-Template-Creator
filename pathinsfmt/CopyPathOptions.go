package pathinsfmt

import "gitlab.com/evatix-go/pathhelper/copyrecursive"

type CopyPathOptions struct {
	IsSkipOnExist      bool          `json:"IsSkipOnExist,omitempty"` // removes all before the action
	IsRecursive        bool          `json:"IsRecursive,omitempty"`
	IsCopyRwx          bool          `json:"IsCopyRwx,omitempty"`
	IsMove             bool          `json:"IsMove,omitempty"`
	IsClearDestination bool          `json:"IsClearDestination,omitempty"`
	IsUseShellOrCmd    bool          `json:"IsUseShellOrCmd,omitempty"`
	ApplyPathModifier  *PathModifier `json:"ApplyPathModifier,omitempty"`
}

func (it *CopyPathOptions) CopyRecursiveOptions() *copyrecursive.Options {
	if it == nil {
		return defaultRecursiveCopyOptions
	}

	return &copyrecursive.Options{
		IsSkipOnExist:      it.IsSkipOnExist,
		IsRecursive:        it.IsRecursive,
		IsMove:             it.IsMove,
		IsClearDestination: it.IsClearDestination,
		IsUseShellOrCmd:    it.IsUseShellOrCmd,
		IsNormalize:        false, // no need to normalize, it is done in source
		IsExpandVar:        false,
	}
}
