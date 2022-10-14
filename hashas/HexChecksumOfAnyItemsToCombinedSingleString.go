package hashas

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func HexChecksumOfAnyItemsToCombinedSingleString(
	isSkipOnNil bool,
	method Variant,
	items ...interface{},
) *errstr.Result {
	results := HexChecksumOfAnyItems(
		isSkipOnNil,
		method,
		items...)

	if results.HasError() {
		return errstr.New.Result.ErrorWrapper(results.ErrorWrapper)
	}

	return method.HexSumOfAny(results.Values)
}
