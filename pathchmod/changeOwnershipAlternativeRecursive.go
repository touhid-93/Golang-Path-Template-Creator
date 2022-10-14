package pathchmod

import (
	"errors"
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func changeOwnershipWindowsRecursive(path, userName, groupName string) *errorwrapper.Wrapper {
	uid, gid, errorWrapper := GetUserGroupId(userName, groupName)
	if errorWrapper.HasError() {
		return errorWrapper
	}

	// https://github.com/gutengo/fil/blob/6109b2e0b5cfdefdef3a254cc1a3eaa35bc89284/file.go#L27-L34
	err := filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}

		if errors.Is(err, filepath.SkipDir) {
			return nil
		}

		if err != nil {
			compiledErrMsg := err.Error() +
				", \nfailed path:" +
				path +
				", \nfailed for chown (name, uid, gid): " +
				name +
				constants.CommaSpace +
				converters.AnyToValueString(uid) +
				constants.CommaSpace +
				converters.AnyToValueString(gid)

			return errors.New(compiledErrMsg)
		}

		return nil
	})

	if err != nil {
		return errnew.Type.Error(errtype.ChmodApplyFailed, err)
	}

	return nil
}
