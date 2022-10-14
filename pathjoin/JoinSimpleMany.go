package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/osconsts"
)

// JoinSimpleMany doesn't apply normalize
func JoinSimpleMany(relatives ...string) string {
	return strings.Join(
		relatives,
		osconsts.PathSeparator)
}
