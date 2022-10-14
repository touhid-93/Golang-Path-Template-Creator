package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func GetLocationsRwxWrappers(
	isContinueOnError bool,
	locations []string,
) (
	filePathToRwxWrapper map[string]*chmodhelper.RwxWrapper,
	errWrap *errorwrapper.Wrapper,
) {
	resultMap, err := chmodhelper.GetExistingChmodRwxWrappers(
		isContinueOnError, locations...)

	if err != nil {
		return resultMap, errnew.Error.Type(
			errtype.ExistingChmodReadFailed,
			err,
		)
	}

	return resultMap, nil
}
