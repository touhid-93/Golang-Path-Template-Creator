package nginxlinuxpath

import (
	"os"

	"gitlab.com/evatix-go/asynchelper/syncparallel"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/extensionsconst"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/knowndirstructure"
)

func NewNginxDir(
	isNormalize bool,
	dirChmod os.FileMode,
	currentNginxRoot,
	username string,
) *NginxDir {
	var nginxRoot, userDir *knowndirstructure.NginxApacheDirectory
	var specificUserRoot, specificUserRootConfig string
	allUsersRoot := normalizeinternal.JoinPathsFixIf(
		isNormalize,
		currentNginxRoot,
		ExtraConfName,
		Users,
	)

	syncparallel.Tasks(
		func() {
			nginxRoot = GetFullDirStructure(
				isNormalize,
				dirChmod,
				currentNginxRoot)
		},
		func() {
			specificUserRoot = normalizeinternal.JoinPathsFixIf(
				isNormalize,
				allUsersRoot,
				username)

			specificUserRootConfig = normalizeinternal.JoinPathsFixIf(
				isNormalize,
				specificUserRoot,
				username+extensionsconst.DotConf,
			)

			userDir = GetFullDirStructure(
				isNormalize,
				dirChmod,
				specificUserRoot)
		})

	return &NginxDir{
		BaseUsername:          *coreinstruction.NewUsername(username),
		Root:                  nginxRoot,
		User:                  userDir,
		AllUsersRoot:          allUsersRoot,
		SpecificUserRoot:      specificUserRoot,
		CurrentUserRootConfig: specificUserRootConfig,
	}
}
