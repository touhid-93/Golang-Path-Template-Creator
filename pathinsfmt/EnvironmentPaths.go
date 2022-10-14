package pathinsfmt

import (
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/reqtype"
)

type EnvironmentPaths struct {
	coreinstruction.BaseSpecPlusRequestIds
	ModifyAs reqtype.Request `json:"ModifyAs"`
	Paths    []string        `json:"Locations,omitempty"`
}

func (e EnvironmentPaths) Length() int {
	return len(e.Paths)
}

func (e EnvironmentPaths) IsEmpty() bool {
	return len(e.Paths) == 0
}

func (e EnvironmentPaths) HasAnyItem() bool {
	return len(e.Paths) > 0
}
