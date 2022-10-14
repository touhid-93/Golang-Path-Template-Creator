package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ExistingChmodRwxWrapper(
	location string,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	existingChmod, err := chmodhelper.GetExistingChmodRwxWrapperPtr(location)

	if err != nil {
		return existingChmod, errnew.
			Path.
			Error(
				errtype.ExistingChmodReadFailed,
				err,
				location)
	}

	return existingChmod, nil
}
