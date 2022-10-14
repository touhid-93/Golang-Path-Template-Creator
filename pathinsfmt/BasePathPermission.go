package pathinsfmt

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
)

type BasePathPermission struct {
	RwxInstruction *chmodins.RwxInstruction `json:"RwxInstruction,omitempty"`
	Chown          *Chown                   `json:"Chown,omitempty"`
}

func (it *BasePathPermission) IsEmptyRwx() bool {
	return it == nil || it.RwxInstruction == nil
}

func (it *BasePathPermission) HasRwx() bool {
	return it != nil || it.RwxInstruction != nil
}

func (it *BasePathPermission) IsEmptyChown() bool {
	return it == nil || it.Chown == nil
}

func (it *BasePathPermission) HasChown() bool {
	return it != nil || it.Chown != nil
}
