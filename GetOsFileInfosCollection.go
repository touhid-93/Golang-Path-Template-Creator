package pathhelper

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/fileinfogetter"
	"gitlab.com/evatix-go/pathhelper/osfileinfos"
)

func GetOsFileInfosCollection(
	allPaths *[]string,
) (
	infos *osfileinfos.Collection,
	errsCollection *errwrappers.Collection,
) {
	rawInfos, errWrappersCollection := fileinfogetter.GetWithErrors(allPaths)

	return osfileinfos.New(rawInfos), errWrappersCollection
}
