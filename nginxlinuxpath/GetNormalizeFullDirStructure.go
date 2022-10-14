package nginxlinuxpath

import "gitlab.com/evatix-go/pathhelper/knowndirstructure"

func GetNormalizeFullDirStructure(
	currentNginxRoot string,
) *knowndirstructure.NginxApacheDirectory {
	return &knowndirstructure.NginxApacheDirectory{
		Root:             fixPath(currentNginxRoot, ""),
		RootConfigFile:   fixPath(currentNginxRoot, RootConfigName),
		ConfigAvailable:  fixPath(currentNginxRoot, ConfigAvailableName),
		ConfigEnabled:    fixPath(currentNginxRoot, ConfigEnabledName),
		SitesAvailable:   fixPath(currentNginxRoot, SitesAvailableName),
		SitesEnabled:     fixPath(currentNginxRoot, SitesEnabledName),
		ExtraConfig:      fixPath(currentNginxRoot, ExtraConfName),
		ModulesAvailable: fixPath(currentNginxRoot, ModulesAvailableName),
		ModulesEnabled:   fixPath(currentNginxRoot, ModulesEnabledName),
	}
}
