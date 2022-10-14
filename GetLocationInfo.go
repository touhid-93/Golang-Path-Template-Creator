package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func GetLocationInfo(location string) *LocationInfo {
	baseDir := splitinternal.GetBaseDir(location)
	fileNameWithExt, dotExt := splitinternal.GetFileNameDotExt(location)
	fileName := strings.TrimSuffix(fileNameWithExt, dotExt)
	var ext string

	if len(dotExt) >= constants.One {
		ext = dotExt[1:]
	}

	if baseDir == fileNameWithExt {
		baseDir = ""
	}

	return &LocationInfo{
		RawLocation:           location,
		FileNameWithExtension: fileNameWithExt,
		BaseDir:               baseDir,
		FileName:              fileName,
		DotExtension:          dotExt,
		Extension:             ext,
	}
}
