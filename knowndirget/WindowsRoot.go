package knowndirget

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Returns Windows root as a string
//
// Drive letter + colon,
// from Windows 7 onwards it is c:\ constant.
// However to be precise it gets windir and gets the dir symbol
// Returns `c:\`
//
// Reference : https://play.golang.org/p/ZtXTSczrevD
func WindowsRoot() string {
	windowsDir := WidowsDirectory()
	colonIndex := strings.Index(windowsDir, constants.Colon)

	return windowsDir[0 : colonIndex+2]
}
