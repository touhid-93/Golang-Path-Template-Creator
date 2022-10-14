package pathhelper

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type LocationInfo struct {
	RawLocation,
	FileNameWithExtension,
	BaseDir,
	FileName,
	DotExtension,
	Extension string
}

func (it *LocationInfo) IsInvalidLocation() bool {
	return it.RawLocation == ""
}

func (it *LocationInfo) HasLocation() bool {
	return it.RawLocation != ""
}

func (it *LocationInfo) HasFileNameWithExtension() bool {
	return it.FileNameWithExtension != ""
}

func (it *LocationInfo) IsInvalidFileNameWithExtension() bool {
	return it.FileNameWithExtension == ""
}

func (it *LocationInfo) HasBaseDir() bool {
	return it.BaseDir != ""
}

func (it *LocationInfo) IsInvalidBaseDir() bool {
	return it.BaseDir == ""
}

func (it *LocationInfo) HasFileName() bool {
	return it.FileName != ""
}

func (it *LocationInfo) IsInvalidFileName() bool {
	return it.DotExtension == ""
}

func (it *LocationInfo) HasDotExtension() bool {
	return it.DotExtension != ""
}

func (it *LocationInfo) IsInvalidDotExtension() bool {
	return it.DotExtension == ""
}

func (it *LocationInfo) HasExtension() bool {
	return it.Extension != ""
}

func (it *LocationInfo) IsInvalidExtension() bool {
	return it.Extension == ""
}

func (it *LocationInfo) IsExtension(ext string, isIgnoreCase bool) bool {
	if isIgnoreCase {
		return strings.EqualFold(it.Extension, ext)
	}

	return it.Extension == ext
}

func (it *LocationInfo) IsDotExtension(dotExt string, isIgnoreCase bool) bool {
	if isIgnoreCase {
		return strings.EqualFold(it.DotExtension, dotExt)
	}

	return it.DotExtension == dotExt
}

func (it *LocationInfo) IsFileName(fileName string, isIgnoreCase bool) bool {
	if isIgnoreCase {
		return strings.EqualFold(it.FileName, fileName)
	}

	return it.FileName == fileName
}

func (it *LocationInfo) IsFileNameExt(fileNameExt string, isIgnoreCase bool) bool {
	if isIgnoreCase {
		return strings.EqualFold(it.FileNameWithExtension, fileNameExt)
	}

	return it.FileNameWithExtension == fileNameExt
}

func (it *LocationInfo) PathCombineWithBase(
	isNormalize,
	isExpandEnv bool,
	relativePath string,
) string {
	return PathCombineWithBase(
		isNormalize,
		isExpandEnv,
		it.BaseDir,
		relativePath)
}

func (it *LocationInfo) PathCombineWithBaseAsLocationInfo(
	isNormalize,
	isExpandEnv bool,
	relativePath string,
) *LocationInfo {
	location := PathCombineWithBase(
		isNormalize,
		isExpandEnv,
		it.BaseDir,
		relativePath)

	return GetLocationInfo(location)
}

func (it *LocationInfo) PathSimpleStat() *pathchmod.SimpleStat {
	return pathchmod.GetSimpleStat(it.RawLocation)
}

func (it LocationInfo) String() string {
	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		it)
}
