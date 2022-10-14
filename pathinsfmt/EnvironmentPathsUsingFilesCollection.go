package pathinsfmt

import "gitlab.com/evatix-go/core/reqtype"

type EnvironmentPathsUsingFilesCollection struct {
	BaseLocationCollection
	ModifyAs reqtype.Request `json:"ModifyAs"`
}
