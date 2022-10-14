package apachelinuxpath

import "gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"

func fixPathIf(isFix bool, root, next string) string {
	return normalizeinternal.JoinFixIf(
		isFix,
		root,
		next)
}
