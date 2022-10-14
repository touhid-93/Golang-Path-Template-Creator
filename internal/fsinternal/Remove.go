package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// Remove Reference : https://t.ly/xnAe
func Remove(location string) *errorwrapper.Wrapper {
	err := os.Remove(location)

	return errnew.
		Path.
		Error(errtype.RemoveFailed, err, location)
}
