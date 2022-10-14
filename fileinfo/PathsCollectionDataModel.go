package fileinfo

import "gitlab.com/evatix-go/errorwrapper"

type PathsCollectionDataModel struct {
	RootPath       string
	PathWrappers   *[]*SimplePathWrapper `json:"SimplePathWrappers"`
	Separator      string
	ErrorWrapper   *errorwrapper.Wrapper
	ParentWrappers *Wrappers
}
