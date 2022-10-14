package deletepaths

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func AllRecursive(locations []string) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return nil
	}

	for _, location := range locations {
		recursiveErr := Recursive(location)

		if recursiveErr.HasError() {
			return recursiveErr
		}
	}

	return nil
}
