package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ParseRwxOwnerGroupOtherToFileMode(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (os.FileMode, *errorwrapper.Wrapper) {
	rwxWrapper, errWrap := ParseRwxOwnerGroupOtherToRwxWrapper(rwxOwnerGroupOther)

	if errWrap.HasError() {
		return 0, errWrap
	}

	if rwxWrapper == nil {
		return 0, errnew.Messages.Many(
			errtype.ChmodInvalid,
			"Cannot process wildcard rwx to convert to fixed file mode")
	}

	return rwxWrapper.ToFileMode(), nil
}
