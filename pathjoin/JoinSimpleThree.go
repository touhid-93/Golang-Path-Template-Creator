package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinSimpleThree Doesn't apply normalize
func JoinSimpleThree(first, second, third string) string {
	return normalize.SimpleJoinPath3(first, second, third)
}
