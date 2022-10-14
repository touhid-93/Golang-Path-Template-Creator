package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

// FilesOfLocationsIf recursive all files recursive locations of
// given locations by expanding using file walk.
func FilesOfLocationsIf(
	isRecursive bool,
	isContinueOnError bool,
	locations []string,
) *errstr.Results {
	return pathsOfLocationsIf(
		isRecursive,
		isContinueOnError,
		Files,
		locations)
}
