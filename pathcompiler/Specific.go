package pathcompiler

import (
	"path"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

type Specific struct {
	Name                     string            `json:"Name,omitempty"`
	Description              string            `json:"Description,omitempty"`
	Url                      string            `json:"Url,omitempty"`
	ReplacerMap              map[string]string `json:"ReplacerMap,omitempty"` // todo later add dynamic way of getting format path
	SpecificPathFileLocation string            // eg. self saving location, unix : "/var/opt/{app-name}/defined-paths/paths.json"
	VarAppRoot               string            // eg. unix : "/var/opt/{app-name}"
	EtcAppRoot               string            // eg. unix : "/etc/{app-name}"
	EtcAppConfigRoot         string            // eg. unix : "/etc/{app-name}/config"
	AppDbRoot                string            // eg. unix : "/var/opt/{app-name}/databases"
	TempRoot                 string            // eg. unix : TempAppRoot => TempPermanentDir + AppNameLower => "/{os-temp}/{app-name}"
	UserTempRoot             string            // eg. unix : "/{os-temp}/{app-name}/users/"
	CacheTempRoot            string            // eg. unix : "/{os-temp}/{app-name}/cache/"
	InstructionTempRoot      string            // eg. unix : "/{os-temp}/{app-name}/instructions/"
	MigrationCacheRoot       string            // eg. unix : "/{os-temp}/{app-name}/migration-cache/"
	PackageTempRoot          string            // eg. unix : "/{os-temp}/{app-name}/packages/"
	LogAppRoot               string            // eg. unix : pathsconst.UnixLogAppRoot + {app-name} => "/var/log/{app-name}/"
	VarCacheRoot             string            // eg. unix : pathsconst.UnixVarAppRoot + "cache" => "/var/opt/{app-name}/cache"
	DownloadsRoot            string            // eg. unix : pathsconst.UnixVarAppRoot + "downloads" => "/var/opt/{app-name}/downloads"
	ScriptsRoot              string            // eg. unix : pathsconst.UnixVarAppRoot + "scripts" => "/var/opt/{app-name}/scripts"
	DecompressRoot           string            // eg. unix : pathsconst.UnixVarAppRoot + "decompress" => "/var/opt/{app-name}/decompress"
	PackagesRoot             string            // eg. unix : pathsconst.UnixVarAppRoot + "packages" => "/var/opt/{app-name}/packages"
	PackagesDownloadRoot     string            // eg. unix : pathsconst.UnixVarAppRoot + "packages-downloaded" => "/var/opt/{app-name}/packages-downloaded/"
	DefaultInstructionsRoot  string            // eg. unix : pathsconst.UnixVarAppRoot + "/instructions/" => "/var/opt/{app-name}/instructions/"
	DefaultEnvRoot           string            // eg. unix : pathsconst.UnixVarAppRoot + "/env/" => "/var/opt/{app-name}/env/"
	DefaultEnvPathRoot       string            // eg. unix : pathsconst.UnixVarAppRoot + "/env-paths/" => "/var/opt/{app-name}/env-paths/"
	BackupRoot               string            // eg. unix : pathsconst.UnixVarAppRoot + "/backups/" => "/var/opt/{app-name}/backups/"
	ArchiveRoot              string            // eg. unix : pathsconst.UnixVarAppRoot + "/archived/" => "/var/opt/{app-name}/archived/"
	ZipsRoot                 string            // eg. unix : pathsconst.UnixVarAppRoot + "/compressed/" => "/var/opt/{app-name}/compressed/"
	DefaultConfigFilePath    string            // eg. unix : pathsconst.UnixVarAppRoot + "/config/default-config.json" => "/var/opt/{app-name}/config/default-config.json"
	SnapshotsRoot            string            // eg. unix : "/var/opt/{app-name}-snapshots/"
	PublicRoot               string            // eg. unix : "/var/www"
	SslRoot                  string            // eg. unix : "/var/opt/{app-name}-ssl/"
}

func (it *Specific) IsNull() bool {
	return it == nil
}

func (it *Specific) IsDefined() bool {
	return it != nil
}

func (it *Specific) JoinWithOptions(
	isSkipEmpty bool,
	isNormalize bool,
	isExpand bool,
	baseDir string,
	relativePaths ...string,
) string {
	return pathjoin.JoinWithBaseDirSep(
		isSkipEmpty,
		isExpand,
		isNormalize,
		osconsts.PathSeparator,
		baseDir,
		relativePaths...)
}

func (it *Specific) JoinWith(
	baseDir string,
	relativePaths ...string,
) string {
	simpleJoin := it.JoinSimpleWith(
		baseDir,
		relativePaths...)

	return path.Clean(simpleJoin)
}

func (it *Specific) JoinNormalized(
	baseDir string,
	relativePaths ...string,
) string {
	return pathjoin.JoinWithBaseDirSep(
		true,
		false,
		true,
		osconsts.PathSeparator,
		baseDir,
		relativePaths...)
}

func (it *Specific) JoinNormalizedExpand(
	baseDir string,
	relativePaths ...string,
) string {
	return pathjoin.JoinWithBaseDirSep(
		true,
		true,
		true,
		osconsts.PathSeparator,
		baseDir,
		relativePaths...)
}

func (it *Specific) JoinSimpleWith(
	baseDir string,
	relativePaths ...string,
) string {
	return pathjoin.JoinSimpleBaseWithMany(
		baseDir,
		relativePaths...)
}

func (it *Specific) JoinWithVarAppRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.VarAppRoot,
		relativePaths...)
}

func (it *Specific) JoinWithEtcAppRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.EtcAppRoot,
		relativePaths...)
}

func (it *Specific) JoinWithAppConfigRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.EtcAppConfigRoot,
		relativePaths...)
}

func (it *Specific) JoinWithAppDbRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.AppDbRoot,
		relativePaths...)
}

func (it *Specific) JoinWithTempRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.TempRoot,
		relativePaths...)
}

func (it *Specific) JoinWithAppTempRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.TempRoot,
		relativePaths...)
}

func (it *Specific) JoinWithTempWithoutApp(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		pathsconst.TempPermanentDir,
		relativePaths...)
}

func (it *Specific) JoinWithUserTempRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.UserTempRoot,
		relativePaths...)
}

func (it *Specific) JoinWithCacheTempRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.CacheTempRoot,
		relativePaths...)
}

func (it *Specific) JoinWithInstructionTempRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.InstructionTempRoot,
		relativePaths...)
}

func (it *Specific) JoinWithMigrationCacheRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.MigrationCacheRoot,
		relativePaths...)
}

func (it *Specific) JoinWithPackageTempRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.PackageTempRoot,
		relativePaths...)
}

func (it *Specific) JoinWithLogAppRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.LogAppRoot,
		relativePaths...)
}

func (it *Specific) JoinWithVarCacheRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.VarCacheRoot,
		relativePaths...)
}

func (it *Specific) JoinWithDownloadsRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.DownloadsRoot,
		relativePaths...)
}

func (it *Specific) JoinWithScriptsRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.ScriptsRoot,
		relativePaths...)
}

func (it *Specific) JoinWithDecompressRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.DecompressRoot,
		relativePaths...)
}

func (it *Specific) JoinWithPackagesRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.PackagesRoot,
		relativePaths...)
}

func (it *Specific) JoinWithPackagesDownloadRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.PackagesDownloadRoot,
		relativePaths...)
}

func (it *Specific) JoinWithDefaultInstructionsRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.DefaultInstructionsRoot,
		relativePaths...)
}

func (it *Specific) JoinWithDefaultEnvRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.DefaultEnvRoot,
		relativePaths...)
}

func (it *Specific) JoinWithDefaultEnvPathRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.DefaultEnvPathRoot,
		relativePaths...)
}

func (it *Specific) JoinWithBackupRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.BackupRoot,
		relativePaths...)
}

func (it *Specific) JoinWithArchiveRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.ArchiveRoot,
		relativePaths...)
}

func (it *Specific) JoinWithZipsRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.ZipsRoot,
		relativePaths...)
}

func (it *Specific) JoinWithSnapshotsRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.SnapshotsRoot,
		relativePaths...)
}

func (it *Specific) JoinWithPublicRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.PublicRoot,
		relativePaths...)
}

func (it *Specific) JoinWithSslRoot(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.SslRoot,
		relativePaths...)
}

func (it *Specific) ExecutableDirPath() string {
	return pathsconst.ExecutableDir
}

func (it *Specific) JoinWithExecutableDirPath(
	relativePaths ...string,
) string {
	return it.JoinNormalized(
		it.ExecutableDirPath(),
		relativePaths...)
}

func (it *Specific) CloneReplacerMap() map[string]string {
	if it == nil || len(it.ReplacerMap) == 0 {
		return map[string]string{}
	}

	newMap := make(map[string]string, len(it.ReplacerMap))

	for key, value := range it.ReplacerMap {
		newMap[key] = value
	}

	return newMap
}

func (it *Specific) ClonePtr() *Specific {
	if it == nil {
		return &Specific{}
	}

	return &Specific{
		Name:                     it.Name,
		Description:              it.Description,
		Url:                      it.Url,
		ReplacerMap:              it.CloneReplacerMap(),
		SpecificPathFileLocation: it.SpecificPathFileLocation,
		VarAppRoot:               it.VarAppRoot,
		EtcAppRoot:               it.EtcAppRoot,
		EtcAppConfigRoot:         it.EtcAppConfigRoot,
		AppDbRoot:                it.AppDbRoot,
		TempRoot:                 it.TempRoot,
		UserTempRoot:             it.UserTempRoot,
		CacheTempRoot:            it.CacheTempRoot,
		InstructionTempRoot:      it.InstructionTempRoot,
		MigrationCacheRoot:       it.MigrationCacheRoot,
		PackageTempRoot:          it.PackageTempRoot,
		LogAppRoot:               it.LogAppRoot,
		VarCacheRoot:             it.VarCacheRoot,
		DownloadsRoot:            it.DownloadsRoot,
		ScriptsRoot:              it.ScriptsRoot,
		DecompressRoot:           it.DecompressRoot,
		PackagesRoot:             it.PackagesRoot,
		PackagesDownloadRoot:     it.PackagesDownloadRoot,
		DefaultInstructionsRoot:  it.DefaultInstructionsRoot,
		DefaultEnvRoot:           it.DefaultEnvRoot,
		DefaultEnvPathRoot:       it.DefaultEnvPathRoot,
		BackupRoot:               it.BackupRoot,
		ArchiveRoot:              it.ArchiveRoot,
		ZipsRoot:                 it.ZipsRoot,
		DefaultConfigFilePath:    it.DefaultConfigFilePath,
		SnapshotsRoot:            it.SnapshotsRoot,
		PublicRoot:               it.PublicRoot,
		SslRoot:                  it.SslRoot,
	}
}

func (it *Specific) Clone() Specific {
	cloned := it.ClonePtr()

	return cloned.ToNonPtr()
}

func (it Specific) ToPtr() *Specific {
	return &it
}

func (it *Specific) ToNonPtr() Specific {
	if it == nil {
		return Specific{}
	}

	return *it
}

func (it *Specific) ReflectSetTo(toPtr interface{}) error {
	return coredynamic.ReflectSetFromTo(it, toPtr)
}

func (it *Specific) ReflectSetToErrWrap(toPtr interface{}) *errorwrapper.Wrapper {
	return errnew.Reflect.SetFromTo(it, toPtr)
}

func (it *Specific) PrettyJsonString() string {
	return corejson.NewPtr(it).PrettyJsonString()
}

func (it *Specific) Json() corejson.Result {
	return corejson.New(it)
}

func (it *Specific) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Specific) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it Specific) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
