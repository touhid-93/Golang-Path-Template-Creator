package deletepaths

import "gitlab.com/evatix-go/errorwrapper"

func AllOnExistIf(
	isRemoveOnExist bool,
	locations ...string,
) *errorwrapper.Wrapper {
	if isRemoveOnExist {
		return AllOnExist(locations...)
	}

	return All(locations...)
}
