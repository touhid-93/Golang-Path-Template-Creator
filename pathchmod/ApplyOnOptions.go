package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func ApplyOnOptions(
	isApply,
	isSkipOnInvalid,
	isApplyOnMismatch bool,
	changeFileMode os.FileMode,
	location string,
) *errorwrapper.Wrapper {
	if !isApply {
		return nil
	}

	if isApplyOnMismatch {
		return ApplyOnMismatch(
			isSkipOnInvalid,
			changeFileMode,
			location)
	}

	return ApplyChmodDefault(
		changeFileMode,
		location)
}
