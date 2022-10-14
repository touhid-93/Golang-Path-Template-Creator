package pathjoin

import "gitlab.com/evatix-go/pathhelper/normalize"

func DbPath(first, second string) string {
	simpleJoin := JoinSimple(first, second)

	return normalize.DbPath(simpleJoin)
}
