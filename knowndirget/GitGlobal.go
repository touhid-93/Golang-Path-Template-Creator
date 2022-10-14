package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to .git. Checks for it on all possible locations. If .git doesn't exist creates it.
// todo should we  use createDirectory to make .git . is it a directory or a file
func GitGlobal() string {
	homePath := UserPath()
	var outputPath, outputPathAlternate string

	if osconsts.IsWindows {
		outputPath = knowndir.GitGlobalWin.CombineWith(homePath)
	} else {
		outputPath = knowndir.GitGlobalUnix.CombineWith(homePath)
		outputPathAlternate = knowndir.GitGlobalUnixXdg.Value()
	}

	if !ispathinternal.Exists(outputPath) {
		outputPath = outputPathAlternate
	}

	return outputPath
}
