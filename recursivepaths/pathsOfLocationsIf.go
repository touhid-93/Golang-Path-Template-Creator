package recursivepaths

import (
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func pathsOfLocationsIf(
	isRecursive bool,
	isContinueOnError bool,
	pathsExpander func(location string) *errstr.Results,
	locations []string,
) *errstr.Results {
	if !isRecursive {
		return errstr.New.Results.Strings(
			locations)
	}

	isExitImmediate := !isContinueOnError
	capacity := defaultcapacity.PredictiveDefault(
		len(locations))
	slice := stringslice.MakeDefault(capacity)
	var sliceErr []string

	for _, location := range locations {
		results := pathsExpander(location)

		if results.HasError() {
			sliceErr = append(sliceErr,
				results.ErrorWrapper.FullString())
		}

		slice = append(
			slice,
			results.SafeValues()...)

		if isExitImmediate && results.HasError() {
			err := errcore.SliceToError(sliceErr)
			return &errstr.Results{
				Values: slice,
				ErrorWrapper: errnew.
					Path.
					Error(
						errtype.FileExpand,
						err,
						location),
			}
		}
	}

	err := errcore.SliceToError(sliceErr)
	if err != nil {
		return errstr.New.Results.ErrorWrapper(
			errnew.
				Path.
				Error(
					errtype.FileExpand,
					err,
					""))
	}

	return errstr.New.Results.Strings(
		slice)
}
