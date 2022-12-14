package pathcompiler

var DefaultSuffix = Suffixes{
	Specific{
		Name:                     "Suffix map",
		Description:              "Only suffix",
		SpecificPathFileLocation: "/defined-paths/paths.json",
		VarAppRoot:               AppNameLower,
		EtcAppRoot:               AppNameLower,
		EtcAppConfigRoot:         defaultConfigRootSuffix,
		AppDbRoot:                "/databases/",
		TempRoot:                 AppNameLower,
		UserTempRoot:             "/users/",
		CacheTempRoot:            "/cache/",
		InstructionTempRoot:      instructionDirName,
		MigrationCacheRoot:       "/migration-cache/",
		PackageTempRoot:          packagesDirName,
		LogAppRoot:               AppNameLower,
		VarCacheRoot:             "/cache/",
		DownloadsRoot:            "/downloads/",
		ScriptsRoot:              "/scripts/",
		DecompressRoot:           "/decompress/",
		PackagesRoot:             packagesDirName,
		PackagesDownloadRoot:     "/packages-downloaded/",
		DefaultInstructionsRoot:  instructionDirName,
		DefaultEnvRoot:           "/env/",
		DefaultEnvPathRoot:       "/env-paths/",
		BackupRoot:               "/backups/",
		ArchiveRoot:              "/archived/",
		ZipsRoot:                 "/compressed/",
		DefaultConfigFilePath:    defaultConfigRootSuffix + "/default-config.json",
		SnapshotsRoot:            "-snapshots/",
		PublicRoot:               "www/",
		SslRoot:                  "-ssl/",
	},
}
