package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

func recursiveOrNonRecursiveLocations(
	isContinueOnError bool,
	isRecursiveCheck bool,
	isNormalize bool,
	locations []string,
) *errstr.Results {
	locationsNormalized := normalize.PathsUsingSingleIfAsync(
		isNormalize,
		locations)

	return recursivepaths.AllOfLocationsIf(
		isRecursiveCheck,
		isContinueOnError,
		locationsNormalized,
	)
}
