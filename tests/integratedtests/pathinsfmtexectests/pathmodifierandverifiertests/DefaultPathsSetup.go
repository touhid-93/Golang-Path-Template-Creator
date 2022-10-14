package pathinsfmtexectests

import (
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers/pathinsfmtexectestwrappers"
)

func DefaultPathsSetup() {
	deletepaths.AllOnExist(
		pathinsfmtexectestwrappers.PathOneTextFile,
		pathinsfmtexectestwrappers.PathTwoTextFile,
	).HandleError()

	_, errWrap := createpath.CreateMany(
		true,
		true,
		[]string{
			pathinsfmtexectestwrappers.PathOneTextFile,
			pathinsfmtexectestwrappers.PathTwoTextFile,
		},
	)

	errWrap.HandleError()
}
