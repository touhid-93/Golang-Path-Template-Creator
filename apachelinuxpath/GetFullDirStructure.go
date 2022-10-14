package apachelinuxpath

import "gitlab.com/evatix-go/pathhelper/knowndirstructure"

func GetFullDirStructure(
	isNormalize bool,
	currentRoot string,
) *knowndirstructure.NginxApacheDirectory {
	return &knowndirstructure.NginxApacheDirectory{
		Root:             fixPathIf(isNormalize, currentRoot, ""),
		RootConfigFile:   fixPathIf(isNormalize, currentRoot, RootConfigName),
		ConfigAvailable:  fixPathIf(isNormalize, currentRoot, ConfigAvailableName),
		ConfigEnabled:    fixPathIf(isNormalize, currentRoot, ConfigEnabledName),
		SitesBackup:      fixPathIf(isNormalize, currentRoot, SitesBackup),
		SitesAvailable:   fixPathIf(isNormalize, currentRoot, SitesAvailableName),
		SitesEnabled:     fixPathIf(isNormalize, currentRoot, SitesEnabledName),
		ExtraConfig:      fixPathIf(isNormalize, currentRoot, ExtraConfName),
		ModulesAvailable: fixPathIf(isNormalize, currentRoot, ModulesAvailableName),
		ModulesEnabled:   fixPathIf(isNormalize, currentRoot, ModulesEnabledName),
		ApachePorts:      fixPathIf(isNormalize, currentRoot, ApachePorts),
		ApacheEnvVars:    fixPathIf(isNormalize, currentRoot, ApacheEnvVars),
	}
}
