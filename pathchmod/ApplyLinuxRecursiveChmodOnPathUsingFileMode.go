package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ApplyLinuxRecursiveChmodOnPathUsingFileMode(
	mode os.FileMode,
	rootDir string,
) *errorwrapper.Wrapper {
	fileMode := chmodhelper.
		New.
		RwxWrapper.
		UsingFileMode(mode)

	chmodErr := fileMode.
		LinuxApplyRecursive(false, rootDir)

	return errnew.
		Path.
		Error(errtype.ChmodApplyFailed, chmodErr, rootDir)
}
