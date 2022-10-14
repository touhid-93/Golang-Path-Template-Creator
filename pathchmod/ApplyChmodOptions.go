package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func ApplyChmodOptions(
	isApply,
	isSkipInvalid bool,
	changeFileMode os.FileMode,
	location string,
) *errorwrapper.Wrapper {
	if !isApply {
		return nil
	}

	if !isSkipInvalid {
		return ApplyChmodDefault(
			changeFileMode, location)
	}

	return ApplyChmodSkipInvalidFile(
		changeFileMode, location)
}
