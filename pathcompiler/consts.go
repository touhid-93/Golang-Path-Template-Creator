package pathcompiler

import "gitlab.com/evatix-go/pathhelper/pathsconst"

const (
	AppName                 = pathsconst.AppName
	AppNameLower            = pathsconst.AppNameLower
	defaultConfigRootSuffix = pathsconst.DefaultConfigRootSuffix // "/config"
	packagesDirName         = pathsconst.PackagesDirName         // "/packages/"
	instructionDirName      = pathsconst.InstructionDirName      // "/instructions/"
	unixLogAppRoot          = pathsconst.UnixLogAppRoot          // /var/opt/{app-name}
	unixVarAppRoot          = pathsconst.UnixVarAppRoot          // /var/opt/{app-name}
	etcApp                  = pathsconst.EtcApp                  // /etc/{app-name}
	unixConfigRoot          = pathsconst.UnixConfigRoot          // /etc/{app-name}/config
)
