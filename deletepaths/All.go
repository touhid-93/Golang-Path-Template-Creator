package deletepaths

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func All(locations ...string) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return nil
	}

	for _, location := range locations {
		recursiveErr := Single(location)

		if recursiveErr.HasError() {
			return recursiveErr
		}
	}

	return nil
}

func AllOnExist(locations ...string) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return nil
	}

	for _, location := range locations {
		recursiveErr := SingleOnExist(location)

		if recursiveErr.HasError() {
			return recursiveErr
		}
	}

	return nil
}
