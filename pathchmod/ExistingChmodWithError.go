package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ExistingChmodWithError(location string) *ChmodWithError {
	if location == "" {
		return &ChmodWithError{
			ErrWrapper: errnew.Path.Empty(),
		}
	}

	chmod, err := chmodhelper.GetExistingChmod(location)
	if err != nil {
		pathErr := errnew.
			Path.
			Messages(
				errtype.File,
				location,
				"ExistingChmodWithError",
				err.Error())

		return &ChmodWithError{
			Chmod:      0,
			ErrWrapper: pathErr,
		}
	}

	return &ChmodWithError{
		Chmod:      chmod,
		ErrWrapper: nil,
	}
}
