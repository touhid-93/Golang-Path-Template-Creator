package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ExistingRwxWrapperWithError(location string) *RwxWrapperWithError {
	if location == "" {
		return &RwxWrapperWithError{
			RwxWrapper: nil,
			ErrWrapper: errnew.Path.Empty(),
		}
	}

	rwxWrapper, err := chmodhelper.GetExistingChmodRwxWrapperPtr(location)
	if err != nil {
		pathErr := errnew.
			Path.
			Messages(
				errtype.File,
				location,
				"ExistingRwxWrapperWithError",
				err.Error())

		return &RwxWrapperWithError{
			RwxWrapper: nil,
			ErrWrapper: pathErr,
		}
	}

	return &RwxWrapperWithError{
		RwxWrapper: rwxWrapper,
		ErrWrapper: nil,
	}
}
