package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/osconsts"
)

// JoinSimpleBaseWithMany doesn't apply normalize
func JoinSimpleBaseWithMany(
	baseDir string,
	relatives ...string,
) string {
	return baseDir +
		osconsts.PathSeparator +
		strings.Join(relatives, osconsts.PathSeparator)
}
