package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

// DirsOfLocationsIf recursive all directories recursive locations of
// given locations by expanding using file walk.
func DirsOfLocationsIf(
	isRecursive bool,
	isContinueOnError bool,
	locations []string,
) *errstr.Results {
	return pathsOfLocationsIf(
		isRecursive,
		isContinueOnError,
		Directories,
		locations)
}
