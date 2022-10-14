package fileinfopath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/namevalue"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

func FileInfoString(fileInfo os.FileInfo) string {
	if fileInfo == nil {
		return constants.EmptyString
	}

	return consts.IndentFileInfoEachLineJoiner + errcore.VarNameValuesJoiner(
		consts.IndentFileInfoEachLineJoiner,
		namevalue.Instance{
			Name:  "Name",
			Value: fileInfo.Name(),
		},
		namevalue.Instance{
			Name:  "Size",
			Value: fileInfo.Size(),
		},
		namevalue.Instance{
			Name:  "LastModifiedDate",
			Value: fileInfo.ModTime(),
		},
		namevalue.Instance{
			Name:  "IsDir",
			Value: fileInfo.IsDir(),
		})
}
