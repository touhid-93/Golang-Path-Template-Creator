package pathstatlinux

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type RwxSimple struct {
	HyphenedRwxValue string
	RwxWrapper       *chmodhelper.RwxWrapper
	ErrorWrapper     *errorwrapper.Wrapper
	IsRwxValid       bool
}

func InvalidRwxSimple() *RwxSimple {
	return &RwxSimple{
		HyphenedRwxValue: constants.EmptyString,
		RwxWrapper:       nil,
		ErrorWrapper:     errnew.Type.Create(errtype.ChmodInvalid),
		IsRwxValid:       false,
	}
}

// func (receiver *RwxSimple) ApplyOnLinuxPaths()  {
// 	// return receiver.RwxWrapper.clo
// }
//
// func (receiver *RwxSimple) ApplyOnLinuxPaths()  {
//
// }
