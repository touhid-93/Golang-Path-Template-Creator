package pathcompiler

import (
	"gitlab.com/evatix-go/core/osconsts"
)

var WindowsOs = Specific{
	Name:                     "Windows",
	Description:              "all windows related os paths",
	Url:                      "",
	SpecificPathFileLocation: windowsProductionDefault + DefaultSuffix.SpecificPathFileLocation,
	VarAppRoot:               windowsProductionDefault,
	EtcAppRoot:               windowsProductionDefault,
	EtcAppConfigRoot:         windowsProductionDefault + DefaultSuffix.EtcAppConfigRoot,
	AppDbRoot:                windowsProductionDefault + DefaultSuffix.AppDbRoot,
	TempRoot:                 AppWindowsUserSpecificTempRoot,
	UserTempRoot:             AppWindowsUserSpecificTempRoot + DefaultSuffix.UserTempRoot,
	CacheTempRoot:            AppWindowsUserSpecificTempRoot + DefaultSuffix.CacheTempRoot,
	InstructionTempRoot:      AppWindowsUserSpecificTempRoot + DefaultSuffix.InstructionTempRoot,
	MigrationCacheRoot:       AppWindowsUserSpecificTempRoot + DefaultSuffix.MigrationCacheRoot,
	PackageTempRoot:          AppWindowsUserSpecificTempRoot + DefaultSuffix.PackageTempRoot,
	LogAppRoot:               windowsProductionDefault + "\\logs",
	VarCacheRoot:             windowsProductionDefault + DefaultSuffix.VarCacheRoot,
	DownloadsRoot:            TempAppRoot + DefaultSuffix.DownloadsRoot,
	ScriptsRoot:              windowsProductionDefault + DefaultSuffix.ScriptsRoot,
	DecompressRoot:           TempAppRoot + DefaultSuffix.DecompressRoot,
	PackagesRoot:             windowsProductionDefault + DefaultSuffix.PackagesRoot,
	PackagesDownloadRoot:     windowsProductionDefault + DefaultSuffix.PackagesDownloadRoot,
	DefaultInstructionsRoot:  windowsProductionDefault + DefaultSuffix.DefaultInstructionsRoot,
	DefaultEnvRoot:           windowsProductionDefault + DefaultSuffix.DefaultEnvRoot,
	DefaultEnvPathRoot:       windowsProductionDefault + DefaultSuffix.DefaultEnvPathRoot,
	BackupRoot:               windowsProductionDefault + DefaultSuffix.BackupRoot,
	ArchiveRoot:              windowsProductionDefault + DefaultSuffix.ArchiveRoot,
	ZipsRoot:                 windowsProductionDefault + DefaultSuffix.ZipsRoot,
	DefaultConfigFilePath:    windowsProductionDefault + DefaultSuffix.DefaultConfigFilePath,
	SnapshotsRoot:            windowsProductionDefault + "\\config" + DefaultSuffix.SnapshotsRoot,
	PublicRoot:               osconsts.WindowsCDrive + "public_root\\" + DefaultSuffix.PublicRoot,
	SslRoot:                  windowsProductionDefault + "\\all" + DefaultSuffix.SslRoot,
}
