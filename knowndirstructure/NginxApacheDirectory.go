package knowndirstructure

import (
	"os"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/createdirinternal"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/internal/pathgetterinternal"
)

type NginxApacheDirectory struct {
	DirChmod         os.FileMode // todo fix for apache
	Root             string      `json:"Root,omitempty"`
	RootConfigFile   string      `json:"RootConfigFile,omitempty"`
	ConfigAvailable  string      `json:"ConfigAvailable,omitempty"`
	ConfigEnabled    string      `json:"ConfigEnabled,omitempty"`
	SitesBackup      string      `json:"SitesBackup,omitempty"`
	SitesAvailable   string      `json:"SitesAvailable,omitempty"`
	SitesEnabled     string      `json:"SitesEnabled,omitempty"`
	ExtraConfig      string      `json:"ExtraConfig,omitempty"`
	ModulesAvailable string      `json:"ModulesAvailable,omitempty"`
	ModulesEnabled   string      `json:"ModulesEnabled,omitempty"`
	ApachePorts      string      `json:"ApachePorts,omitempty"`   // root + ports.conf
	ApacheEnvVars    string      `json:"ApacheEnvVars,omitempty"` // root + env-vars
}

func (it *NginxApacheDirectory) IsRootExist() bool {
	return fsinternal.IsPathExists(it.Root)
}

func (it *NginxApacheDirectory) IsRootConfigFile() bool {
	return fsinternal.IsPathExists(it.RootConfigFile)
}

func (it *NginxApacheDirectory) IsConfigAvailable() bool {
	return fsinternal.IsPathExists(it.ConfigAvailable)
}

func (it *NginxApacheDirectory) IsConfigEnabled() bool {
	return fsinternal.IsPathExists(it.ConfigEnabled)
}

func (it *NginxApacheDirectory) IsSitesBackup() bool {
	return fsinternal.IsPathExists(it.SitesBackup)
}

func (it *NginxApacheDirectory) IsSitesAvailable() bool {
	return fsinternal.IsPathExists(it.SitesAvailable)
}

func (it *NginxApacheDirectory) IsSitesEnabled() bool {
	return fsinternal.IsPathExists(it.SitesEnabled)
}

func (it *NginxApacheDirectory) IsExtraConfig() bool {
	return fsinternal.IsPathExists(it.ExtraConfig)
}

func (it *NginxApacheDirectory) IsModulesAvailable() bool {
	return fsinternal.IsPathExists(it.ModulesAvailable)
}

func (it *NginxApacheDirectory) IsModulesEnabled() bool {
	return fsinternal.IsPathExists(it.ModulesEnabled)
}

func (it *NginxApacheDirectory) MkDirRoot(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.Root, mode)
}

func (it *NginxApacheDirectory) MkDirConfigAvailable(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.ConfigAvailable, mode)
}

func (it *NginxApacheDirectory) MkDirConfigEnabled(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.ConfigEnabled, mode)
}

func (it *NginxApacheDirectory) MkDirSitesBackup(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.SitesBackup, mode)
}

func (it *NginxApacheDirectory) MkDirSitesAvailable(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.SitesAvailable, mode)
}

func (it *NginxApacheDirectory) MkDirSitesEnabled(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.SitesEnabled, mode)
}

func (it *NginxApacheDirectory) MkDirExtraConfig(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.ExtraConfig, mode)
}

func (it *NginxApacheDirectory) MkDirModulesAvailable(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.ModulesAvailable, mode)
}

func (it *NginxApacheDirectory) MkDirModulesEnabled(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.ModulesEnabled, mode)
}

func (it *NginxApacheDirectory) MkDirAll(mode os.FileMode) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	errCollection.AddWrapperPtr(it.MkDirRoot(mode))
	errCollection.AddWrapperPtr(it.MkDirConfigAvailable(mode))
	errCollection.AddWrapperPtr(it.MkDirConfigEnabled(mode))
	errCollection.AddWrapperPtr(it.MkDirSitesBackup(mode))
	errCollection.AddWrapperPtr(it.MkDirSitesAvailable(mode))
	errCollection.AddWrapperPtr(it.MkDirSitesEnabled(mode))
	errCollection.AddWrapperPtr(it.MkDirExtraConfig(mode))
	errCollection.AddWrapperPtr(it.MkDirModulesAvailable(mode))
	errCollection.AddWrapperPtr(it.MkDirModulesEnabled(mode))

	return errCollection
}

func (it *NginxApacheDirectory) MkDirAllDefault() *errwrappers.Collection {
	return it.MkDirAll(it.DirChmod)
}

func (it NginxApacheDirectory) Json() corejson.Result {
	return corejson.New(it)
}

func (it NginxApacheDirectory) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it NginxApacheDirectory) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func (it NginxApacheDirectory) JsonModelAny() interface{} {
	return it
}

func (it *NginxApacheDirectory) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *NginxApacheDirectory) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *NginxApacheDirectory) AsJsoner() corejson.Jsoner {
	return it
}

func (it *NginxApacheDirectory) CombinedSitesAvailable(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		it.SitesAvailable,
		combinedPaths)
}

func (it *NginxApacheDirectory) CombinedSitesEnabled(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		it.SitesEnabled,
		combinedPaths)
}

func (it *NginxApacheDirectory) CombinedRoot(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		it.Root,
		combinedPaths)
}

func (it *NginxApacheDirectory) AllFilesAtSitesBackup() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.SitesBackup)
}

func (it *NginxApacheDirectory) AllFilesAtSitesAvailable() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.SitesAvailable)
}

func (it *NginxApacheDirectory) AllFilesAtSitesEnabled() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.Root)
}

func (it *NginxApacheDirectory) AllPathsAtRoot() *errstr.Results {
	return pathgetterinternal.GetAllPaths(
		true,
		osconsts.PathSeparator,
		it.Root)
}

func (it *NginxApacheDirectory) AllFilesAtRoot() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.Root)
}

func (it *NginxApacheDirectory) IsAllExist() bool {
	return it.IsRootExist() &&
		it.IsRootConfigFile() &&
		it.IsConfigAvailable() &&
		it.IsConfigEnabled() &&
		it.IsSitesAvailable() &&
		it.IsSitesEnabled() &&
		it.IsExtraConfig() &&
		it.IsModulesAvailable()
}
