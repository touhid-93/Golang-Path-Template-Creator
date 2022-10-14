package pathfuncs

import "gitlab.com/evatix-go/errorwrapper"

func InvalidFilterResultUsingErrWp(
	fullPath string,
	errWrap *errorwrapper.Wrapper,
) *FilterResult {
	return &FilterResult{
		FullPath:     fullPath,
		ErrorWrapper: errWrap,
	}
}
