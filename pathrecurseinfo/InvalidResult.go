package pathrecurseinfo

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
)

func InvalidResult(
	root string,
	errorWrapper *errorwrapper.Wrapper,
	stat *chmodhelper.PathExistStat,
) *Result {
	return &Result{
		Root:            root,
		IsInvalidResult: true,
		PathStat:        stat,
		ErrorWrapper:    errorWrapper,
		PathsResult:     EmptyPathsResult(),
	}
}
