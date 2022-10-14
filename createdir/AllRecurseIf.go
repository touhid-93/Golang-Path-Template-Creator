package createdir

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/dirinfo"
)

func AllRecurseIf(
	condition bool,
	path string,
	fileMode os.FileMode,
) *dirinfo.Result {
	if condition {
		return AllRecurse(path, fileMode)
	}

	return dirinfo.Empty()
}
