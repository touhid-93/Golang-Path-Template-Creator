package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
)

type RwxWrapperWithError struct {
	RwxWrapper *chmodhelper.RwxWrapper
	errorwrapper.ErrWrapper
}
