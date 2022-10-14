package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func nginxPathTest01() {
	nginxRoot := pathsconst.TempPermanentDir + "/nginx-test01/"

	nginxDir := nginxlinuxpath.NewNginxDirDefaultChmod(
		true,
		nginxRoot,
		"alim")

	fmt.Println(nginxDir.JsonString())
	nginxDir.MkDirDefaultSpecific().HandleError()

	alimSite01 := nginxDir.AbsPathOfAvailableSiteAddConfExt("alim-site-01")

	fs.WriteStringToFile(
		false,
		alimSite01,
		"some with alim").HandleError()

	fs.WriteStringToFile(
		false,
		nginxDir.CurrentUserRootConfig,
		"some with alim - root").HandleError()

	rs2 := nginxDir.CopyCurrentUserRootConfigToTempRel("alim-test2/newroot.conf")

	rs2.ErrorWrapper.HandleError()
	fmt.Println(rs2.String())

	availableSites, err := nginxDir.AllSiteFiles()

	fmt.Println(availableSites)
	fmt.Println(err)
	fmt.Println(nginxDir.AllFilesAtSpecificUserRoot())
	fmt.Println(nginxDir.UsersAvailableSites())
	fmt.Println(nginxDir.UsersEnabledSites())

	nginxDir2 := nginxlinuxpath.NewNginxDirDefaultChmod(
		true,
		nginxRoot,
		"alim2")

	nginxDir2.MkDirDefaultSpecific()
	fmt.Println(nginxDir.AllUsersDirs())
	fmt.Println(nginxDir.AllUsersNames())
}
