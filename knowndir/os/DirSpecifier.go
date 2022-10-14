package os

import "gitlab.com/evatix-go/enum/osarchs"

type DirSpecifier interface {
	Arch32() string
	Arch64() string
	Arch32Ptr() *string
	Arch64Ptr() *string
	Generated() string
	String() string
	GetDir(architecture osarchs.Architecture) string
}
