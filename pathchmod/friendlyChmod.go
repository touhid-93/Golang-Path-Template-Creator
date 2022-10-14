package pathchmod

import (
	"fmt"
	"path/filepath"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/core/simplewrap"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type friendlyChmod struct{}

func (it friendlyChmod) OfPath(
	location string,
) string {
	fileChmod, isInvalid := chmodhelper.
		GetExistingChmodOfValidFile(location)

	if isInvalid {
		return "path invalid : " + simplewrap.WithDoubleQuote(location)
	}

	return chmodhelper.FileModeFriendlyString(fileChmod) +
		constants.SpaceColonSpace +
		simplewrap.WithDoubleQuote(location)
}

func (it friendlyChmod) OfDirFiles(
	location string,
) string {
	if chmodhelper.IsPathInvalid(location) {
		return "path invalid : " + simplewrap.WithDoubleQuote(location)
	}

	parentDir := it.parentDir(location)

	scriptBuilder := errcmd.
		New.
		ScriptBuilder.
		DefaultDependingOnOs()

	if osconsts.IsWindows {
		scriptBuilder.Args(fmt.Sprintf(
			powershellListWithOwner,
			normalize.PathFixWithoutLongPathIf(
				true,
				parentDir)))
	} else {
		scriptBuilder.Args(simplewrap.WithDoubleQuote(parentDir))
	}

	dirChmodDisplay := it.OfPath(location)

	return dirChmodDisplay +
		constants.DefaultLine +
		scriptBuilder.DetailedOutput()
}

func (it friendlyChmod) LogOfPath(
	location string,
) {
	chmodMsg := it.OfPath(location)

	fmt.Println(chmodMsg)
}

func (it friendlyChmod) LogOfDirFiles(
	location string,
) {
	chmodMsg := it.OfDirFiles(location)

	fmt.Println(chmodMsg)
}

func (it friendlyChmod) parentDir(location string) string {
	if chmodhelper.IsDirectory(location) {
		return location
	}

	return filepath.Dir(location)
}
