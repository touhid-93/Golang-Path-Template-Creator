package createdir

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func AllDefaultMode(location string) *errorwrapper.Wrapper {
	return All(location, DefaultDirectoryFileMode)
}
