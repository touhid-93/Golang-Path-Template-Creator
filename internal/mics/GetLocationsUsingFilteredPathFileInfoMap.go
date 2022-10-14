package mics

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
)

func GetLocationsUsingFilteredPathFileInfoMap(
	infoMap *chmodhelper.FilteredPathFileInfoMap,
) []string {
	if infoMap == nil || infoMap.IsEmptyValidFileInfos() {
		return []string{}
	}

	locations := make(
		[]string,
		constants.Zero,
		len(infoMap.FilesToInfoMap))

	for s := range infoMap.FilesToInfoMap {
		locations = append(locations, s)
	}

	return locations
}
