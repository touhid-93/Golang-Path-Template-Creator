package apachelinuxpath

import "gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"

func fixPath(root, next string) string {
	return normalizeinternal.JoinFixIf(
		true,
		root,
		next)
}
