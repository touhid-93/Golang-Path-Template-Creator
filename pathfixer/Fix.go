package pathfixer

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Fix(isExpand bool, path string) string {
	if path == "" {
		return path
	}

	expand := expandpath.ExpandVariablesIf(
		isExpand, path)

	return normalize.Path(
		expand)
}
