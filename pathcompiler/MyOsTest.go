package pathcompiler

var MyOsTest = Specific{
	Name:                     "MyOS",
	Description:              "all MyOS related os paths",
	SpecificPathFileLocation: "sys/" + unixVarAppRoot + DefaultSuffix.SpecificPathFileLocation,
	VarAppRoot:               "sys/" + unixVarAppRoot,
	EtcAppRoot:               "sys/" + etcApp,
	EtcAppConfigRoot:         "sys/" + unixConfigRoot,
	AppDbRoot:                "sys/" + unixVarAppRoot + DefaultSuffix.AppDbRoot,
	TempRoot:                 "sys/" + TempAppRoot,
	UserTempRoot:             "sys/" + TempAppRoot + DefaultSuffix.UserTempRoot,
	CacheTempRoot:            "sys/" + TempAppRoot + DefaultSuffix.CacheTempRoot,
	InstructionTempRoot:      "sys/" + TempAppRoot + DefaultSuffix.InstructionTempRoot,
	MigrationCacheRoot:       "sys/" + TempAppRoot + DefaultSuffix.MigrationCacheRoot,
	PackageTempRoot:          "sys/" + TempAppRoot + DefaultSuffix.PackageTempRoot,
	LogAppRoot:               "sys/" + unixLogAppRoot,
	VarCacheRoot:             "sys/" + unixVarAppRoot + DefaultSuffix.VarCacheRoot,
	DownloadsRoot:            "sys/" + unixVarAppRoot + DefaultSuffix.DownloadsRoot,
	ScriptsRoot:              "sys/" + unixVarAppRoot + DefaultSuffix.ScriptsRoot,
	DecompressRoot:           "sys/" + TempAppRoot + DefaultSuffix.DecompressRoot,
	PackagesRoot:             "sys/" + etcApp + DefaultSuffix.PackagesRoot,
	PackagesDownloadRoot:     "sys/" + unixVarAppRoot + DefaultSuffix.PackagesDownloadRoot,
	DefaultInstructionsRoot:  "sys/" + unixVarAppRoot + DefaultSuffix.DefaultInstructionsRoot,
	DefaultEnvRoot:           "sys/" + unixVarAppRoot + DefaultSuffix.DefaultEnvRoot,
	DefaultEnvPathRoot:       "sys/" + unixVarAppRoot + DefaultSuffix.DefaultEnvPathRoot,
	BackupRoot:               "sys/" + unixVarAppRoot + DefaultSuffix.BackupRoot,
	ArchiveRoot:              "sys/" + unixVarAppRoot + DefaultSuffix.ArchiveRoot,
	ZipsRoot:                 "sys/" + unixVarAppRoot + DefaultSuffix.ZipsRoot,
	DefaultConfigFilePath:    "sys/" + etcApp + DefaultSuffix.DefaultConfigFilePath,
	SnapshotsRoot:            "sys/" + unixVarAppRoot + DefaultSuffix.SnapshotsRoot,
	PublicRoot:               "sys/" + "/var/" + DefaultSuffix.PublicRoot,
	SslRoot:                  "sys/" + unixVarAppRoot + DefaultSuffix.SslRoot,
}
