package mics

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
)

func GetFileInfosUsingFilteredPathFileInfoMap(
	infoMap *chmodhelper.FilteredPathFileInfoMap,
) []os.FileInfo {
	if infoMap == nil || infoMap.IsEmptyValidFileInfos() {
		return []os.FileInfo{}
	}

	fileInfos := make(
		[]os.FileInfo,
		constants.Zero,
		len(infoMap.FilesToInfoMap))

	for _, fileInfo := range infoMap.FilesToInfoMap {
		fileInfos = append(fileInfos, fileInfo)
	}

	return fileInfos
}
