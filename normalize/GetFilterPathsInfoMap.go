package normalize

import "gitlab.com/evatix-go/core/chmodhelper"

func GetFilterPathsInfoMap(
	isNormalize,
	isSkipOnInvalid bool,
	locations []string,
) *chmodhelper.FilteredPathFileInfoMap {
	locationsNormalized :=
		PathsUsingSingleIfAsync(
			isNormalize,
			locations)

	return chmodhelper.GetExistsFilteredPathFileInfoMap(
		isSkipOnInvalid,
		locationsNormalized...)
}
