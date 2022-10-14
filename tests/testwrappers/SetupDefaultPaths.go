package testwrappers

import (
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathscreateinsexec"
)

func SetupDefaultPathsUnix() []string {
	errCollection := pathscreateinsexec.ApplyPathsCreatorCollectionsReturnErrorCollection(
		true,
		true,
		true,
		false,
		PathsCreateInstructionsUnix)

	if errCollection.HasError() {
		errCollection.HandleWithMsg(
			"Failed to create default paths." + AllPathsString)
	}

	return DefaultWorkingPaths
}
