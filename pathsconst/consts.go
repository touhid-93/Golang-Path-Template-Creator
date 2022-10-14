package pathsconst

import "gitlab.com/evatix-go/core/appname"

const (
	AppName                 = appname.Cimux
	AppNameLower            = appname.CimuxLower
	VarOpt                  = "/var/opt/"
	Etc                     = "/etc/"
	VarLog                  = "/var/log/"
	TestDirPatternName      = AppNameLower + "-tests"
	DefaultConfigRootSuffix = "/config"
	PackagesDirName         = "/packages/"
	InstructionDirName      = "/instructions/"
	UnixLogAppRoot          = VarLog + AppNameLower
	UnixVarAppRoot          = VarOpt + AppNameLower
	EtcApp                  = Etc + AppNameLower
	UnixConfigRoot          = EtcApp + DefaultConfigRootSuffix
	WindowsDir              = "c:\\Windows"
)
