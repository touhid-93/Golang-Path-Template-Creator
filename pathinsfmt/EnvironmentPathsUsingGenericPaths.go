package pathinsfmt

import (
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/reqtype"
)

type EnvironmentPathsUsingGenericPaths struct {
	coreinstruction.BaseSpecPlusRequestIds
	BaseGenericPathsCollection
	ModifyAs reqtype.Request `json:"ModifyAs"`
}
