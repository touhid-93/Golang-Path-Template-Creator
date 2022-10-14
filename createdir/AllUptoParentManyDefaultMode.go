package createdir

import "gitlab.com/evatix-go/errorwrapper"

func AllUptoParentManyDefaultMode(locations ...string) *errorwrapper.Wrapper {
	return AllUptoParentMany(DefaultDirectoryFileMode, locations...)
}
