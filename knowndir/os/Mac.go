package os

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/enum/osarchs"
)

type Mac struct {
	X32, X64  string
	generated *string
}

func (it *Mac) Arch32() string {
	return it.X32
}

func (it *Mac) Arch64() string {
	return it.X64
}

func (it *Mac) Arch32Ptr() *string {
	return &it.X32
}

func (it *Mac) Arch64Ptr() *string {
	return &it.X64
}

func (it *Mac) GetDir(architecture osarchs.Architecture) string {
	if architecture.IsX32() {
		return it.Arch32()
	}

	return it.Arch64()
}

func (it *Mac) Generated() string {
	if it.generated != nil {
		return *it.generated
	}

	if osconsts.IsX64Architecture {
		it.generated = it.Arch64Ptr()
	} else {
		it.generated = it.Arch32Ptr()
	}

	return *it.generated
}

func (it Mac) String() string {
	return it.Generated()
}
