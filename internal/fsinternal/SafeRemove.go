package fsinternal

import (
	"gitlab.com/evatix-go/errorwrapper"
)

// SafeRemove Reference : https://t.ly/xnAe
func SafeRemove(location string) *errorwrapper.Wrapper {
	if IsPathExists(location) {
		return Remove(location)
	}

	return nil
}
