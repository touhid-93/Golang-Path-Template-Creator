package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

type ChmodWithError struct {
	Chmod os.FileMode
	errorwrapper.ErrWrapper
}
