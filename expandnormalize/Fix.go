package expandnormalize

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Fix(
	location string,
) string {
	location = expandpath.ExpandVariables(
		location)

	return normalize.Path(
		location)
}
