package recursivepaths

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

// AllOfLocationsIf recursive all (files + dirs) recursive locations of
// given locations by expanding using file walk.
func AllOfLocationsIf(
	isRecursive bool,
	isContinueOnError bool,
	locations []string,
) *errstr.Results {
	return pathsOfLocationsIf(
		isRecursive,
		isContinueOnError,
		All,
		locations)
}
