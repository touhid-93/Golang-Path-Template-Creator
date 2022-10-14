package elitepath

import "gitlab.com/evatix-go/pathhelper/pathfixer"

type Directory struct {
	WithPermission
	IsClearDir                bool `json:"IsClearDir,omitempty"`
	IsApplyDefaultChmodGroups bool `json:"IsApplyDefaultChmodGroups,omitempty"`
}

func NewDirectory(isClear, isDefaultChmod bool, location string) *Directory {
	return &Directory{
		WithPermission: WithPermission{
			Path: Path{
				Location: pathfixer.Location{
					Path: location,
				},
			},
		},
		IsClearDir:                isClear,
		IsApplyDefaultChmodGroups: isDefaultChmod,
	}
}
