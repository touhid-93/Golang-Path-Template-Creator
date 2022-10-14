package testwrappers

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func GetSetupPaths() []string {
	linkedCollection := corestr.Empty.LinkedCollections()

	for _, ins := range PathsCreateInstructionsUnix {
		linkedCollection.AddStrings(
			ins.LazyFlatPaths()...,
		)
	}

	locations := linkedCollection.
		ToCollection(0).
		ListStrings()

	return normalize.PathsUsingSingleIfAsync(
		true, locations)
}
