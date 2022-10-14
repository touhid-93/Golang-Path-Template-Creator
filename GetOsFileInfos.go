package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/fileinfogetter"
)

func GetOsFileInfos(
	allPaths *[]string,
) (
	infos *[]os.FileInfo,
	errsCollection *errwrappers.Collection,
) {
	return fileinfogetter.GetWithErrors(allPaths)
}
