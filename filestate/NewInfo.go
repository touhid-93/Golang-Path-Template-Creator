package filestate

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func NewInfo(
	hashMethod hashas.Variant,
	isNormalize bool,
	filePath string,
) (*Info, *errorwrapper.Wrapper) {
	finalPath := normalize.PathIf(isNormalize, filePath)
	stat := chmodhelper.GetPathExistStat(finalPath)

	if stat.HasError() {
		return InvalidInfo(hashMethod, finalPath), errnew.
			Path.
			Error(
				errtype.PathMissingOrInvalid,
				stat.MeaningFullError(),
				finalPath)
	}

	isInvalidFile := !stat.IsFile()
	hexChecksumFileResult := hashMethod.HexSumOfFileIf(
		isInvalidFile,
		finalPath)

	if hexChecksumFileResult.HasError() {
		return InvalidInfo(hashMethod, finalPath), hexChecksumFileResult.ErrorWrapper
	}

	isFile := stat.IsFile()
	size := stat.FileInfo.Size()
	lastModifiedDate := stat.FileInfo.ModTime()
	chmod := stat.FileInfo.Mode()
	fileInfoPath := fileinfopath.NewUsingStat(filePath, stat)

	return &Info{
		FullPath:           finalPath,
		LastModified:       lastModifiedDate,
		Chmod:              chmod,
		UserGroupId:        pathsysinfo.GetUserGroupIdUsingUnix(fileInfoPath),
		Size:               size,
		HexContentChecksum: hexChecksumFileResult.Value,
		IsFile:             isFile,
		IsInvalid:          !stat.IsExist,
		HashMethod:         hashMethod,
	}, nil
}
