package nginxlinuxpath

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/knowndirstructure"
)

func GetFullDirStructure(
	isNormalize bool,
	dirChmod os.FileMode,
	currentNginxRoot string,
) *knowndirstructure.NginxApacheDirectory {
	return &knowndirstructure.NginxApacheDirectory{
		DirChmod:         dirChmod,
		Root:             fixPathIf(isNormalize, currentNginxRoot, ""),
		RootConfigFile:   fixPathIf(isNormalize, currentNginxRoot, RootConfigName),
		ConfigAvailable:  fixPathIf(isNormalize, currentNginxRoot, ConfigAvailableName),
		ConfigEnabled:    fixPathIf(isNormalize, currentNginxRoot, ConfigEnabledName),
		SitesBackup:      fixPathIf(isNormalize, currentNginxRoot, SitesBackup),
		SitesAvailable:   fixPathIf(isNormalize, currentNginxRoot, SitesAvailableName),
		SitesEnabled:     fixPathIf(isNormalize, currentNginxRoot, SitesEnabledName),
		ExtraConfig:      fixPathIf(isNormalize, currentNginxRoot, ExtraConfName),
		ModulesAvailable: fixPathIf(isNormalize, currentNginxRoot, ModulesAvailableName),
		ModulesEnabled:   fixPathIf(isNormalize, currentNginxRoot, ModulesEnabledName),
	}
}
