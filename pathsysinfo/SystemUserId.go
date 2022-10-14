package pathsysinfo

import (
	"os/user"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SystemUserId(userResult *user.User) (int, *errorwrapper.Wrapper) {
	uid, errUidConvert := strconv.Atoi(userResult.Uid)

	if errUidConvert != nil {
		return constants.InvalidValue, errnew.Type.Error(errtype.ConversionValueToInteger, errUidConvert)
	}

	return uid, nil
}
