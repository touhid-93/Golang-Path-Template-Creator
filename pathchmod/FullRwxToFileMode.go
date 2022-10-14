package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func FullRwxToFileMode(rwxFull string) (os.FileMode, *errorwrapper.Wrapper) {
	rwxWrapper, errWrap := FullRwxToRwxWrapper(rwxFull)

	if errWrap.HasError() {
		return 0, errWrap
	}

	return rwxWrapper.
		ToFileMode(), nil
}
